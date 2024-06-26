package state;

import (
    "os"
    "path/filepath"
    "errors"
)

type State struct {
    Path string
    SinkInitialized bool
    SinkFile string
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
    } else if err == nil {
        state.SinkInitialized = true
        state.SinkFile = sinkPath
    } else {
        return state, err
    }

    return state, nil
}
