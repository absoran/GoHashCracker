package internal

import (
	"fmt"
	"os"
	"time"

	"github.com/absoran/goproject/models"
)

// CalculationTime function takes timestamp at begining and end of the scope. Take time between 2 value and print that time value for calculation time.
func CalculationTime(start time.Time, description string) {
	elapsed := time.Since(start)
	fmt.Printf("%s Total Time:%s", description, elapsed)
}

// PrintResult function takes output object as parameter and prints its values.
func PrintResult(result models.Output) {
	fmt.Printf("hash cracked with in word list with %d elements, at index : %d ; \nWord : %s \nHash : %s\n", result.WordlistSize, result.FoundAt, result.Word+result.Rules, result.Hash)
}

// Parse rule function parse rules taken from web and converts it to hard-coded rule flag value.
func ParseRule(rule string) string {
	switch rule {
	case "No Rule":
		{
			return ""
		}
	case "Make first letter great":
		{
			return "%u"
		}
	case "Add 1 to end of word":
		{
			return "1"
		}
	case "Add 12 to end of word":
		{
			return "12"
		}
	case "Add 123 to end of word":
		{
			return "123"
		}
	}
	return ""
}
func Exit(err error) {
	fmt.Fprintf(os.Stderr, "error : %v\n", err)
	os.Exit(1)
}
func CheckError(err error) {
	if err != nil {
		Exit(err)
	}
}
