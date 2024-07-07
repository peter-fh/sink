package sink;

import (
    "fmt"
    "net/http"
    "io"
    "bufio"
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


func cloneStructureRequest(repository_name string) (string, error) {
    clone_address := SinkAddress + "/clone/structure"
    req, err := http.NewRequest("GET", clone_address, nil)

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

    msg := ""
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        msg += scanner.Text()
    }

    return msg, nil
}


func cloneRead(resp *http.Response) error {

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    fmt.Println(body)

    return nil
}
