package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// As of now, main doesn't do anything
// @todo check user's OS and perform actions based on that
func main() {
	Watcher()
}

// Starts watching specified folder with fsnotify
func Watcher() {

	hostev, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = hostev.WatchFlags("/etc/", 2)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go func() {
		for {
			select {
			case ev := <-hostev.Event:
				log.Println("event:", ev)
				WatchTasks()
			case err := <-hostev.Error:
				log.Println("error", err)

			}
		}
	}()

	<-done
	hostev.Close()

}

// Perform tasks when a change is detected in the watched directory
func WatchTasks() {
	binary, lookErr := exec.LookPath("sh") // checks the path of the sh executable
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"sh", "script.sh"} // Name of command must be first argument
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
