package sink;

import (
    "fmt"
)
func (sink *Sink) Init() (string, error){
    if !sink.State.SinkInitialized {
        return "", fmt.Errorf("Sink repository already initialized in %s", sink.State.Path)
    }

    return fmt.Sprintf("New repository initialized in %s", sink.State.Path), nil 
}


