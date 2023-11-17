package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Params struct {
	file   string
	level  string
	output string
}

func AppParameters() Params {
	var file string
	var level string
	var output string

	flag.StringVar(&file, "file", "", "Log-file name.")
	flag.StringVar(&level, "level", "info", "Log level to analyze.")
	flag.StringVar(&output, "output", "", "Output file name.")

	flag.Parse()

	if file == "" {
		file = os.Getenv("LOG_ANALYZER_FILE")
	}
	if level == "" {
		level = os.Getenv("LOG_ANALYZER_LEVEL")
		if level == "" {
			level = "info"
		}
	}
	if output == "" {
		output = os.Getenv("LOG_ANALYZER_OUTPUT")
	}

	params := Params{
		file:   file,
		level:  level,
		output: output,
	}

	return params
}

func LogData(filename string) ([][]string, error) {
	var res [][]string
	file, err := os.OpenFile(filename, os.O_RDONLY, 0o644)
	defer func() {
		err = file.Close()
	}()
	if err != nil {
		return res, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ";")
		res = append(res, line)
	}

	if err = scanner.Err(); err != nil {
		return res, err
	}

	return res, err
}

func LogStats(logData [][]string, level string) []string {
	var res []string
	res = append(res, fmt.Sprintf("Total lines number: %d", len(logData)))

	count := 0
	for _, l := range logData {
		if l[0] == level {
			count++
		}
	}
	res = append(res, fmt.Sprintf("Lines number with level '%s': %d", level, count))

	return res
}

func WriteStats(stats []string, output string) {
	var writer *bufio.Writer
	if output == "" {
		writer = bufio.NewWriter(os.Stdout)
	} else {
		file, err := os.Create(output)
		defer func() {
			if err = file.Close(); err != nil {
				panic(err)
			}
		}()
		if err != nil {
			panic(err)
		}

		writer = bufio.NewWriter(file)
	}

	for _, l := range stats {
		writer.WriteString(l + "\n")
	}
	writer.Flush()
}

func main() {
	param := AppParameters()

	if param.file == "" {
		fmt.Println("You must enter log-file name with parameter -file")
		return
	}

	logData, err := LogData(param.file)
	if err != nil {
		fmt.Println(err)
		return
	}
	stats := LogStats(logData, param.level)

	WriteStats(stats, param.output)
}
