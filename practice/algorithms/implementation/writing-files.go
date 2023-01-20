package main

import (
	"fmt"
	"os"
)

func writeToFile(bytesChannel chan []byte, doneChannel chan bool, errChannel chan error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	errChannel <- err
	for {
	  select {
		case bytesString := <- bytesChannel:
		  _, err := f.Write(bytesString)
		  errChannel <- err
		case <- doneChannel:
		  return
	  }
	}
}

const filename string = "hello.txt"

func main() {
  inputArray := []string{"hello", " world"}
  bytesChannel, doneChannel, errChannel := make(chan []byte), make(chan bool), make(chan error)
  go writeToFile(bytesChannel, doneChannel, errChannel)
  err := <- errChannel
  if err != nil {
    panic(err)
  }
  for _, b := range inputArray {
    bytesChannel <- []byte(b)
    err := <- errChannel
    if err != nil {
      fmt.Printf("Critical error: %s", err.Error())
      break
    }
  }
  doneChannel <- true
  
  // ...

  fmt.Println("done")
}