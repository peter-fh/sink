package main

import (
    "fmt"
    "os"
    "cli/sink"
)


func main(){
    args := os.Args[1:]
    
    command_instance, err := sink.Initialize(args)

    if err != nil {
        panic(err)
    }

    msg, err := command_instance.Exec()
    if err != nil {
        fmt.Println("Fatal error occurred:")
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(msg)
    sink.WriteLog(command_instance)

}
