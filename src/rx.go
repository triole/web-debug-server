package main

import (
	"regexp"
)

func rxFindSubMatch(rx string, sub, str string) (r string) {
	temp, _ := regexp.Compile(rx)
	matches := temp.FindStringSubmatch(str)
	names := temp.SubexpNames()
	for i, match := range matches {
		if i != 0 && names[i] == sub {
			r = match
		}
	}
	return
}
