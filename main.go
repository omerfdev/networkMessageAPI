package main

import (
    "fmt"
    "os/exec"
    "strings"
    "net/http"
    "io/ioutil"
)

func scanNetwork() ([]string, error) {
    cmd := exec.Command("nmap", "-sn", "192.168.1.0/24")
    output, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(output), "\n")
    var devices []string
    for _, line := range lines {
        if strings.Contains(line, "Nmap scan report for") {
            parts := strings.Split(line, " ")
            devices = append(devices, parts[len(parts)-1])
        }
    }
    return devices, nil
}

func sendMessage(deviceIP string, message string) error {
    url := fmt.Sprintf("http://%s:8080/sendMessage", deviceIP)
    resp, err := http.Post(url, "text/plain", strings.NewReader(message))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }
    fmt.Printf("Response from %s: %s\n", deviceIP, string(body))
    return nil
}

func main() {
    devices, err := scanNetwork()
    if err != nil {
        fmt.Printf("Error scanning network: %v\n", err)
        return
    }

    message := "Hello, this is a test message!"
    for _, device := range devices {
        fmt.Printf("Sending message to %s\n", device)
        err := sendMessage(device, message)
        if err != nil {
            fmt.Printf("Error sending message to %s: %v\n", device, err)
        }
    }
}
