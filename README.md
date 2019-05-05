# go-secretstr

[![Build Status](https://travis-ci.org/75py/secretstr.svg?branch=master)](https://travis-ci.org/75py/secretstr)
[![codecov](https://codecov.io/gh/75py/secretstr/branch/master/graph/badge.svg)](https://codecov.io/gh/75py/secretstr)
[![GoDoc](https://godoc.org/github.com/75py/secretstr?status.svg)](https://godoc.org/github.com/75py/secretstr)

## Install

```bash
go get github.com/75py/secretstr
```

## Documents

https://godoc.org/github.com/75py/secretstr

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/75py/secretstr"
)

type UnsafeLoginForm struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type LoginForm struct {
	ID       secretstr.SecretString /* instead of string */ `json:"id"`
	Password secretstr.SecretString /* instead of string */ `json:"password"`
}

func main() {
	input := []byte(`{"id":"raw_id","password":"raw_password"}`)

	var unsafeLoginForm UnsafeLoginForm
	_ = json.Unmarshal(input, &unsafeLoginForm)
	var loginForm LoginForm
	_ = json.Unmarshal(input, &loginForm)

	fmt.Printf("fmt.Printf(\"%%s\", unsafeLoginForm) => %s \n", unsafeLoginForm)
	fmt.Printf("fmt.Printf(\"%%v\", unsafeLoginForm) => %v \n", unsafeLoginForm)
	fmt.Printf("fmt.Printf(\"%%+v\", unsafeLoginForm) => %+v \n", unsafeLoginForm)
	fmt.Printf("fmt.Printf(\"%%#v\", unsafeLoginForm) => %#v \n", unsafeLoginForm)

	fmt.Printf("fmt.Printf(\"%%s\", loginForm) => %s \n", loginForm)
	fmt.Printf("fmt.Printf(\"%%v\", loginForm) => %v \n", loginForm)
	fmt.Printf("fmt.Printf(\"%%+v\", loginForm) => %+v \n", loginForm)
	fmt.Printf("fmt.Printf(\"%%#v\", loginForm) => %#v \n", loginForm)
}
```

Output
```
fmt.Printf("%s", unsafeLoginForm) => {raw_id raw_password} 
fmt.Printf("%v", unsafeLoginForm) => {raw_id raw_password} 
fmt.Printf("%+v", unsafeLoginForm) => {ID:raw_id Password:raw_password} 
fmt.Printf("%#v", unsafeLoginForm) => main.UnsafeLoginForm{ID:"raw_id", Password:"raw_password"} 
fmt.Printf("%s", loginForm) => {[FILTERED] [FILTERED]} 
fmt.Printf("%v", loginForm) => {[FILTERED] [FILTERED]} 
fmt.Printf("%+v", loginForm) => {ID:[FILTERED] Password:[FILTERED]} 
fmt.Printf("%#v", loginForm) => main.LoginForm{ID:[FILTERED], Password:[FILTERED]} 
```
