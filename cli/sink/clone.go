package sink

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
)


type CloneCommand struct {
    sinkInstance *Sink
}


func (c CloneCommand) Exec() (string, error) {
    args := c.sinkInstance.Args
    if len(args) == 0 {
        return "", fmt.Errorf("No repository specified for cloning")
    } else if len(args) > 1 {
        return "", fmt.Errorf("Too many arguments to clone command")
    }

    repository_name := args[0]
    resp_body, err := cloneStructureRequest(repository_name)
    if err != nil {
        return "", err
    }
    fmt.Println(resp_body)
    return "Cloned "+ repository_name, nil
}

func (c CloneCommand) Log() (bool, string) {
    return false, ""
}

func MakeCloneCommand(s *Sink) Command {
    return CloneCommand{s}
}


type RepositoryStructure struct {
    Dirs []string   `json:"dirs"`
    Files []string  `json:"files"`
}


func cloneStructureRequest(repository_name string) (string, error) {
    clone_structure_address := SinkAddress + "/clone/structure"
    req, err := http.NewRequest("GET", clone_structure_address, nil)

    if err != nil {
        return "", err
    }

    agent := fmt.Sprintf("SinkCLI/%s", SinkVersion)
    req.Header.Set("User-Agent", agent)
    req.Header.Set("Accept", "application/json")
    req.Header.Set("Repository", repository_name)

    client := &http.Client{}
    resp, err := client.Do(req)


    if err != nil {
        return "", err
    }

    structure := RepositoryStructure{}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    json.Unmarshal(body, &structure)
    fmt.Println("dirs:", structure.Dirs)
    fmt.Println("files:", structure.Files)

    err = clone(repository_name, structure)
    if err != nil {
        return "", err
    }
    return "cloned the repo", nil
}


func clone(repo string, structure RepositoryStructure) error {
    clone_file_address := SinkAddress + "/clone/file"
    for _, dir := range structure.Dirs {
        err := os.Mkdir(dir, 0755)
        if err != nil {
            return fmt.Errorf("Error creating directory structure: %v", err) 
        }
    }


    agent := fmt.Sprintf("SinkCLI/%s", SinkVersion)
    client := &http.Client{}
    for _, file := range structure.Files {
        req, err := http.NewRequest("GET", clone_file_address, nil)
        if err != nil {
            return err
        }
        req.Header.Set("User-Agent", agent)
        req.Header.Set("Accept", "text/plain")
        req.Header.Set("Repository", repo)
        req.Header.Set("File", file)
        resp, err := client.Do(req)
        if err != nil {
            return err
        }
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            return err
        }
        err = os.WriteFile(file, body, 0755)
        if err != nil {
            return err
        }

    }

    return nil
}


