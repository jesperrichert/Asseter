package util

import "regexp"

func CheckFile(filePath string) bool {
	matched, err := regexp.MatchString(`^[a-zA-Z0-9_\-.]+$`, filePath)
	if err != nil {
		return false
	}
	return matched
}
