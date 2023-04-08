package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)
  
func main()  {
	mux := http.NewServeMux()
    mux.HandleFunc("/", homePage)
    mux.HandleFunc("/close", closePage)
  
    srv := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }
 
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()
 
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
 
	<-sigChan
 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
 
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}
 
	fmt.Println("Server shutdown successfully")
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	for i := 0; i < 20; i++ {
		fmt.Fprintf(w, "Response %d\n", i)
		w.(http.Flusher).Flush()
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintln(w,"Server Working Normally")
}
func closePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"Server Shutting Down Gracefully")
	pid := os.Getpid()
	fmt.Println(pid)
	process, err := os.FindProcess(pid)
	fmt.Println(process)
	fmt.Println(err)


	if err != nil {
		fmt.Println(err)
		return
	}
	
	if err := process.Signal(os.Interrupt); err != nil {
		fmt.Println(err)
		return
	}
}