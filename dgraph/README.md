# Dgraph Notes
Dgraph is (for now) an open source graph database with first-class GraphQL-ish support. It is written in GO and designed to be deployed horizontally across the web. 

## Security
[Dgraph Docs](https://dgraph.io/docs/graphql/security/) address this by saying: handle your own security!

.... but it may be handled if I never expose the database to the outside world, and only allow pocketbase to interact with Dgraph inside an isolated Docker network. 

## Hello_World
The folder `go-hw-dgraph` holds some example code, written in GO to show how it is possible to interact with Dgraph. 

## Configuration
[Dgraph Docs](https://dgraph.io/docs/deploy/config/) address configuration a few ways: config files, but those are overridden by environment variables and CLI commands. The configuration file here called `config.yaml` is not used, but could serve as an example for exploration.

## Example Requests
The `ex_req.md` file holds a few example requests to get started with Dgraph. Further reading is also linked.

These are copy/paste examples to use with Ratel, taken direct from [here](https://dgraph.io/docs/dql/dql-get-started/) in the Dgraph docs. 