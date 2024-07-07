package main

import (
    "net/http"
    "fmt"
    "os"
    "path/filepath"
    "encoding/json"
)

var repositoryPaths string = "./repositories"

func stringifyFile(file_path string) (string, error){
    data, err := os.ReadFile(file_path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}




func pingHandler(writer http.ResponseWriter, req *http.Request){
    fmt.Fprintf(writer, "ping")
}




func cloneHandler(writer http.ResponseWriter, req *http.Request){
    fmt.Println("Incoming clone request:")
    for name, values := range req.Header {
        for _, value := range values {
            fmt.Printf("%s: %s\n", name, value)
        }
    }


    repository := req.Header.Get("Repository")
    if repository == "" {
        fmt.Fprintln(writer, http.StatusBadRequest)
        return
    }
    
    repo_filepath := filepath.Join("repositories", repository)
    dirs, files, err := WalkRepo(repo_filepath)
    if err != nil {
        fmt.Fprintln(writer, http.StatusBadRequest)
        return
    }


    writer.Header().Set("Content-Type", "application/json")
    json.NewEncoder(writer).Encode(dirs)
    json.NewEncoder(writer).Encode(files)
    fmt.Println("Sent:", dirs)
}


func cloneFileHandler(writer http.ResponseWriter, req *http.Request){
    fmt.Println("Incoming Clone File Request")
    for name, values := range req.Header {
        for _, value := range values {
            fmt.Printf("%s: %s\n", name, value)
        }
    }
    repository := req.Header.Get("Repository")
    if repository == "" {
        fmt.Fprintln(writer, http.StatusBadRequest)
        return
    }
    fmt.Println(req.Header.Get("Repository"))
    fmt.Fprintln(writer, "die")
    fmt.Println()
}

func server(){
    address := "localhost:8080"
    http.HandleFunc("/ping", pingHandler)
    http.HandleFunc("/clone/structure", cloneHandler)
    http.HandleFunc("/clone/file", cloneHandler)

    fmt.Println("Serving:", address) 
    err := http.ListenAndServe(address, nil)
    if err != nil {
        panic(err)
    }
}
func main(){
    server()
}
