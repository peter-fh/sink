package sink;

import (
)


type TemplateCommand struct {
    sinkInstance *Sink
}


func (c TemplateCommand) Exec() (string, error) {
    msg := ""
    return msg, nil
}

func (c TemplateCommand) Log() (bool, string) {
    return false, ""
}

func MakeTemplateCommand(s *Sink) Command {
    return TemplateCommand{s}
}

