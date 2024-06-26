package main

import (
    "fmt"
    "os"
    "cli/sink"
)


func main(){
    args := os.Args[1:]
    fmt.Println("Welcome to codesync!")
    fmt.Println()

    if len(args) > 0 {
        fmt.Println(args)
    }
    
    sink, err := sink.Initialize(args)
    if err != nil {
        panic(err)
    }


    fmt.Printf("Directory in state is: %s\n", sink.State.Path)
    fmt.Printf("Initialized sink directory: %t\n", sink.State.SinkInitialized)
    sink.Command(sink)

}
