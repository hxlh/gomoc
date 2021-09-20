package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	path := os.Args[1]
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	os.Remove("moc_files")
	os.Mkdir("moc_files", os.ModePerm)
	for i := 0; i < len(files); i++ {
		if strings.Contains(files[i].Name(), ".h") {
			log.Println("Building MOC " + files[i].Name())
			index := strings.LastIndex(files[i].Name(), ".h")
			
			cmd:=exec.Command("moc", path + "/" + files[i].Name(), "-o", "moc_files/moc_"+files[i].Name()[:index] + ".cc")
			
			var out bytes.Buffer
			var stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				log.Println(fmt.Sprint(err) + ": " + stderr.String())
				return
			}
		}
	}
	log.Println("finished")
}
