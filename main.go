package main

import (
	"./fsnotify/"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	hostev, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-hostev.Event:
				log.Println("event:", ev)

				// Define the players
				cmd := "/usr/bin/killall"
				arg := []string{"mDNSResponder"}

				// Game on.
				err := exec.Command(cmd, arg[0]).Start()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				fmt.Println("Sucess !")

			case err := <-hostev.Error:
				log.Println("error", err)

			}
		}
	}()

	err = hostev.WatchFlags("/etc/", 2)
	if err != nil {
		log.Fatal(err)
	}

	<-done

	hostev.Close()

}
