package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"github.com/recws-org/recws"
)

func startServer(port int) (*exec.Cmd, error) {
	cmd := exec.Command("../server/server", "-port", strconv.Itoa(port))
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func doMain() error {
	port := 8082
	ws := recws.RecConn{
		KeepAliveTimeout: 10 * time.Second,
	}
	fmt.Print("Starting server\n")
	cmd, err := startServer(port)
	if err != nil {
		return err
	}
	fmt.Print("Dialing\n")
	ws.Dial(fmt.Sprintf("ws://localhost:%d", port), nil)
	fmt.Print("Killing server\n")
	err = cmd.Process.Kill()
	if err != nil {
		return err
	}
	fmt.Print("Waiting for server to die\n")
	err = cmd.Wait()
	fmt.Printf("Child process: %s\n", err.Error())
	fmt.Print("Starting server again\n")
	_, err = startServer(port)
	if err != nil {
		return err
	}
	fmt.Print("Reading a message\n")
	_, _, err = ws.ReadMessage()
	fmt.Printf("%s\n", err)
	fmt.Print("Reading a message\n")
	_, _, err = ws.ReadMessage()
	fmt.Printf("%s\n", err)
	return nil
}

func main() {
	err := doMain()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
