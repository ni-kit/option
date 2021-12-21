package httpresult_test

import (
	"fmt"
	"github.com/debudda/option/httpresult"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

type User struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func ExampleJSON() {
	res := httpresult.JSON[User](http.Get(url))
	res.Switch(
		func(u User) {
			fmt.Println(u.ID)
		},
		func(err error) {
			fmt.Println("couldn't fetch user", err)
		},
	)
	// Output: 1
}

func ExampleJSON_2() {
	res := httpresult.JSON[User](http.Get(url))
	res.Ok(func(u User) {
		fmt.Println(u.ID)
	})
	// Output: 1
}

func ExampleJSONCode() {
	res, code := httpresult.JSONCode[User](http.Get(url))
	res.Ok(func(u User) {
		fmt.Println(code, u.ID)
	})
	// Output: 200 1
}

func ExampleResponseJSON() {
	res := httpresult.ResponseJSON[User](http.Get(url))
	res.Switch(
		func(res *httpresult.SimpleJSONResponse[User]) {
			fmt.Println(res.StatusCode, res.Body.ID)
		},
		func(err error) {
			fmt.Println("couldn't fetch user", err)
		},
	)
	// Output: 200 1
}
