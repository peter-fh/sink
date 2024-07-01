package sink;

import (
    "os"
    "path/filepath"
    "io"
)


type TrackCommand struct {
    sinkInstance *Sink
}


func (c TrackCommand) Exec() (string, error) {
    msg := ""
    return msg, nil
}

func (c TrackCommand) Log() (bool, string) {
    return false, ""
}


func (s *Sink) track(file_path string) {
    info, err := os.Stat(file_path)

    if err != nil {
        panic(err)
    }
    if !info.IsDir() {
    }
}

func (s *Sink) cpTrackedFile(tracked_file_path string, tracked_file_info os.FileInfo) error {
    sink_path := s.State.Path
    sink_data_path := filepath.Join(sink_path, ".sink", "data")
    copied_file_path := filepath.Join(sink_data_path, tracked_file_info.Name())

    source_file, err := os.Open(tracked_file_path)
    if err != nil {
        return err
    }
    destination_file, err := os.Create(copied_file_path)
    if err != nil {
        return err
    }
    _, err = io.Copy(source_file, destination_file)
    if err != nil {
        return err
    }
    err = destination_file.Sync()
    if err != nil {
        return err
    }
    return nil
}


func (s *Sink) cpDir(tracked_dir_path string, tracked_dir_info os.FileInfo) error {
    panic("cpDir not implemented yet")
    sink_path := s.State.Path
    sink_data_path := filepath.Join(sink_path, ".sink", "data")
    copied_dir_path := filepath.Join(sink_data_path, tracked_dir_info.Name())

    source_file, err := os.Open(tracked_file_path)
    if err != nil {
        return err
    }
    destination_file, err := os.Create(copied_file_path)
    if err != nil {
        return err
    }
    _, err = io.Copy(source_file, destination_file)
    if err != nil {
        return err
    }
    err = destination_file.Sync()
    if err != nil {
        return err
    }
    return nil
}


func MakeTrackCommand(s *Sink) Command {

    return TrackCommand{s}
}

