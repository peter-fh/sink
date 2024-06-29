package sink;


import (
    "cli/state"
    "os"
    "fmt"
)


type Sink struct {
    State *state.State
    Args []string
}

type Command interface {
    Exec() (string, error)
    Log() (bool, string)
}

var commandMap = map[string]func(*Sink) Command {
    "": MakeDefaultCommand,
    "init": MakeInitCommand,
    "status": MakeStatusCommand,
    "help": MakeHelpCommand,
    "nuke": MakeNukeCommand,
}


func Initialize(args_without_exe[] string) (Command, error){
    sink_state, err := state.GetState()
    if err != nil {
        return nil, err
    }


    args := make([]string, 0)
    if len(args_without_exe) > 1 {
        args = args_without_exe[1:]
    }

    command := ""
    if len(args_without_exe) != 0 {
        command = args_without_exe[0]
    }

    sinkInstance := Sink{sink_state, args}
    makeCommand, found := commandMap[command]
    if !found {
        makeCommand = MakeHelpCommand
    }
    commandInstance := makeCommand(&sinkInstance)

    return commandInstance, nil 
}

func WriteLog(c Command){

    loggable, msg := c.Log()
    if !loggable {
        return
    }

    
    msg_bytes := []byte(msg)
    fmt.Println("logging:", msg)
    err := os.WriteFile("log", msg_bytes, 0644)
    
    
    if err != nil {
        panic(err)
    }


}

