package svg

import (
	"fmt"
	"os"
)

func Read(file_name string) string {
	fileContents, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Print(err)
	}
	return string(fileContents)
}
