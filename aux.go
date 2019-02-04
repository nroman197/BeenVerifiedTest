package main

import (
	"strconv"
	"strings"
	)

// function to check errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// check if string is numeric
func isNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

// check if there are 2 params separated by &
func isRangeFormat(s string) bool {
  err := len(strings.Split(s, "&"))
  return err == 2
}
