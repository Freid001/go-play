# Go Play

## About 
* A series of simple cli games.
* The purpose of this project was to gain some experience writing an application in Golang.

### Requirements
* [Docker](https://www.docker.com/) 

### Installation
```bash
// build 
docker build -t go_play .

// run
docker run -it go_play bash
```

## Useage

```bash
// run
go run *.go

// hello
go run *.go hello -name="your_name"

// time
go run *.go time

// guess number
go run *.go guessnumber

// fizzbuss
go run *.go fizzbuzz -numbers="comma_separated_number_string"
```
