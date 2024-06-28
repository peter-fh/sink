package sink;


import (
    "cli/state"
)


type Command interface {
    Exec() (string, error)
}

var commandMap = map[string]func(*Sink) Command {
    "": MakeDefaultCommand,
    "init": MakeInitCommand,
    "status": MakeStatusCommand,
    "help": MakeHelpCommand,
}

type Sink struct {
    State *state.State
    Args []string
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
    makeCommand, _ := commandMap[command]
    commandInstance := makeCommand(&sinkInstance)

    return commandInstance, nil 
}

