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
	args := os.Args[1:]
	file, err := os.Open(args[0])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	jobname := args[1]

	files := ""
	created := ""
	backup_date := ""

	//var allTransferedFiles []string

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

		//allTransferedFiles = append(allTransferedFiles, x[4])

    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	hostname, error := os.Hostname()  
	if error != nil {  
	 panic(error)  
	}  

	reqBody, err := json.Marshal(map[string]string{
		"Hostname": hostname,
		"backupDate": backup_date,
		"numerOfFiles": files,
		"created": created,
		"backupType": jobname,
	})

	if err != nil {
		fmt.Println(err)
	}
	
	resp, err := http.Post("http://192.168.112.205:15000/create", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
	}

	sc := resp.StatusCode

	//defer resp.Body.Close()

	if (sc == 200) {
		fmt.Println("Ok")
		os.Exit(0)
	}

}