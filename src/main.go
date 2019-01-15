package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// flags
var isRemove = flag.Bool("remove", false, "Perform remove operation")
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

	// pre-run URL validation top stop wasting resources
	if *urlPtr != "" && !validateURL(*urlPtr) {
		fmt.Println("Invalid URL: Please use full URL, including protocols.")
		os.Exit(-1)
	}

	redirMap := loadAll()
	if *isRemove {
		if *linkPtr != "" {

			// if -url flag is present, check if url matches
			if redirMap[*linkPtr] == nil ||
				(*urlPtr != "" &&
					redirMap[*linkPtr] != nil &&
					redirMap[*linkPtr][0] != *urlPtr) {

				fmt.Println("No action performed: No matching entry")
				os.Exit(-1)

			}

			url := redirMap[*linkPtr][0]
			delete(redirMap, *linkPtr)
			replaceAll(redirMap)

			fmt.Printf("Removed %s >> %s", *linkPtr, url)

		} else if *urlPtr != "" {

			for k, v := range redirMap {
				if v[0] == *urlPtr {
					delete(redirMap, k)
					replaceAll(redirMap)

					fmt.Printf("Removed %s >> %s", k, v[0])
					os.Exit(0)
				}
			}

			fmt.Println("No action performed: No matching entry")
			os.Exit(0)

		} else {

			fmt.Println("Missing required flag: url or link")
			os.Exit(-1)

		}
	} else {
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

			if redirMap[link] != nil {
				fmt.Println("No action performed: Duplicate link entry")
				os.Exit(0)
			}
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
		m[c[0][1:]] = newTarget(c[1], c[2])
	}

	return m
}

func replaceAll(m map[string]target) error {

	var s string

	for k, v := range m {
		s += fmt.Sprintf("/%s\t\t%s\t\t%s\r\n", k, v[0], v[1])
	}

	return ioutil.WriteFile(configName, []byte(s), 0600)
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
	} else {
		fmt.Printf("Added %s >> %s", link, url)
	}
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func validateURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}

	return true
}
