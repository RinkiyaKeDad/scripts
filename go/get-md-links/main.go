package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./test.md")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(GetAllLinks(string(input)))
}

func GetAllLinks(markdown string) map[string]string {
	// Holds all the links and their corresponding values
	m := make(map[string]string)

	// Regex to extract link and text attached to link
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	scanner := bufio.NewScanner(strings.NewReader(markdown))
	stop := false
	// Scans line by line
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "```") == true {
			stop = !stop
		}

		if stop == false {
			// Make regex
			matches := re.FindAllStringSubmatch(scanner.Text(), -1)

			// Only apply regex if there are links and the link does not start with #
			if matches != nil {
				if strings.HasPrefix(matches[0][2], "#") == false {
					// fmt.Println(matches[0][2])
					m[matches[0][1]] = matches[0][2]
				}
			}
		}
	}
	return m
}
