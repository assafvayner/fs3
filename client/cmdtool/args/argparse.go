package args

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

// parse user command to know which operation to
func ParseArgs() (string, string, error) {
	parser := argparse.NewParser("parser", "Parse user input")

	operation := parser.String("o", "operation", &argparse.Options{Required: true, Help: "Operation to execute"})
	filePath := parser.String("f", "file_name", &argparse.Options{Required: true, Help: "Path to the file"})
	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Fprintln(os.Stderr, parser.Usage(err))
		return "", "", err
	}
	fmt.Println(*operation)
	fmt.Println(*filePath)
	return *operation, *filePath, nil
}
