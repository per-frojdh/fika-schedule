# fika-schedule
Fika scheduler for Kartena

This is a learning project, so there is going to be a lot of ugliness. You have been warned.

## How to get started (don't clone me!)
Note that this is definitely not *the* way to go about this. It's more along a quick and dirty solution, and I hope it sort of works.

* Download [golang binaries](https://golang.org/dl/) (>1.5)
* Make sure the binaries are added to PATH
* Create what's known as a GOPATH, typically it would be a source project folder, like $HOME/work
    * It's like an additional PATH variable, specific for all your Go related stuff.
* Navigate to the folder with your favourite terminal.
* Run ```go get github.com/per-frojdh/fika-schedule```
* Watch as the magic unfolds (or complain to me if something doesnt work):
    * Go just created the necessary folder structure for you.
    * Go downloaded all of the required dependencies of this project.
    * Go even compiled the package for you (look under ```\bin```)
    * Run it as any other binary file.
    
## How to get started (for real this time)

### Endpoints
* Create a user ```POST -> :8000/user``` with a "name" field.
    * Send in a username, this will be the key for the db.
* Fetch all registered users ```GET -> :8000/user```
    * Get a single user ```GET -> :8000/user/name```
* Fetch the next user to do fika ```GET -> :8000/fika```
    * This crashes sometimes, unsure why (yet)
* When a user has done their fika they should do ```PUT -> :8000/fika/name```
    * This will update their fika timestamp.
* To delete a user ```DELETE -> :8000/user/name```

## How does it all work
* We use a embedded database (built in Go) called [BoltDB](https://github.com/boltdb/bolt), it's a key/value db.
    * It's incredibly fast, but locks down the file so it cannot be used by multiple processes.
    * I'm not always closing the connection, so we crash sometimes (this will be fixed, sometime)
* We use the standard lib for handling connections.
    * We use a pretty standard router though, because it's something thats lacking in the standard lib.
    * It's called mux and adds context and some more features like parsing queryStrings and url params.
* The way most handlers (at least the ones working with the database is the sort of following)
    * Access the db, we receive JSON (I think, Go handles it as raw byte arrays)?
    * Unmarshal the JSON(or is it bytes?) into structs that we can read, write, manipulate etc.
    * Marshal the struct into JSON that we can return to the user.

## What's left to do
* Error handling is a mess, but that's mostly cause nothing is validated.
* Database handling needs to get improved, occasionally we get locked out of file access.
* Cleaning up the structs and objects.
    * Right now it's a semi-structured chaotic mess.
    * I've switched from "Object-Oriented" style, like adding static methods to structs back and forth a few times.
* Timestamps aren't really readable at the moment, my attempts at making it human readable failed so I've resorted to the standard output.
    * The standard lib for time seems good, I just haven't figured it out yet.
* Some sort of logs would be sweet.
* Keep track of "authentication" keys and require them for communicating with the API
* Much much more.
    
## Fun things left to look at
* [EasyJSON](https://github.com/mailru/easyjson)
    * Basically we skip the standard lib json marshaller/unmarshaller because they are too generic
    * In CLI, we produce a struct straight into a file with a marshaller/unmarshaller interface specific for the properties (less generic, much faster)
    
