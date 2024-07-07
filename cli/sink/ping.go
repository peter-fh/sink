package sink;

import (
    "net/http"
    "fmt"
    "bufio"
)


type PingCommand struct {
    sinkInstance *Sink
}


func (c PingCommand) Exec() (string, error) {
    return ping()
}

func (c PingCommand) Log() (bool, string) {
    return false, ""
}




func MakePingCommand(s *Sink) Command {
    return PingCommand{s}
}


func ping() (string, error) {
    formatted_response := ""
    resp, err := http.Get(SinkAddress + "/ping")
    if err != nil {
        return "", err
    }
    formatted_response += fmt.Sprintf("Response Status: %s\n", resp.Status)
   
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        formatted_response += scanner.Text()
    }

    return formatted_response, scanner.Err()
}
