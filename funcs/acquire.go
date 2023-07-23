package funcs

import (
	"encoding/json"
	"strings"
)

func AcquireWithoutDots(fields []string, jsonStr map[string]interface{}) []map[string]interface{} {
	var collectedSlices []map[string]interface{}
	newMap := make(map[string]interface{})

	for _, filter := range fields {
		for key, value := range jsonStr {
			if strings.EqualFold(key, filter) {
				newMap[key] = value
			}
		}

	}

	collectedSlices = append(collectedSlices, newMap)
	return collectedSlices

}

func AcquireWithDots(fields []string, jsonStr map[string]interface{}) []map[string]interface{} {
	var collectedSlices []map[string]interface{}

	newMap1 := make(map[string]interface{})
	finalMap := make(map[string]map[string]any)
	sFinalMap := make(map[string]interface{})

	for _, filter := range fields {
		for key, value := range jsonStr {
			// first filter
			newFilter := filter[:strings.IndexByte(filter, '.')]
			// second filter
			newFilter2 := filter[strings.IndexByte(filter, '.'):]
			// get location
			newFiltered := strings.ReplaceAll(newFilter2, ".", "")

			if strings.EqualFold(key, newFilter) {
				// check if exist
				boo := CheckOnMap(key, finalMap)
				// marshal into new Map
				b, _ := json.Marshal(value)

				json.Unmarshal(b, &newMap1)
				// TODO create new map
				if !boo {
					finalMap[key] = make(map[string]any)
				}

				for k, v := range newMap1 {
					if strings.EqualFold(k, newFiltered) {
						finalMap[key][k] = v
						break
					}
				}
				break
			}
		}
	}

	for i, j := range finalMap {
		sFinalMap[i] = j
		collectedSlices = append(collectedSlices, sFinalMap)
	}

	return collectedSlices
}

func CheckOnMap(key string, mapSample map[string]map[string]any) bool {
	check := false
	for k := range mapSample {
		if k == key {
			check = true
		}
	}
	return check
}
