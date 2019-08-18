package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
)

var memes = map[string]string{
	"Xresources": "/.Xresources",
}

func main() {
	args := os.Args
	file, err := findConfig(args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// TODO: backup config file

	txtlines := readByLines(file)

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}

func findConfig(config string) (*os.File, error) {
	if ps, found := memes[config]; found {
		// try to find file
		dir, err := homeDir()
		if err != nil {
			return nil, err
		}

		if _, err := os.Stat(dir + ps); err == nil {
			file, err := os.Open(dir + ps)
			if err != nil {
				return nil, err
			}
			return file, nil
		} else if os.IsNotExist(err) {
			return nil, err
		}
	}
	return nil, nil
}

func readByLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	return txtlines
}

func homeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
