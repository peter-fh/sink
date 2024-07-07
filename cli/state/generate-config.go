package state 

import (
    "fmt"
    "path/filepath"
    "encoding/json"
    "os"
    "io"
)


var currentVersion int = 1
type Config struct {
    Version     int     `json:"version"`
    Name        string  `json:"name"`
    IsLinked    bool    `json:"is_linked"` 
}


func (s *State) UpdateConfig() error {
    return s.writeConfig()
}

func (s *State) InitializeConfig() error {
    name := filepath.Base(s.Path)
    s.ConfigOptions = Config{
        Version: currentVersion,
        Name: name,
        IsLinked: false,
    }
    return s.writeConfig() 
}


func (s *State) readConfig() error {
    if !s.SinkInitialized {
        return fmt.Errorf("Shit programmer alert\nReading config while sink state is uninitialized")
    }

    config_path := filepath.Join(s.SinkPath(), "config.json")
    json_config, err := os.Open(config_path)

    if err != nil {
        return err
    }

    byte_config, err := io.ReadAll(json_config)

    if err != nil {
        return err
    }
    config := Config{}
    err = json.Unmarshal(byte_config, &config)

    if err != nil {
        return err
    }

    s.ConfigOptions = config

    return nil
}

func (s *State) writeConfig() error {
    if !s.SinkInitialized {
        return fmt.Errorf("Shit programmer alert\nWriting config in uninitialized state")
    }

    json_bytes, err := json.Marshal(s.ConfigOptions)
    if err != nil {
        return err
    }
    
    config_file_path := s.ConfigPath()

    return os.WriteFile(config_file_path, json_bytes, os.ModePerm)

}
