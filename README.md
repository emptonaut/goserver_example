# goserver_example
An exercise in demonstrating simple authentication between client and server.

## Usage

The server takes no flags--just `make all` and run `bin/server`.

The client has full help output. To see it, run `bin/client`.

## Some random notes:

- Database models following a loose repository pattern [reference](https://medium.com/bumpers/our-go-is-fine-but-our-sql-is-great-b4857950a243)
- [Gingko BDD](http://onsi.github.io/ginkgo/)

## Prerequisites

These executables need to be available and generally aren't installed by default on Mac or Linux:

- go
- sqlite3

## Flaws

These are _some_ of the things that would need to be changed, or otherwise should be changed, before this application would be deployed.

- Passwords are entered on the client in plaintext
- Default ServeHTTP handler requires actual endpoints to reparse request bodies
- Helper function commonality/design (ties into first point)
    - There's a lot of boilerplate in the client functions and endpoints
- Naming consistencies
- Server's hardcoded configuration in main.
- Client should not be configured to read a specific root CA.
- Documentation (at least sequence diagram in this readme)
- The http.Client usage in the client should be replaced with an interface (ideally a built in type) for dependency injection
- Very limited request format (see next point)
- Should have used gRPC in hindsight

## Screenshot of Demo

![demo working](https://raw.githubusercontent.com/shoelick/goserver_example/master/demo.png)
