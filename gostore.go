package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const filename = "./data.json"

func main() {
	interruptHandler()

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var store map[string]string
	if len(data) > 0 {
		json.Unmarshal(data, &store)
	} else {
		store = make(map[string]string)
	}

	for {
		fmt.Printf("> ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		tokens := strings.Split(text, " ")
		op, operands := tokens[0], tokens[1]

		if op == "get" {
			value, present := store[operands]

			if present {
				fmt.Println(value)
			} else {
				fmt.Println("nil")
			}
		}

		if op == "set" {
			writeFile, openErr := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0755)
			if openErr != nil {
				panic(openErr)
			}
			defer writeFile.Close()

			setOperands := strings.Split(operands, "=")
			key, value := setOperands[0], setOperands[1]

			store[key] = value

			serializedData, _ := json.Marshal(store)

			_, writeErr := writeFile.Write(serializedData)
			if writeErr != nil {
				panic(writeErr)
			}

			writeFile.Sync()
		}

		log.Println("Current Store:", store)
	}
}

func interruptHandler() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		<-sigs
		fmt.Println("")
		fmt.Println("Exiting")
		os.Exit(0)
	}()
}
