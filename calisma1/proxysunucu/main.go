package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// func sunucuClose(w http.ResponseWriter, r * http.Request){
// 	fmt.Println("server kapat")
// 	server:= &http.Server{
// 		Addr: ":8086",
// 	}

// 	hata:=server.Close()
// 	fmt.Println(hata)

// }
func main() {
	cloneUrl, err := url.Parse("http://localhost:8085")
	if err!=nil{
		log.Fatal(err)

	}
	proxy:=httputil.NewSingleHostReverseProxy(cloneUrl)
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w,r)

	})
	// http.HandleFunc("/close",sunucuClose)
	http.ListenAndServe(":8086",nil)
}
