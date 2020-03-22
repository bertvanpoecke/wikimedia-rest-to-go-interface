# wikimedia-rest-to-go-interface

## Info

Go interface to [Wikimedia REST v1 API](https://en.wikipedia.org/api/rest_v1/?spec)

Supported endpoints so far:

* /page/title/{title}
* /page/summary/{title}
* /page/random/title
* /page/random/summary
* /feed/onthispoint/{type}/{mm}/{dd}


## Usage

The wikimedia package can be used to create an interface to the desired Wikimedia project REST API.

```go
func main() {
	w, err := wikimedia.NewWikimedia("https://en.wikipedia.org")
	if err != nil {
		panic(err)
	}

	summary, err := w.GetPageSummary("Belgium")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", summary)
}
```

## License

This package is licensed under MIT license. See LICENSE for details.

