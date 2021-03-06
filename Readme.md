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
user := maybeUser.Default(defaultUser)
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
user := maybeUser.Default(defaultUser)
fmt.Println(user == defaultUser) // => true
```

Unpack _None[User]_ (verbose):
```go
user := maybeUser.Switchv(
    func(u User) User { return u },
    func() User { return defaultUser })
fmt.Println(user == defaultUser) // => true
```

Switch on Option:
```go 
maybeUser := option.O(aUser)
maybeUser.Switch(
    func(u User) {
        fmt.Printf("Got a user %s, aged %d\n", u.Name, u.Age)
    },
    func() {
        fmt.Println("Got nothing")
    })
```

### Iterate

```go 
opts := option.Slice[User](
    User{Name: "Douglas", Age: 42},
    User{Name: "Neil", Age: 61},
)

opts.EachPtr(func(s *User) {
    s.Age*=2
})
opts.Each(func(s User) {
    fmt.Println(s)
})
// => {Douglas 84}
// => {Neil 122}

ultimateWriter := option.Foldli(opts, func(i int, res User, next User) User {
    if i != len(opts)-1 {
        res.Name += next.Name + " "
    }
    res.Age += next.Age
    return res
}, User{})
fmt.Println(ultimateWriter) // {Douglas Neil 206}
```

### Http 

```go
type User struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

res := httpresult.JSON[User](http.Get("https://jsonplaceholder.typicode.com/todos/1"))
res.Ok(func(u User) {
    fmt.Printf("%#v\n", u)
})
// => User{UserID:1, ID:1, Title:"delectus aut autem", Completed:false}

res = httpresult.JSON[User](http.Get("https://incorrect"))
u := res.Default(User{Title: "so default"})
fmt.Printf("%#v\n", u)
// => User{UserID:0, ID:0, Title:"so default", Completed:false}
```