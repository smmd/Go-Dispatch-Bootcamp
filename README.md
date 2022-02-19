# Go-Dispatch-Bootcamp

This API is an exercises that should contain:

- An endpoint for reading from an external API
  - Write the information in a CSV file
- An endpoint for reading the CSV
  - Display the information as a JSON
- An endpoint for reading the CSV concurrently with some criteria (details below)
- Unit testing for the principal logic
- Follow conventions, best practices
- Clean architecture
- Go routines usage
- Authentication token to use as header in all endpoints

External API used: [Pokemon](https://pokeapi.co/)

### Usage examples for this API.

Running server:
```bash
go run main.go
```

To access API endpoint's is needed generate a token.
```curl
curl http://localhost:3001/generate-token/
```

GET All Poke monsters:
```curl
curl http://localhost:3001/pokemonsters/
```

GET Poke monster by ID:
```curl
curl http://localhost:3001/pokemonsters/{id}
```

Consume external api and populate CSV:
```curl
http://localhost:3001/fill-pokedex/
```

GET Poke monsters with worker pool:
```curl
http://localhost:3001/:type/:items/:items_per_workers
```
`type: "odd" or "even"`,
`items: amount of valid items response`,
`items_per_workers: amount of valid items per worker`

### How to run Unit test suite.

```bash
go test ./...
```

## Documentation

### Must to learn

- [Go Tour](https://tour.golang.org/welcome/1)
- [Go basics](https://www.youtube.com/watch?v=C8LgvuEBraI)
- [Git](https://www.youtube.com/watch?v=USjZcfj8yxE)
- [Tool to practice Git online](https://learngitbranching.js.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [How to write code](https://golang.org/doc/code.html)
- [Go by example](https://gobyexample.com/)
- [Go cheatsheet](http://cht.sh/go/:learn)
- [Any talk by Rob Pike](https://www.youtube.com/results?search_query=rob+pike)
- [The Go Playground](https://play.golang.org/)

### Self-Study Material

- [Golang Docs](https://golang.org/doc/)
- [Constants](https://www.youtube.com/watch?v=lHJ33KvdyN4)
- [Variables](https://www.youtube.com/watch?v=sZoRSbokUE8)
- [Types](https://www.youtube.com/watch?v=pM0-CMysa_M)
- [For Loops](https://www.youtube.com/watch?v=0A5fReZUdRk)
- [Conditional statements: If](https://www.youtube.com/watch?v=QgBYnz6I7p4)
- [Multiple options conditional: Switch](https://www.youtube.com/watch?v=hx9iHend6jM)
- [Arrays and Slices](https://www.youtube.com/watch?v=d_J9jeIUWmI)
- [Clean Architecture](https://medium.com/@manakuro/clean-architecture-with-go-bce409427d31)
- [Maps](https://www.youtube.com/watch?v=p4LS3UdgJA4)
- [Functions](https://www.youtube.com/watch?v=feU9DQNoKGE)
- [Error Handling](https://www.youtube.com/watch?v=26ahsUf4sF8)
- [Structures](https://www.youtube.com/watch?v=w7LzQyvriog)
- [Structs and Functions](https://www.youtube.com/watch?v=RUQADmZdG74)
- [Pointers](https://tour.golang.org/moretypes/1)
- [Methods](https://www.youtube.com/watch?v=nYWa5ECYsTQ)
- [Interfaces](https://tour.golang.org/methods/9)
- [Interfaces](https://gobyexample.com/interfaces)
- [Packages](https://www.youtube.com/watch?v=sf7f4QGkwfE)
- [Failed requests handling](http://www.metabates.com/2015/10/15/handling-http-request-errors-in-go/)
- [Modules](https://www.youtube.com/watch?v=Z1VhG7cf83M)
  - [Part 1 and 2](https://blog.golang.org/using-go-modules)
- [Unit testing](https://golang.org/pkg/testing/)
- [Go tools](https://dominik.honnef.co/posts/2014/12/an_incomplete_list_of_go_tools/)
- [More Go tools](https://dev.to/plutov/go-tools-are-awesome-bom)
- [Functions as values](https://tour.golang.org/moretypes/24)
- [Concurrency (goroutines, channels, workers)](https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3)
  - [Concurrency Part 2](https://www.youtube.com/watch?v=LvgVSSpwND8)
