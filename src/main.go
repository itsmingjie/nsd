package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// flags
var opPtr = flag.String("op", "", "Operation (add, remove)")
var linkPtr = flag.String("link", "", "Custom URL")
var urlPtr = flag.String("url", "", "Redirection target")
var statusPtr = flag.Int("status", 301, "HTTP status code")
var configName = "_redirect"

func init() {
	// randomization seed for default random url
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func main() {
	flag.Parse()

	if *opPtr == "" {
		fmt.Println("Missing required flag: op")
		os.Exit(-1)
	}

	redirMap := loadAll()

	if *opPtr == "add" {
		var link string

		if *urlPtr == "" {
			fmt.Println("Missing required flag: url")
			os.Exit(-1)
		}

		if *linkPtr == "" {
			for {
				// generate random string until it's new
				link = randStringRunes(6)

				if redirMap[link] == nil {
					break
				}
			}
		} else {
			link = *linkPtr
		}

		add(link, *urlPtr, *statusPtr)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadAll() map[string]target {

	f, err := os.Open(configName)
	var m map[string]target
	m = make(map[string]target)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		c := strings.Fields(s.Text())
		m[c[0]] = newTarget(c[1], c[2])
	}

	return m
}

func add(link string, url string, status int) {
	f, err := os.OpenFile(configName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	appendLine := fmt.Sprintf("%s\t\t%s\t\t%s\r\n", "/"+link, url, strconv.Itoa(status))

	if _, err = f.WriteString(appendLine); err != nil {
		panic(err)
	}
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
