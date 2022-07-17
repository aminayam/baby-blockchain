package main

import (
	"baby-blockchain/Blocks"
	"baby-blockchain/OperTx"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}
var bc Blocks.Blockchain

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	//bc := Blocks.InitBlockchain()
	go func() {
		genesisBlock := Blocks.NewGenesisBlock()
		spew.Dump(genesisBlock)

		mutex.Lock()
		bc.Blocks = append(bc.Blocks, genesisBlock)
		mutex.Unlock()
	}()
	log.Fatal(run())

}

// web server
func run() error {
	muxRouter := makeMuxRouter()
	httpPort := os.Getenv("PORT")
	log.Println("HTTP Server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        muxRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

// create handlers
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

// write blockchain when we receive an http request
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err = io.WriteString(w, string(bytes)); err != nil {
		log.Fatal(err)
	}
}

// takes JSON payload as an input for heart rate (BPM)
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg OperTx.Operation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Blocks.CreateBlock(OperTx.CreateTransaction(msg, 0), prevBlock.BlockID)

	if Blocks.ValidateBlock(newBlock, prevBlock) {
		bc.Blocks = append(bc.Blocks, newBlock)
		spew.Dump(bc)
	}
	mutex.Unlock()

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)

}
