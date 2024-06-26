package sink;


import (
    "cli/state"
)


var commandMap = map[string]func(*Sink)(string, error){
    "": (*Sink).Default,
    "init": (*Sink).Init,
    "status": (*Sink).Status,
    "help": (*Sink).Help,
}

type Sink struct {
    State *state.State
    Command func(*Sink) (string, error)
    Args []string
}

func Initialize(args_without_exe[] string) (*Sink, error){
    sink_state, err := state.GetState()
    if err != nil {
        return new(Sink), err
    }


    command := ""
    if len(args_without_exe) != 0 {
        command = args_without_exe[0]
    }

    function, exists := commandMap[command]
    if !exists {
        function = commandMap["help"]
    }

    args := make([]string, 0)
    if len(args_without_exe) > 1 {
        args = args_without_exe[1:]
    }
    sink := &Sink{State: sink_state, Command: function, Args: args}
    return sink, err
}

