# std

![GitHub CI](https://github.com/go-board/std/actions/workflows/go.yml/badge.svg)

[![Go Reference](https://pkg.go.dev/badge/github.com/go-board/std.svg)](https://pkg.go.dev/github.com/go-board/std)

![License](https://badgen.net/github/license/go-board/std)

![Last Commit](https://badgen.net/github/last-commit/go-board/std)

![Last Version](https://badgen.net/github/tag/go-board/std)

## Introduction

🥂 **`Std` is an enhanced version of the standard library based the new `Generics` feature.**

This project aims to provide a set of useful tools and libraries for the Go programming language.

Unlike using `interface{}`, this library can be used to create a generic type that can be used to create a type-safe API. And with `Generics`, no longer need use `reflect` package, so we can benifit from the performance.

## Installation
```bash
go get -u github.com/go-board/std
```

## Packages Hierarchy
- [algorithm](https://github.com/go-board/std/blob/master/algorithm) common used algorithms
    - [dp](https://github.com/go-board/std/blob/master/algorithm/dp) dynamic programming 
- [clone](https://github.com/go-board/std/blob/master/clone) clone a object
- [codec](https://github.com/go-board/std/blob/master/codec) encode and decode
- [collections](https://github.com/go-board/std/blob/master/collections) common used collections
    - [btree](https://github.com/go-board/std/blob/master/collections/btree) btree based map & set
    - [linkedlist](https://github.com/go-board/std/blob/master/collections/linkedlist) linked list
    - [queue](https://github.com/go-board/std/blob/master/collections/queue) double ended queue
- [cond](https://github.com/go-board/std/blob/master/cond) conditional operator
- [core](https://github.com/go-board/std/blob/master/core) core types & constraints
- [delegate](https://github.com/go-board/std/blob/master/delegate) delegate function signature
- [hash](https://github.com/go-board/std/blob/master/hash) hash a object
- [iterator](https://github.com/go-board/std/blob/master/iterator) iterators
    - [adapters](https://github.com/go-board/std/blob/master/iterator/adapters) adapter to create iterators & streams
    - [stream](https://github.com/go-board/std/blob/master/iterator/stream) stream processing
- [lazy](https://github.com/go-board/std/blob/master/lazy) lazy evaluation & variables
- [optional](https://github.com/go-board/std/blob/master/optional) optional values
- [ptr](https://github.com/go-board/std/blob/master/ptr) convenient pointer operator
- [result](https://github.com/go-board/std/blob/master/result) result values
- [service](https://github.com/go-board/std/blob/master/service) service abstractions
- [sets](https://github.com/go-board/std/blob/master/sets) hashset using builtin map
- [slices](https://github.com/go-board/std/blob/master/slices) slice functors
- [tuple](https://github.com/go-board/std/blob/master/tuple) tuple type from 2 to 5
