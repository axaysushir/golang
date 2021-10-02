package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	fmt.Println("URLS in golang")

	u, err := url.Parse("http://bing.com/search?q=dotnet")

	if err != nil {
		log.Fatal(err)
	}

	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()

	fmt.Println(u)
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u)

	r := &url.URL{
		Scheme: "https",
		User:   url.UserPassword("user", "password"),
		Host:   "example.com",
		Path:   "foo/bar",
	}
	fmt.Println(r) // https://user:password@example.com/foo/bar

}
