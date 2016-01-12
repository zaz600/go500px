# go500px
Package go500px provides a client for using the [500px API](https://github.com/500px/api-documentation) in Go.

This package contains *unstable code*. Work in progress.

### Installation

```
go get github.com/zaz600/go500px
```

### Example

```go
package main

import (
	"log"
	"net/url"

	"github.com/zaz600/go500px"
)

func main() {
	client := go500px.NewClient(nil)
	//add your ConsumerKey (https://500px.com/settings/applications)
	client.ConsumerKey = "ConsumerKey"
	v := url.Values{}
	v.Set("feature", "editors")
	v.Add("only", "Animals")
	p, err := client.Photos.GetStream(v)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v", p.Photos)
}
```

### License

[MIT](http://zaz600.mit-license.org)

