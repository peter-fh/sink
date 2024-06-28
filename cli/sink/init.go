package sink;

import (
    "fmt"
)

type InitCommand struct {
    sinkInstance *Sink
}

func (i InitCommand) Exec() (string, error){
    if i.sinkInstance.State.SinkInitialized {
        return "", fmt.Errorf("Sink repository already initialized in %s", i.sinkInstance.State.Path)
    }

    return fmt.Sprintf("New repository initialized in %s", i.sinkInstance.State.Path), nil 
}

func MakeInitCommand(s *Sink) (Command) {
    return InitCommand{s}
}


