package strings

import (
	"encoding/json"

	l "fase-4-hf-order/external/logger"
)

func MarshalString(s interface{}) string {
	if s == nil {
		return ""
	}

	o, err := json.Marshal(s)
	if err != nil {
		l.Errorf("error in MarshalString client ", " | ", err)
		return ""
	}

	return string(o)
}
