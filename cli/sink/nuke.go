package sink;

import (
    "os"
    "path/filepath"
)
type NukeCommand struct {
    sinkInstance *Sink
}


func (n NukeCommand) Exec() (string, error) {
    args := n.sinkInstance.Args
    working_path := n.sinkInstance.State.Path
    if len(args) == 1 {
        working_path = filepath.Join(working_path, args[0])
    }

    msg := "if this command makes into production I've failed horribly"
    err := nuke(working_path)
    return msg, err
}

func nuke(dir string) error{
    c, err := os.ReadDir(dir)
    if err != nil {
        return err
    }
    for _, entry := range c {
        if entry.Name() == ".sink" {
            sink_path := filepath.Join(dir, entry.Name())
            err := os.RemoveAll(sink_path)
            if err != nil {
                return err
            }
        } else if entry.IsDir() {
            dir := filepath.Join(dir, entry.Name())
            err := nuke(dir)
            if err != nil {
                return err
            }
        }
    }
    return nil
    
}

func (n NukeCommand) Log() (bool, string) {
    return false, ""
}

func MakeNukeCommand(s *Sink) Command {
    return NukeCommand{s}
}
