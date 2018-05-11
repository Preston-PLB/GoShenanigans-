package main

import (
	"net/http"
	"fmt"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Hey there</h1>")
}

func main(){
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}