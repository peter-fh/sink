package state;

import (
    "os"
    "path/filepath"
    "errors"
)

type State struct {
    Path string
    SinkInitialized bool
    RepositoryName string
    IsLinked bool
    ConfigOptions Config
}


func GetState() (*State, error) {
    state := new(State)
    cwd, err := os.Getwd()

    if err != nil {
        return state, err
    }

    state.Path = cwd
    sinkPath := filepath.Join(cwd, ".sink")
    _, err = os.Stat(sinkPath)
    if errors.Is(err, os.ErrNotExist) {
        state.SinkInitialized = false
    } else if err != nil {
        state.SinkInitialized = false
        return state, err
    } else {
        state.SinkInitialized = true
    }


    return state, nil
}

func (s *State) SinkPath() string {
    return filepath.Join(s.Path, ".sink")
}

func (s *State) ConfigPath() string {
    return filepath.Join(s.SinkPath(), "config.json")
}
