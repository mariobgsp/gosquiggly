package funcs

import (
	"encoding/json"
)

func MergeMapInterface(wtd []map[string]interface{}, wd []map[string]interface{}) []byte {
	wtdIntf := wtd[0]
	wdIntf := wd[0]

	for k, v := range wdIntf {
		wtdIntf[k] = v
	}
	b, _ := json.Marshal(wtdIntf)
	return b

}

func FilterSingleMap(wtd []map[string]interface{}, wd []map[string]interface{}) []byte {
	wtdIntf := wtd
	wdIntf := wd

	if len(wtdIntf) == 0 {
		wtdIntf = wdIntf
	}

	b, _ := json.Marshal(wtdIntf)
	return b

}
