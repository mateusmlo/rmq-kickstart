package utils

import "log"

// FailOnError logs error and panics if something goes bad
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
