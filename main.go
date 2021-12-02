package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// AOC contains method for all challenges
type AOC struct{}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "You need to pass the challenge number as an argument\n")
		os.Exit(1)
	}
	dayNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		if os.Args[1] == "gen" {
			GenTodayFiles()
		} else {
			fmt.Fprintf(os.Stderr, "Argument should be \"gen\" or a number\n")
			os.Exit(1)
		}
		return
	}
	if dayNum < 1 || dayNum > 25 {
		fmt.Fprintf(os.Stderr, "Challenge number should be between 1 and 25\n")
		os.Exit(1)
	}

	aoc := AOC{}
	challenge := reflect.ValueOf(aoc).MethodByName("Day" + os.Args[1])
	var emptyMethod reflect.Value
	if challenge == emptyMethod {
		fmt.Fprintf(os.Stderr, "Challenge is not implemented yet\n")
		os.Exit(1)
	}
	challenge.Call(nil)
}

// ReadLines returns slice of the file's lines
func ReadLines(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n"), nil
}

// ReadLinesSep returns slice of the file's lines, using a custom separator
func ReadLinesSep(path string, separator string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.ReplaceAll(string(content), "\r", ""), separator), nil
}

// ReadLinesToInt returns slice of the file's lines converted to integers
func ReadLinesToInt(path string) ([]int, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	var intLines []int
	for _, line := range lines {
		intVal, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		intLines = append(intLines, intVal)
	}
	return intLines, err
}

// GenTodayFiles generates source and test files for today's challenge
func GenTodayFiles() {
	day := time.Now().Format("2")
	sourcePath := fmt.Sprintf("day%s.go", day)
	testPath := fmt.Sprintf("day%s_test.go", day)
	data := map[string]string{
		"day": day,
	}
	CreateFromTemplate("templates/day.go.tpl", sourcePath, data)
	CreateFromTemplate("templates/day_test.go.tpl", testPath, data)
	AddBenchmark(day)
}

// CreateFromTemplate create a file from a template
func CreateFromTemplate(source, dest string, data map[string]string) bool {
	t, err := template.ParseFiles(source)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	_, err = os.Stat(dest)
	if os.IsNotExist(err) {
		sourceFile, err := os.Create(dest)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error creating output file\n")
			return false
		}
		err = t.Execute(sourceFile, data)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error executing template\n")
			return false
		}
		sourceFile.Close()
	} else {
		fmt.Fprintf(os.Stderr, "Skipping: %s already exists\n", dest)
		return false
	}
	return true
}

func AddBenchmark(day string) {
	readme, err := os.OpenFile("README.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer readme.Close()
	if _, err := readme.WriteString(fmt.Sprintf("\n### Day %s\n\n```\n```\n", day)); err != nil {
		log.Println(err)
	}
}
