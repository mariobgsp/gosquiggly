package squiggly

import (
	"encoding/json"
	"log"

	"github.com/mariobgsp/gosquiggly/funcs"
)

func Filter(jsonString string, fields string) (b []byte, error error) {
	var result map[string]interface{}
	w, d := funcs.SeparateFilter(fields)

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		log.Panicf("error %s", err)
		return nil, err
	}

	// filtering #1
	wtd := funcs.AcquireWithoutDots(w, result)
	// filtering #2
	wd := funcs.AcquireWithDots(d, result)

	resultFiltered := funcs.MergeMapInterface(wtd, wd)
	return resultFiltered, nil
}
