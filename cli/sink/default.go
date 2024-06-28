package sink;

import (
    "fmt"
)

type DefaultCommand struct {
    sinkInstance *Sink
}

func (d DefaultCommand) Exec() (string, error){

    msg := ""
    msg += fmt.Sprintf("Welcome to Sink, called from: %s\n\n", d.sinkInstance.State.Path)
    
    if d.sinkInstance.State.SinkInitialized {
        msg += fmt.Sprintln("You are in a sink tracked repository.")
        msg += fmt.Sprintln("Functionality to actually do source control coming soon!")
        return msg, nil
    }

    msg += fmt.Sprintln("No tracked repository in current directory.")
    msg += fmt.Sprintln("Run 'sink init' to initialize a repository.")
    return msg, nil
}


func MakeDefaultCommand(s *Sink) (Command) {
    return DefaultCommand{s} 
}
