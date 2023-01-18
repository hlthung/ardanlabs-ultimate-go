# Other concepts

## Pointers

- https://github.com/ardanlabs/gotraining/tree/master/topics/go/language/pointers

- https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html

- https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac

## Blank Identifier

`updateUser` returns update stats and error. But we don't care about the update stats, so we can use `_` cause compiler will later complain the variable declared but not used.
```
if _, err := updateUser(&u); err != nil {
    fmt.Println(err)
    return
}
```

If you don't want anything from the function, don't do this
```
if _, _ := updateUser(&u); err != nil {
    fmt.Println(err)
    return
}
```

Just do this

```
if updateUser(&u); err != nil {
    fmt.Println(err)
    return
}
```

There's linter out there will complain that you are not handling the error. It's kinda no, no in the Go community. Though there's exception, but overall you should be checking error.

## Function Idioms
1. If your function return more than 2 values (up to 3), should consider refactoring
2. Avoid using `else`, use if > switch > if.. else..
3. After calling a function, return immediately if there's an error
```
func retrieveUser(name string) (*user, error) {
    r, err := getUser(name)
    if err != nil {
        return nil, err
    }

    // Unmarshal the json document
    var u user
    err = json.Unmarshal([]byte(r), &u)
    return &u, err
}

func main() {
    // Retrieve the user profile.
    u, err := retrieveUser("sally")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Display the user profile.
    fmt.Printf("%+v\n", *u)
}
```
For this piece of code below

```
    // Unmarshal the json document
    var u user
    err = json.Unmarshal([]byte(r), &u)
    return &u, err
```

You can also write like (recommended)

```
    // Unmarshal the json document
    var u user
    if err := json.Unmarshal([]byte(r), &u); err != nil {
        return nil, err
    }
    return &u, nil
```
So that reader won't get confused with the err, whether it has error or is actually nil

If you need to return two values, such as below, then do two lines, else above line is the recommended way of writing

```
    r, err := getUser(name)
    if err != nil {
        return nil, err
    }
```

## Literal Functions
Though we can write something like this, normally we don't do this:
```
    // Declare an anonymous function and call it.
    func() {
        fmt.Println("Test:", n)
    }()
   
```
Can see people doing this in some scnenarios:
```
    // Declare an anonymous function and call it.
    f := func() {
        fmt.Println("Test:", n)
    }

    f()
```
Most of the time we do with defer():
```
    // Defer the call to the anonymous function till after main returns
    defer func() {
        fmt.Println("Test:", n) // print n = 3
    }()

    n = 3
```

## Short Variable Declaration
Don't do err1, err2 etc as you don't need to! When you use short variable declaration operation, it will construct the new variable, and reuse err on the call. 

```
    u, err := getUser()
    if err != nil {
        return 
    }
    
    c, err := getCustomer(u)
    if err != nil {
        return 
    }
```

Note: you need to have at least 1 new variables on left side of := to make it work

## Array


Do you know you can do this? Yes, you can but DON'T DO THIS
```
numbers := [4]int{0:10, 2:20, 1:30, 3:40}
numbers := [...]int{10, 20, 30, 40}
```
Do this
```
numbers := [4]int{10, 20, 30, 40}
```
Go has the traditional for loop syntax as below
```
for i := 0; i < len(numbers); i++ {
    fmt.Println(i, numbers[i])
}
```

But we should make use of this
```
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
```
This is when you want to operate on the copies. 

Or when you want to operate on the orignal data, do this
```
for i := range fruits {
    fmt.Println(i, fruit[i])
}
```

Also be aware that you can't do
```
var five [5]int
four := [4]int{10, 20, 30, 40}

five = four
```
You will get error message: cannot use four (type [4]int) as type [5]int in assignment
Because they are different type (notice the wording, we are not saying size, but type). 

Do note that size is part of the type information.
