<p align="center" style="margin-bottom: 20px; width: 100px; height: 100px; margin: auto">
<img src="https://gophercises.com/img/gophercises_jumping.gif" width="250px"/>
</p>

# go-calculator

Calculator based on shunting-yard algorithm whish is a method for parsing mathematical expressions specified in infix notation.

See documentation [Shunting-yard_algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm).

## Prerequisites

Install [Golang](https://golang.org/doc/install) or you could use gvm (Golang version manager) to do it.

Follow those steps for install GVM : [github.com/moovweb/gvm](https://github.com/moovweb/gvm#installing)

## Usage

* go get project locally

```
$ go get github.com/Woukiwoukiwa/go-calculator
```

* Run calculator with infix expression

```
$ go-calculator "1 + 3"
```

## Available operators

| Operator | Operation | What it returns |
|---|---|---|
| + | x + y | Sum of x and y |
| - | x - y | Difference of x and y |
| * | x * y | Product of x and y |
| / | x / y | Quotient of x and y |
| % | x % y | Remainder of x / y |
| ( |  | Open parenthesis |
| ) |  | Closed parenthesis |

## Issues

- [ ] Error management improvment
- [ ] Square calculation
- [ ] Cos, Sin, Tan calculation
- [ ] Pi integration

## Dependencies

This project depends on

* [golang](http://golang.org/)