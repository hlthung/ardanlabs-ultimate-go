# JSON - The Fine Print: 
- Webinar with Miki Tebeka (Aug 24, 2022)
- https://github.com/tebeka/talks/tree/master/json-the-fine-print

## Chunked Transfer Encoding

### Background
- https://bunny.net/academy/http/what-is-chunked-encoding/

### Chunking Example - HTTP Handler

```
func handler(w http.ResponseWriter, r *http.Request) {
    fluser, ok := w.(http.FLusher)
    if !ok {
        http/Error(w, "no streaming support", http.StatusInternalServerError)
        return
    }

    enc := json.NewEncoder(w)
    for i := 0; i < 10; i++ {
        if err := enc.ERncode(Point{i,i}); err != nil {
            // cant set error
            log.Printf("encoding: %s", err)
            return
        }
        flusher.Flush()
    }
}
```

## Zero vs Missing values
```
type Payment struct {
    Time time.Time      `json:"time"`
    From string         `json:"form"`
    To string           `json:"to"`
    Amount float64      `json:"amount"`
}
```
If we run below code, we will get `{Time:Time.Data(1, time.January, 1, 0, 0, 0, 0, time.UTC)...}` which is actually a zero Time data
```
data := []byte(`{
    "from": "User1",
    "to": "User2",
    "amount": 123.45
}`)

var p Payment
if err := json.Unmarshal(data, &p); err != nil {
    return err
}
fmt.Printf("%#v\n",p)
```
We want to know if it's empty because sender didnt send or they actually send us a zero value. Practically, it's actually fine (like, just don't care, honestly), but if we want to distinguish, there are a  few ways:

1. Update Payment struct's Time as pointer
```
type Payment struct {
    Time *time.Time     `json:"time"`
    From string         `json:"form"`
    To string           `json:"to"`
    Amount float64      `json:"amount"`
}
```
When we run it, we will see `{Time:<nil>...}`

Consequences: Now you will need handle it carefully, as you might get panic.

2. Use a library, eg. [mapstructure](https://pkg.go.dev/github.com/mitchellh/mapstructure)

```
data := []byte(`{
    "from": "User1",
    "to": "User2",
    "amount": 123.45
}`)

var m map[string]interface{}
if err := json.Unmarshal(data, &m); err != nil {
    return err
}

var p Payment
if err := mapstructure.Decode(m, &p); err != nil {
    return err
}

if _, ok := m["time"]; !ok {
    p.Time = time.Now()
}
fmt.Printf("%#v\n",p)
```

3. Set time when receiving - Recommended, most of the time we see industry doing this also
```
data := []byte(`{
    "from": "User1",
    "to": "User2",
    "amount": 123.45
}`)

p := Payment{
    Time: time.Now()
}
if err := json.Unmarshal(data, &p); err != nil { // gonna overwrite 
    return err
}
fmt.Printf("%#v\n",p)
```

## Validate
Note: Valid JSON != Valid Data, remember to do validation still.

Standard:
```
func (loc Location) Validate() error {
    if loc.Lat < -90 || loc.Lat > 90 {
        return fmt.Errorf("invalid latitude: %#v\n", loc.Lat)
    }

    if loc.Lng < -180 || loc.Lng > 180 {
        return fmt.Errorf("invalid latitude: %#v\n", loc.Lng)
    }

    return nil
}
```
```
data := []byte(`{"lat": 132.50941, "lng": 34.9440}`)

var loc Location
if err := json.Unmarshal(data, &loc); err != nil {
    return err
}

if err := loc.Validate(); err != nil {
    return err
}

fmt.Println(loc)
```

## The Great Escape

```
s := "<JSON>"
enc := json.NewEncorder(os.Stdout)
enc.Encode(s)
```

will escape but we can do this to switch it off

```
s := "<JSON>"
enc := json.NewEncorder(os.Stdout)
enc.SetEscapeHTML(false)
enc.Encode(s)
```
Also check out: https://yourbasic.org/golang/multiline-string/

