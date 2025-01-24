package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type Config struct {
	Server string
	Port   int
}

var (
	config Config
	wg     sync.WaitGroup
	mutex  sync.Mutex
)

func readConfigFile(filename string) {
	defer wg.Done()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
}

func main() {
	wg.Add(1)
	go readConfigFile("config.json")
	wg.Wait()

	mutex.Lock()
	fmt.Println("Server:", config.Server)
	fmt.Println("Port:", config.Port)
	mutex.Unlock()
}
