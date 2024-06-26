package sink;

import (
    "fmt"
)
func (sink *Sink) Default() (string, error){

    msg := ""
    msg += fmt.Sprintf("Welcome to Sink, called from: %s\n\n", sink.State.Path)
    
    if sink.State.SinkInitialized {
        msg += fmt.Sprintln("You are in a sink tracked repository.")
        msg += fmt.Sprintln("Functionality to actually do source control coming soon!")
        return msg, nil
    }

    msg += fmt.Sprintln("No tracked repository in current directory.")
    msg += fmt.Sprintln("Run 'sink init' to initialize a repository.")
    return msg, nil
}
