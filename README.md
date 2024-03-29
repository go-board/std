# std

![GitHub CI](https://github.com/go-board/std/actions/workflows/go.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-board/std.svg)](https://pkg.go.dev/github.com/go-board/std)
![License](https://badgen.net/github/license/go-board/std)

[![codecov](https://codecov.io/gh/go-board/std/branch/master/graph/badge.svg?token=SYZ3UBT4GD)](https://codecov.io/gh/go-board/std)

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
- [clone](https://github.com/go-board/std/blob/master/clone) clone a object
- [codec](https://github.com/go-board/std/blob/master/codec) encode and decode
- [collections](https://github.com/go-board/std/blob/master/collections) common used collections
    - [btree](https://github.com/go-board/std/blob/master/collections/btree) btree based map & set
    - [linkedlist](https://github.com/go-board/std/blob/master/collections/linkedlist) linked list
    - [queue](https://github.com/go-board/std/blob/master/collections/queue) double ended queue
- [cond](https://github.com/go-board/std/blob/master/cond) conditional operator
- [constraints](https://github.com/go-board/std/blob/master/constraints) core constraints
- [fp](https://github.com/go-board/std/blob/master/fp) functional programing
- [hash](https://github.com/go-board/std/blob/master/hash) hash a object
- [iter](https://github.com/go-board/std/blob/master/iter) iterators
    - [collector](https://github.com/go-board/std/blob/master/iterator/collector) consume iter and collect to another type
    - [source](https://github.com/go-board/std/blob/master/iterator/source) adapter to create iterators & streams
- [lazy](https://github.com/go-board/std/blob/master/lazy) lazy evaluation & variables
- [optional](https://github.com/go-board/std/blob/master/optional) optional values
- [ptr](https://github.com/go-board/std/blob/master/ptr) convenient pointer operator
- [result](https://github.com/go-board/std/blob/master/result) result values
- [service](https://github.com/go-board/std/blob/master/service) service abstractions
- [sets](https://github.com/go-board/std/blob/master/sets) hashset using builtin map
- [slices](https://github.com/go-board/std/blob/master/slices) slice functors
- [tuple](https://github.com/go-board/std/blob/master/tuple) tuple type from 2 to 5
