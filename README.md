# Http Client
- Has **Get, Post, Put, Patch, Delete, Head** and **Options** methods

## Simple Example

```go
package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	c "github.com/ncostamagna/go_http_client/client"
)

func main() {
	client := c.New(nil, "https://base-url", 5000*time.Millisecond, true)
	u := url.URL{}
	u.Path += "/my-path"
	q := u.Query()

	// query string. for example 'page'
	q.Set("page", strconv.Itoa(2))
	u.RawQuery = q.Encode()

	reps := client.Get(u.String())

	if reps.Err != nil {
		fmt.Println(reps.Err)
	} else {
		fmt.Println(reps)
	}
}

```