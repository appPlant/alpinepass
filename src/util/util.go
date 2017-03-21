package util

import "strings"
import "github.com/appPlant/alpinepass/src/data"

import "os"
import "fmt"

//CheckError throws an exception if an error exists.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Filler is used to replace spaces in strings.
const Filler string = "_"

//CleanString removes whitespace in string.
func CleanString(s string) string {
	return strings.Replace(s, " ", Filler, -1)
}

//ThrowError prints the error information and exits the application.
func ThrowError(message string) {
	//TODO read about Go error handling and use this.
	os.Stderr.WriteString(message)
	fmt.Println()
	os.Exit(-1)
}

//ThrowConfigError prints the error, some information about a Config and exits the application.
func ThrowConfigError(message string, config data.Config) {
	//TODO read about Go error handling and use this.
	os.Stderr.WriteString(message)
	//TODO pint out config
	fmt.Println()
	os.Exit(-1)
}
