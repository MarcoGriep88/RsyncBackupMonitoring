package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"net/http"
	"encoding/json"
	"bytes"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	openFile := os.Args[1:]
	file, err := os.Open(openFile[0])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	files := ""
	created := ""
	backup_date := ""

	var allTransferedFiles []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		ln := scanner.Text()
		x := strings.Split(ln, " ")

		if (backup_date == "") {
			backup_date = x[0]
			fmt.Println("Backup Date: " + backup_date)
		}

		if (strings.Contains(ln, "Number of files")) {
			files = x[6]
			fmt.Println("Files backuped: " + files)
			continue
		}

		if (strings.Contains(ln, "Number of created files")) {
			created = x[7]
			fmt.Println("Created Files: " + created)
			continue
		}

		allTransferedFiles = append(allTransferedFiles, x[4])

		
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	reqBody, err := json.Marshal(map[string]string{
		"backupDate": backup_date,
		"numerOfFiles":    files,
		"created": created,
	})

	if err != nil {
		print(err)
	}
	
	resp, err := http.Post("http://api-backup-monitoring/create", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
}