package squiggly

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/mariobgsp/gosquiggly/funcs"
)

func TestSqu(t *testing.T) {

	jsonString := ""
	fields := "paymentmethod.bank"

	var result map[string]interface{}
	w, d := funcs.SeparateFilter(fields)

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		log.Panicf("error %s", err)
		return
	}

	// filtering #1
	wtd := funcs.AcquireWithoutDots(w, result)
	// filtering #2
	wd := funcs.AcquireWithDots(d, result)

	if len(wtd) >= 1 && len(wd) >= 1 {
		resultFiltered := funcs.MergeMapInterface(wtd, wd)
		log.Print(string(resultFiltered))
	} else if len(wtd) == 0 || len(wd) == 0 {
		resultFiltered := funcs.FilterSingleMap(wtd, wd)
		log.Print(string(resultFiltered))
	}

}
