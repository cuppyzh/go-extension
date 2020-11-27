package strinformatter

import (
	"errors"
	"regexp"
)

// Compose is a function mimic-ing string concationation in other language
func Compose(strs ...string) (string, error) {
	pattern := regexp.MustCompile("%v")
	vCounter := len(pattern.FindAll([]byte(strs[0]), -1))

	if vCounter != len(strs)-1 {
		return "", errors.New("mismatch parameter")
	}

	var results = []string{}
	i := 0
	loc := pattern.FindIndex([]byte(strs[0]))
	stringToCheck := strs[0]

	for {
		left := stringToCheck[:loc[0]]
		right := stringToCheck[loc[1]:]

		results = append(results, left+strs[1+i])
		stringToCheck = right
		loc = pattern.FindIndex([]byte(stringToCheck))
		i++

		if i == vCounter {
			break
		}
	}

	result := ""

	for _, str := range results {
		result += str
	}

	return result, nil
}
