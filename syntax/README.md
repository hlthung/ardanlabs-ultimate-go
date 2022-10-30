
## Struct Best Practices

## Basic
Struct example represents a type with different fields.
```
type example struct {
    fla bool
    counter int 16
    pi float 32
}
```

Declare a variable of type example set to its zero value
```
var e1 example
```
Declare a variable of type example and init using a 'struct literal'

```
e2 := example {
    flag: true, 
    counter: 10,
    pi: 3.141592,
}

e2 := example {
    flag: true, 
    // counter: 10, // if you want a zero value for this, simply remove
    pi: 3.141592,
}
```
## Literal 
Not recommended to create empty literal construction syntax like this.

Reason: it doesn't always giving zero. It's an empty literal construction, not a zero value
```
e1 := example{}
```

Also not recommended to do as below, cause you'll get trouble later on
```
e.flag = true 
```
Just do
```
variable := true
e := example {
    flag: variable,
}
```
If you don't need the struct in multiple places, can just do on the fly (lireral type, type that doesn't have name) to avoid type polution.
```
// Anonymous type set to its zero value
var e1 struct {
    flag bool
    counter int16
    pi float32
}

// If you want literal type and something is not zero, then do
e2 := struct {
    flag bool
    counter int16
    pi float32
}{
    flag: true,
    counter: 10,
    pi: 3.141592,
}
```

## Display the value
Recommended, print key and value
```
fmt.Printf("%+v\n", e1) 
```
As this is gonna print just the value
```
fmt.Printf("%v\n", e1) 
```