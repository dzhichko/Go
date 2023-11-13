package main

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"

	"lab_3.2/internal/cmd"
)

func main() {
	file, err := os.ReadFile("./textfiles/config.yaml")
	if err != nil {
		fmt.Println("Cannot to open file")
	}

	var config cmd.Config
	errp := yaml.Unmarshal(file, &config)
	if errp != nil {
		fmt.Println("Error in parsing yaml file")
	}

	prog := cmd.NewApp()

	for _, fileConfig := range config.Files {
		inStr, err := prog.SearchSubstr(fileConfig.Filename, fileConfig.Substring)
		if err != nil {
			fmt.Println("Error in searching substring")
		} else if inStr == true {
			fmt.Printf("File %s contains substring %s\n", fileConfig.Filename, fileConfig.Substring)
		} else {
			fmt.Printf("File %s doesn't contains substring %s\n", fileConfig.Filename, fileConfig.Substring)
		}
	}
}
