package sink;

import (
    "fmt"
    "path/filepath"
    "os"
    "time"
)

type InitCommand struct {
    sinkInstance *Sink
}

func (i InitCommand) Exec() (string, error){
    path := i.sinkInstance.State.Path
    if i.sinkInstance.State.SinkInitialized {
        err := fmt.Errorf("Sink repository already initialized in %s", path)
        return "", err
    }
    if len(i.sinkInstance.Args) > 1 {
        return "", fmt.Errorf("Too many arguments to init")
    }

    err := i.sinkInstance.createDirectory()
    if err != nil {
        return "", err
    }

    msg := fmt.Sprintf("New repository initialized in %s", path)
    return msg, nil 
}

func (i InitCommand) Log() (bool, string) {
    now := time.Now()
    command := "init"
    if len(i.sinkInstance.Args) == 1 {
        command += " " + i.sinkInstance.Args[0]
    }
    return true, fmt.Sprintf("<%s> <%s>", now, command)
}

func MakeInitCommand(s *Sink) (Command) {
    return InitCommand{s}
}


func (s *Sink) createDirectory() error {
    dir_path := s.State.Path
    if len(s.Args) == 1 {
        dir_path = filepath.Join(dir_path, s.Args[0])
        err := os.Mkdir(dir_path, 0755)
        if err != nil {
            return err
        }
        err = os.Chdir(dir_path)

        if err != nil {
            return err
        }
    }

    sink_path := filepath.Join(dir_path, ".sink")
    
    err := os.Mkdir(sink_path, 0755)

    if err != nil {
        return err
    }

    err = os.Chdir(sink_path)

    return nil
}
