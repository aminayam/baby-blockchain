## Terms of reference

Looking through articles on the Internet, I came across an interesting case of using blockchain technology - __gems accounting system__, in particular diamonds.
I implemented a simpler version of it on Golang.

* `Purpose of the system.` Accounting for unique non-fungible tokens. Each stone in the system initially has the correct original id.
* `System content.` In my model, I will consider only the transfer of ownership within the system, "money" transfers will remain outside the system. The user's balance will be a list of original identifiers that he owns, therefore, with the help of transactions, he can transfer a certain id, overwriting it in the owner's database
* `Interaction with other products` Not expected yet
* `Product features.`
  * the ability to find out the current owner (public database data);
  * confirmation of ownership and transfer of ownership for users of the system;
  * view your own "balance".
* `Characteristics of users.` Users = verified owners (and potential owners) of tokens = jewels.

## Working with the product

The entry point is the `main.go` file in root of project.
There you can see an introductory tour of the system, the comments describe the functionality and features. Also, you can manually change the code according to the prompts in order to conduct your research.