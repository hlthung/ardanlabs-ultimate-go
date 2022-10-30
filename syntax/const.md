
# Constants

## Untyped vs Typed
Untyped constants
```
const ui = 12345 // kind: integer
const uf = 3.141592 // kind: floating-point
```

Typed constants - still use the constant type system but their precision is restricted

```
const ti int = 12345 // type: int
const tf float64 = 3.141592 // type: float64
```

Compiler complains that constant 1000 overflows unit8 cause we are using value that's larger than its type which is unit8

```
const myUnit8 unit8 = 1000
const myUint8 = 1000 // won't complain
```

## Kind promotion
Constant arithmetic supports different kinds. Kind Promotion is used to determine kind in these scenarios.

```
// Variable answer will of type float64.

var answer = 3 * 0.333 // kindFloat(3) * kindFloat(0.333)

// Constant third will be of kind floating point.

const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

// Constant zero will be of kind integer.

const zero = 1 / 3 // KindInt(1) / KindInt(3)

```
## Iota

Iota only work in const block like below

```
const (
    A1 = iota // 0
    B1 = iota // 1
    C1 = iota // 2
)

const (
    A2 = iota // 0
    B2 // 1
    C2 // 2
)

const (
    A3 = iota + 1 // 1
    B3 // 2
    C3 // 3
)

const (
    A4 = 1 << iota // 1
    B4 // 2
    C4 // 4
    
    // ... 8, 16, 32
)
```

Also checkout https://yourbasic.org/golang/iota/



Btw not just import, const, you can use type block too but it's not common

```
type (
    bill struct{ ... }
    ann struct { ... }
)
```

## Bad practices
Bad example 
```
type handle int

func Foo(h handle) {
    fmt.Println(h)
}

func main() {
    var h handle
    Foo(h)
}
```

Don't do this

```
type color int

const (
    Red color = 10
    Blue color = 20
)
```

Do this

```
type color int

const (
    Red = 10
    Blue = 20
)
```