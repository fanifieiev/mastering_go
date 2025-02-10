package main

import (
	"log"
	"os"
	"path/filepath"
)

// func init() {
// 	logFile, err := os.OpenFile("which.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatal("Failed to open log file", err)
// 	}
// 	defer logFile.Close()

// 	log.SetOutput(logFile)
// }

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Println("Missing arguments!")
		return
	}

	file := args[1]
	pathList := filepath.SplitList(os.Getenv("PATH"))
	found := false
	for _, dir := range pathList {
		fullPath := filepath.Join(dir, file)
		fileInfo, err := os.Stat(fullPath)
		//path exists?
		if err == nil {
			mode := fileInfo.Mode()
			//is executable?
			if mode.IsRegular() {
				if !found {
					found = true
				}
				if mode&0111 != 0 {
					log.Println(fullPath)
				}
			}
		}
	}
	if !found {
		log.Println("Program not found!")
	}
}
