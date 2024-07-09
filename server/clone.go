package main


import (
    "path/filepath"
    "io/fs"
    "os"
    "fmt"
    "net/http"
    "encoding/json"
)

func stringifyFile(file_path string) (string, error){
    data, err := os.ReadFile(file_path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}


type RepositoryStructure struct {
    Dirs []string   `json:"dirs"`
    Files []string  `json:"files"`
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
    
    os.Chdir("repositories")
    //repo_filepath := filepath.Join("repositories", repository)
    dirs, files, err := walkRepo(repository)
    if err != nil {
        fmt.Fprintln(writer, http.StatusBadRequest)
        return
    }


    writer.Header().Set("Content-Type", "application/json")
    structure := RepositoryStructure{Dirs: dirs, Files: files}
    json.NewEncoder(writer).Encode(structure)
    fmt.Println("Sent:", dirs)
    fmt.Println()
}


func cloneFileHandler(writer http.ResponseWriter, req *http.Request){
    fmt.Println("Incoming Clone File Request")
    for name, values := range req.Header {
        for _, value := range values {
            fmt.Printf("%s: %s\n", name, value)
        }
    }
    repository := req.Header.Get("Repository")
    relative_path := req.Header.Get("File")
    if relative_path == ""  || repository == ""{
        fmt.Printf("Got bad relative path (%s) or repository (%s)\n", relative_path, repository)
        fmt.Fprintln(writer, http.StatusBadRequest)
        return
    }

    file_path := relative_path
    //file_path := filepath.Join("repositories", repository, relative_path)

    fmt.Println(file_path)
    file_content, err := stringifyFile(file_path)
    if err != nil {
        fmt.Printf("Didn't find file: %s, got error: %v", file_path, err)
        fmt.Fprintln(writer, http.StatusBadRequest)
        return
    }

    fmt.Fprintln(writer, file_content)
    fmt.Println()
}




func walkRepo(repo_path string) ([]string, []string, error) {
    var directories []string
    var files [] string
    err := filepath.Walk(repo_path, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() && info.Name() != "."{
            directories = append(directories, path)
        } else if !info.IsDir() {
            files = append(files, path)
        }

        return nil
    })

    if err != nil {
        return nil, nil, err
    }

    return directories, files, nil
}
