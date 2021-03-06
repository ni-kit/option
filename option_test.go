package option_test

import (
	"errors"
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

	isSome := maybeUser.Switch(
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
	isSome := maybeUser.Switch(
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

func ExampleSwitchPtr() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	isSome := maybeUser.SwitchPtr(
		func(u *User) {
			fmt.Printf("Got user %s of age %d\n", u.Name, u.Age)
			u.Age++
		},
		func() {
			fmt.Println("Oops! No user")
		})
	fmt.Println("is some?", isSome)
	maybeUser.Some(func(u User) {
		fmt.Println("age", u.Age)
	})
	// Output: Got user Douglas Adams of age 42
	// is some? true
	// age 43
}

func ExampleSwitcht() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := maybeUser.Switcht(
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

	someUser := maybeUser.Switcht(
		func(u *User) {
			u.Age = 24
		},
		func() {
			fmt.Println("Oops! No user")
		})
	fmt.Printf("User: %v\n", *someUser)
	// Output: User: &{Douglas Adams 24}
}

func ExampleSwitchPtrt() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := maybeUser.SwitchPtrt(
		func(u *User) {
			fmt.Printf("Got user %s of age %d\n", u.Name, u.Age)
			u.Age++
		},
		func() {
			fmt.Println("Oops! No user")
		})
	fmt.Println(*someUser)
	// Output: Got user Douglas Adams of age 42
	// {Douglas Adams 43}
}

func ExampleTestSwitchv() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := maybeUser.Switchv(
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

func ExampleTestSwitchPtrv() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := maybeUser.SwitchPtrv(
		func(u *User) {
			fmt.Println("got user age", u.Age)
			u.Age = 24
		},
		func() User {
			fmt.Println("Oops! No user")
			return User{Name: "Default"}
		})
	fmt.Println(someUser)
	// Output: got user age 42
	// {Douglas Adams 24}
}

func ExampleTestDefault() {
	maybeUser := option.O[User]()

	someUser := maybeUser.Default(User{Name: "Douglas Adams", Age: 42})
	fmt.Println(someUser)
	// Output: {Douglas Adams 42}
}

func ExampleTestDefault_2() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	someUser := maybeUser.Default(User{Name: "Douglas Adams", Age: 24})
	fmt.Println(someUser)
	// Output: {Douglas Adams 42}
}

func ExampleTestDefaultv() {
	maybeUser := option.O[User]()

	someUser := maybeUser.Defaultv(User{Name: "Douglas Adams", Age: 42},
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

	someUser := maybeUser.Defaultv(User{Name: "Douglas Adams", Age: 24},
		func(u User) User {
			u.Age = 49
			return u
		})
	fmt.Println(someUser)
	// Output: {Douglas Adams 49}
}

func ExampleTestDefaultPtrv() {
	maybeUser := option.O[User]()

	someUser := maybeUser.DefaultPtrv(User{Name: "Douglas Adams", Age: 42},
		func(u *User) {
			u.Age = 24
		})
	fmt.Println(someUser)
	// Output: {Douglas Adams 42}
}

func ExampleOption_Some() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	maybeUser.Some(func(u User) {
		fmt.Println(u.Name)
	})
	// Output: Douglas Adams
}

func ExampleOption_Some_2() {
	maybeUser := option.O[User]()

	maybeUser.Some(func(u User) {
		fmt.Println(u.Name)
	})
	// Output:
}

func ExampleOption_SomePtr() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	maybeUser.SomePtr(func(u *User) {
		fmt.Println(u.Name)
	})
	// Output: Douglas Adams
}

func ExampleOption_SomePtr_2() {
	maybeUser := option.O[User]()

	maybeUser.SomePtr(func(u *User) {
		fmt.Println(u.Name)
	})
	// Output:
}

func ExampleOption_SomePtr_3() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	maybeUser.SomePtr(func(u *User) {
		u.Age++
	})

	user := maybeUser.Default(User{})
	fmt.Println(user.Age)
	// Output: 43
}

func ExampleOption_SomePtrt() {
	maybeUser := option.O(User{
		Name: "Douglas Adams",
		Age:  42,
	})

	u := maybeUser.SomePtrt(func(u *User) *User {
		u.Age = 49
		return u
	})
	fmt.Println(u.Age)
	// Output: 49
}

func ExampleOption_OkOrElse() {
	maybeUser := option.O(User{
		Name: "John Doe",
		Age:  123,
	})

	maybeUser.OkOrElse(func() error {
		return errors.New("got no user")
	}).Switch(
		func(u User) {
			fmt.Printf("%#v\n", u)
		},
		func(err error) {
			fmt.Println("got error", err)
		},
	)
	// Output: option_test.User{Name:"John Doe", Age:123}
}

func Example_And() {
	type Writer struct {
		Name string
	}
	maybeWriter := option.O(Writer{
		Name: "Douglas Adams",
	})
	maybeUser := option.O(User{
		Name: "John Doe",
		Age:  123,
	})

	option.And(maybeWriter, maybeUser).Some(func(u User) {
		fmt.Printf("%#v\n", u)
	})
	// Output: option_test.User{Name:"John Doe", Age:123}
}

func ExampleOption_AndThen() {
	sq := func(n int) option.Option[int] { return option.O(n * n) }

	option.O(2).AndThen(sq, sq).Some(func(n int) {
		fmt.Println(n)
	})
	// Output: 16
}
