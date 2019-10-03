package main

import (
    "fmt"
    
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Server created successfully")
    })

    log.Fatal(http.ListenAndServe(":3062", nil)) //3062 is my roll number :-P

}