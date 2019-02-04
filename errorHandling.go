package main

import (
	"strings"
)

// strcut to store validation errors
type ErrorHandler struct{
  Function string "json:function, omitempty"
  Message string "json:message, omitempty"
}

// function to validate params for searchSongsbyLength 
func checkLengths(lenghts string) []ErrorHandler {
	var errs []ErrorHandler

	// check if there are 2 params
	if !isRangeFormat(lenghts) {
    newError := ErrorHandler{Function: "searchSongsbyLength", Message: "Min and max length required!"}
		errs = append(errs, newError)
	} else
    {
      minMaxLengths := strings.Split(lenghts, "&")
      minLength := minMaxLengths[0]
      maxLength := minMaxLengths[1]
    // check if 2 params are numeric
    if !isNumeric(minLength) || !isNumeric(maxLength) {
      newError := ErrorHandler{Function: "searchSongsbyLength", Message: "Min and max must be numbers!"}
      errs = append(errs, newError)
    }
  }

	return errs
}
