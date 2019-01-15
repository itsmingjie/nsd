package main

import "strconv"

type target []string

func newTarget(url, status string) target {
	set := []string{url, status}
	return set
}

func (t target) getStatus() int {
	status, _ := strconv.Atoi(t[1])
	return status
}
