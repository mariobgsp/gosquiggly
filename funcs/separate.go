package funcs

import (
	"strings"
)

func SeparateFilter(fields string) ([]string, []string) {
	sliceFields := strings.Split(fields, ",")
	var withoutDots []string
	var withDots []string

	if len(sliceFields) == 1 {
		//check if dots exist
		if strings.Contains(fields, ".") {
			withDots = append(withDots, fields)
		} else {
			withoutDots = append(withoutDots, fields)
		}
	} else {
		for _, vield := range sliceFields {
			if strings.Contains(vield, ".") {
				withDots = append(withDots, vield)
			} else {
				withoutDots = append(withoutDots, vield)
			}
		}
	}

	return withoutDots, withDots
}
