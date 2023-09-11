package utils

import (
	"log"
	"os"
	"strings"
)

// FailOnError logs error and panics if something goes bad
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// BodyFrom parses a message from STDIN args
func BodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}
