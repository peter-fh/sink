package sink;

import (
)


type LinkCommand struct {
    sinkInstance *Sink
}


func (c LinkCommand) Exec() (string, error) {
    msg := ""
    return msg, nil
}

func (c LinkCommand) Log() (bool, string) {
    return false, ""
}

func MakeLinkCommand(s *Sink) Command {
    return LinkCommand{s}
}

