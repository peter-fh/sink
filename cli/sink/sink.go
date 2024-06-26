package sink;


import (
    "cli/state"
)


var commandMap = map[string]func(*Sink){
    "": (*Sink).Default,
    "init": (*Sink).Init,
    "status": (*Sink).Status,
    "help": (*Sink).Help,
}

type Sink struct {
    State *state.State
    Command func(*Sink)
    Args []string
}

func Initialize(argsWithoutExe[] string) (*Sink, error){
    sink_state, err := state.GetState()
    if err != nil {
        return new(Sink), err
    }


    command := ""
    if len(argsWithoutExe) != 0 {
        command = argsWithoutExe[0]
    }

    function, exists := commandMap[command]
    if !exists {
        function = commandMap["help"]
    }

    args := make([]string, 0)
    if len(argsWithoutExe) > 1 {
        args = argsWithoutExe[1:]
    }
    sink := &Sink{State: sink_state, Command: function, Args: args}
    return sink, err
}

