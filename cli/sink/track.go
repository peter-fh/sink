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
        s.cpFile(file_path)
    } else {
        s.cpDir(file_path)
    }
}


func (s *Sink) dataPath(tracked_file string) string{
    return filepath.Join(s.State.Path, ".sink", "data", tracked_file)
}

func (s *Sink) repositoryPath(tracked_file string) string {
    return filepath.Join(s.State.Path, tracked_file)
}

func (s *Sink) cpFile(tracked_file string) error {
    copied_file_path := s.dataPath(tracked_file)
    source_file_path := s.repositoryPath(tracked_file)


    source_file, err := os.Create(source_file_path)
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


func (s *Sink) cpDir(tracked_dir string) error {
    copied_dir_path := s.dataPath(tracked_dir)

    err := os.Mkdir(copied_dir_path, 0755)
    return err
}


func MakeTrackCommand(s *Sink) Command {

    return TrackCommand{s}
}

