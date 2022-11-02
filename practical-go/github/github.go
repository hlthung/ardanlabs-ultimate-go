package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()
	fmt.Println(githubInfo(ctx, "tebeka"))
}

// githubInfo returns name and number of public repos for login
func githubInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	// resp,err := http.Get(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	defer resp.Body.Close()

	// Note: User-Agent is useful when you want to get result in compact form or user-friendly form
	// Checkout jq:
	// https://medium.com/how-tos-for-coders/https-medium-com-how-tos-for-coders-parse-json-data-using-jq-and-curl-from-command-line-5aa8a05cd79b
	// fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	// var r Reply

	var r struct { // anonymous struct
		Name string
		// Public_Repos int
		NumRepos int `json:"public_repos"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Name, r.NumRepos, nil
}

/*
if err != nil {
	log.Fatalf("error: %s", err)
	// log.Printf("error: %s", err) + os.Exit(1)
}
type Reply struct {
	Name string
	// Public_Repos int
	NumRepos int `json:"public_repos"`
}
*/

/* JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int32, int64, int, uint8, ...
array <-> []any ([]interface{}) // cause Json we can have an array of mixed type eg. ["hi", 7, true, null]
object <-> map[string]any, struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON: json.Marshal
*/
