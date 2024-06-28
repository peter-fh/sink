package sink;


type HelpCommand struct {
    sinkInstance *Sink
}


func (h HelpCommand) Exec() (string, error) {
    panic("Help() not implemented yet")
}

func MakeHelpCommand(s *Sink) Command {
    return HelpCommand{s}
}


