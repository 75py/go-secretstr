package main

import (
	"encoding/json"
	"fmt"
	"github.com/75py/go-secretstr"
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
