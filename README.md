GBASIC
===

[Golang](https://go.dev/) implementation of the [BASIC programming language](https://en.wikipedia.org/wiki/BASIC).

### Install Antler4

Install the [ANTLR4](https://www.antlr.org/) tool and generate the parser code for the BASIC grammar.
Install ANTLR4 tool by following the instructions provided in the official documentation: https://www.antlr.org/

```
$ pip install antlr4-tools
```

## Build

```
$ antlr4 -Dlanguage=Go parser/BASIC.g4
```

## Run

```
$ go run main.go
```
