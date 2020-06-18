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
	"io/ioutil"
	"strconv"
)

type BackupResult struct {
	Id int64
	Hostname string
}

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
	
	if (len(args) < 3) {
		fmt.Println("Invalid count of Start Parameters")
		fmt.Println("main.go [LogfilePath] [Web-API Adress] [Job-Name]")
	}

	api_adress := args[1]
	jobname := args[2]

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
	
	resp, err := http.Post(api_adress+ "/create", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
	}

	sc := resp.StatusCode

	defer resp.Body.Close()

	if (sc == 200) {
		fmt.Println("Ok")

		var r BackupResult

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		
		json.Unmarshal([]byte(bodyString), &r)
		fmt.Println(r.Hostname)
		backup_id := r.Id

		for i:=0; i < len(allTransferedFiles); i++ {
			fmt.Println(allTransferedFiles[i])
			string_backup_id := strconv.FormatInt(int64(backup_id), 10)
			reqBody, err := json.Marshal(map[string]string{
				"fk_backupId": string_backup_id,
				"file": allTransferedFiles[i],
			})

			fmt.Println(string(reqBody))
			
			resp, err := http.Post(api_adress+ "/file", "application/json", bytes.NewBuffer(reqBody))
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				fmt.Println("OK")
			}

			if resp.StatusCode == 201 {
				fmt.Println("OK")
			}
		}

		os.Exit(0)
	}

}
