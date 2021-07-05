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
	input, err := ioutil.ReadFile("./test2.md")
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
	re2 := regexp.MustCompile(`\[([^\]]*)\]\:\ ([^ ]*)\w`)

	scanner := bufio.NewScanner(strings.NewReader(markdown))
	stop := false
	// Scans line by line
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "```") {
			stop = !stop
		}

		if !stop {
			// Make regex
			matches := re.FindAllStringSubmatch(scanner.Text(), -1)
			matches2 := re2.FindAllStringSubmatch(scanner.Text(), -1)
			// Only apply regex if there are links and the link does not start with #
			if matches != nil {
				if strings.HasPrefix(matches[0][2], "#") == false {
					// fmt.Println(matches[0][2])
					m[matches[0][1]] = matches[0][2]
				}
			}
			if matches2 != nil {
				if strings.HasPrefix(matches2[0][2], "#") == false {
					// fmt.Println(matches2[0][2])
					m[matches2[0][1]] = matches2[0][2]
				}
			}
		}
	}
	return m
}
