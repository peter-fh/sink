package main

import (
    "net/http"
    "fmt"
)

var repositoryPaths string = "./repositories"


func server(){
    address := "localhost:8080"
    http.HandleFunc("/clone/structure", cloneHandler)
    http.HandleFunc("/clone/file", cloneFileHandler)

    fmt.Println("Serving:", address) 
    err := http.ListenAndServe(address, nil)
    if err != nil {
        panic(err)
    }
}
func main(){
    server()
}
