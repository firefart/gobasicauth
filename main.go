package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var usernames, passwords string
	flag.StringVar(&usernames, "u", "", "usernames")
	flag.StringVar(&passwords, "p", "", "passwords")
	flag.Parse()

	if usernames == "" || passwords == "" {
		log.Fatal("please provide all parameters")
	}

	fileUsernames, err := os.Open(usernames)
	if err != nil {
		log.Fatal(err)
	}
	defer fileUsernames.Close()
	filePasswords, err := os.Open(passwords)
	if err != nil {
		log.Fatal(err)
	}
	defer filePasswords.Close()

	scannerUsernames := bufio.NewScanner(fileUsernames)
	for scannerUsernames.Scan() {
		user := scannerUsernames.Text()
		_, err := filePasswords.Seek(0, 0)
		if err != nil {
			log.Fatal(err)
		}
		scannerPasswords := bufio.NewScanner(filePasswords)
		for scannerPasswords.Scan() {
			pass := scannerPasswords.Text()
			x := fmt.Sprintf("%s:%s", user, pass)
			fmt.Println(base64.StdEncoding.EncodeToString([]byte(x)))
		}
		if err := scannerPasswords.Err(); err != nil {
			log.Fatal(err)
		}
	}

	if err := scannerUsernames.Err(); err != nil {
		log.Fatal(err)
	}
}
