package option_test

import (
	"fmt"
	"github.com/debudda/option"
)

type User struct {
	Name string
	Age  int
}

func ExampleSwitch() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	isSome := option.Switch(maybeUser,
		func(u User) {
			fmt.Printf("Got user %s of age %d\n", u.Name, u.Age)
		},
		func() {
			fmt.Println("Oops! No user")
		})
	fmt.Println("is some?", isSome)
	// Output: Got user Douglas Adams of age 42
	// is some? true
}

func ExampleSwitch_none() {
	maybeUser := option.O[User]()
	isSome := option.Switch(maybeUser,
		func(u User) {
			fmt.Printf("Got user %s of age %d\n", u.Name, u.Age)
		},
		func() {
			fmt.Println("Oops! No user")
		})
	fmt.Println("is none?", !isSome)
	// Output: Oops! No user
	// is none? true
}

func ExampleSwitcht() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := option.Switcht(maybeUser,
		func(u User) {
			fmt.Printf("Got user %s of age %d\n", u.Name, u.Age)
		},
		func() {
			fmt.Println("Oops! No user")
		})
	fmt.Println(*someUser)
	// Output: Got user Douglas Adams of age 42
	// {Douglas Adams 42}
}

func ExampleSwitcht_ptr() {
	maybeUser := option.O(&User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := option.Switcht(maybeUser,
		func(u *User) {
			u.Age = 24
		},
		func() {
			fmt.Println("Oops! No user")
		})
	fmt.Printf("User: %v\n", *someUser)
	// Output: User: &{Douglas Adams 24}
}

func ExampleTestSwitchv() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := option.Switchv(maybeUser,
		func(u User) User {
			u.Age = 24
			return u
		},
		func() User {
			fmt.Println("Oops! No user")
			return User{Name: "Default"}
		})
	fmt.Println(someUser)
	// Output: {Douglas Adams 24}
}

func ExampleTestDefault() {
	maybeUser := option.O[User]()

	someUser := option.Default(maybeUser,
		User{Name: "Douglas Adams", Age: 42})
	fmt.Println(someUser)
	// Output: {Douglas Adams 42}
}

func ExampleTestDefault_2() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := option.Default(maybeUser,
		User{Name: "Douglas Adams", Age: 24})
	fmt.Println(someUser)
	// Output: {Douglas Adams 42}
}

func ExampleTestDefaultv() {
	maybeUser := option.O[User]()

	someUser := option.Defaultv(maybeUser,
		User{Name: "Douglas Adams", Age: 42},
		func(u User) User {
			u.Age = 24
			return u
		})
	fmt.Println(someUser)
	// Output: {Douglas Adams 42}
}

func ExampleTestDefaultv_2() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := option.Defaultv(maybeUser,
		User{Name: "Douglas Adams", Age: 24},
		func(u User) User {
			u.Age = 49
			return u
		})
	fmt.Println(someUser)
	// Output: {Douglas Adams 49}
}
