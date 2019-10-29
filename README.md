# goserver_example
An exercise in demonstrating simple authentication between client and server.

Some notes:

- Database models following a loose repository pattern [reference](https://medium.com/bumpers/our-go-is-fine-but-our-sql-is-great-b4857950a243)
- [Gingko BDD](http://onsi.github.io/ginkgo/)

## Flaws

_ TODO elucidate_

- ServeHTTP handler generality
- Naming consistencies
- Server's hardcoded configuration in main
- Documentation (at least sequence diagram in this readme)
- Helper function commonality/design (ties into first point)
- Lack of client dependency injection
- Very limited request format (see next point)
- Should have used gRPC in hindsight
