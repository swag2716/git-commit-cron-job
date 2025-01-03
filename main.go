package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	logFile, err := os.OpenFile("task_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Task started!")
	filePath := "counter.txt"

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println("error reading file:", err)
		return
	}

	counter, err := strconv.Atoi(string(data))

	if err != nil {
		log.Println("error converting array of byte to int:", err)
		return
	}

	counter++

	err = ioutil.WriteFile(filePath, []byte(fmt.Sprintf("%d", counter)), 0644)

	if err != nil {
		log.Println("error writing counter to file:", err)
		return
	}

	err = exec.Command("git", "add", ".").Run()

	if err != nil {
		log.Println("error in staging the file:", err)
		return
	}

	commitmsg := "Increment counter to " + strconv.Itoa(counter)
	err = exec.Command("git", "commit", "-m", commitmsg).Run()

	if err != nil {
		log.Println("error in commiting:", err)
		return
	}
	err = exec.Command("git", "push", "-u", "origin", "main").Run()

	if err != nil {
		log.Println("error in git push:", err)
		return
	}

	log.Println("File updated and changes pushed to GitHub successfully.")
}
