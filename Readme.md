### Option

Option provides an Option container which can be used to force some additional presence assertions on values.  

### Example  

**Having**:
```go 
type User struct {
    Name string
    Age  int
}
defaultUser := User{Name: "A", Age: 1}
aUser := User{Name: "B", Age: 2}
```

Unpack _Some[User]_:
```go
// option.Option[User]
maybeUser := option.O(aUser)
// option is a user, use value, ignore default
user := option.Default(maybeUser, defaultUser)
fmt.Println(user == aUser) // => true
```

Call a function if Option is Some:
```go
// option.Option[User]
maybeUser := option.O(aUser)
maybeUser.Some(func(u User) {
	fmt.Println("name: ", u.Name)
}) // name: B
```

Unpack _None[User]_:
```go
maybeUser := option.O[User]() // option.None[User]
// No user, use default
user := option.Default(maybeUser, defaultUser)
fmt.Println(user == defaultUser) // => true
```

Unpack _None[User]_ (verbose):
```go
user := option.Switchv(maybeUser,
    func(u User) User { return u },
    func() User { return defaultUser })
fmt.Println(user == defaultUser) // => true
```

Switch on Option:
```go 
maybeUser := option.O(aUser)
option.Switch(maybeUser,
    func(u User) {
        fmt.Printf("Got a user %s, aged %d\n", u.Name, u.Age)
    },
    func() {
        fmt.Println("Got nothing")
    })
```
