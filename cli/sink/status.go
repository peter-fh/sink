package sink;

import (
    "fmt"
)


type StatusCommand struct {
    sinkInstance *Sink
}


func (s StatusCommand) Exec() (string, error) {
    msg := "\n"

    if !s.sinkInstance.State.SinkInitialized {
        msg += fmt.Sprintln("No tracked repository in current directory.")
        msg += fmt.Sprintln("Run 'sink init' to initialize a repository.")
        return msg, nil
    }
    panic("Status for tracked directory not implemented yet")
    return msg, nil
}

func (s StatusCommand) Log() (bool, string) {
    return false, ""
}

func MakeStatusCommand(s *Sink) Command {
    return StatusCommand{s}
}

