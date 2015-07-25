package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	Watcher()
}

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

func WatchTasks() {
	binary, lookErr := exec.LookPath("sh")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"sh", "script.sh"}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
