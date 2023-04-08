package main

import (
	"context"
	// "fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func sunucu1(w http.ResponseWriter,r *http.Request){

	w.Write([]byte("sunucu açıldı"))
}

func sunucuAdmin(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("admin paneli"))
}

func sunucuPages(w http.ResponseWriter, r * http.Request){
	w.Write([]byte("sayfalar"))
}



func main()  {
	
	http.HandleFunc("/",sunucu1)
	http.HandleFunc("/admin",sunucuAdmin)
	http.HandleFunc("/sayfalar",sunucuPages)
	
	shutdownChan := make(chan os.Signal, 1)
signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
    shutdownChan <- syscall.SIGTERM
    w.WriteHeader(http.StatusOK)
})
server := &http.Server{Addr: ":8084", Handler: nil}
go func() {
    if err := server.ListenAndServe(); err != nil {
        log.Printf("Server error: %s", err)
    }
}()
<-shutdownChan // wait for shutdown signal
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
if err := server.Shutdown(ctx); err != nil {
    log.Fatalf("Server shutdown error: %s", err)
}
	http.ListenAndServe(":8085",nil)
	
}