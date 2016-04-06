package main

import (
	"net/http"
	"fmt"
	"log"
)

func main()  {
	http.HandleFunc("/",handlerFn)
	log.Fatal( http.ListenAndServe("localhost:8000",nil) )
}

func handlerFn(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}