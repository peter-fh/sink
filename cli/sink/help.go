package sink;


type HelpCommand struct {
    sinkInstance *Sink
}


func (h HelpCommand) Exec() (string, error) {
    panic("Help() not implemented yet")
}

func (h HelpCommand) Log() (bool, string) {
    panic("log for help not implemented yet")
}

func MakeHelpCommand(s *Sink) Command {
    return HelpCommand{s}
}


