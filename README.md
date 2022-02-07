# GO Entropy

[![Build Status](https://github.com/prongbang/goentropy/workflows/Go/badge.svg)](https://github.com/prongbang/goentropy/actions)
[![Codecov](https://img.shields.io/codecov/c/github/prongbang/goentropy.svg)](https://codecov.io/gh/prongbang/goentropy) 
[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/goentropy)](https://goreportcard.com/report/github.com/prongbang/goentropy)

## Install

```shell
go get github.com/prongbang/goentropy
```

## Used

```go
e := goentropy.New()

r := e.Calc("123456")

fmt.Println("Entropy:", r)
```

Output

```
Entropy: 2.584963
```
