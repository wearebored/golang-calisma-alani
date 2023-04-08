package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
)

func main()  {
	fmt.Print("hello")
	e := echo.New()

	// create the reverse proxy
	url, _ := url.Parse("https://jsonplaceholder.typicode.com")
	proxy := httputil.NewSingleHostReverseProxy(url)
	// fmt.Println(e)
	// fmt.Println("")
	// fmt.Printf("%v",proxy)
	e.Start(":2957")
}