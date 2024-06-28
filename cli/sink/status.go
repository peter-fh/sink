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

    return msg, nil
}

func MakeStatusCommand(s *Sink) Command {
    return StatusCommand{s}
}

