package xmlutil

import (
	"strconv"
)

func ConvertIntToString(v interface{}) {
	switch value := v.(type) {
	case map[string]interface{}:
		for k, v := range value {
			switch va := v.(type) {
			case float64:
				value[k] = strconv.FormatFloat(va, 'f', -1, 64)
			default:
				ConvertIntToString(va)
			}
			//if reflect.TypeOf(v).Kind() == reflect.Float64 {
			//	value[k] = strconv.FormatFloat(v.(float64), 'f', -1, 64)
			//	continue
			//}

		}
	case []interface{}:
		for i, v := range value {
			switch va := v.(type) {
			case float64:
				value[i] = strconv.FormatFloat(va, 'f', -1, 64)
			default:
				ConvertIntToString(va)
			}
			//if reflect.TypeOf(v).Kind() == reflect.Float64 {
			//	value[i] = strconv.FormatFloat(v.(float64), 'f', -1, 64)
			//	continue
			//}
			//ConvertIntToString(v)
		}

	}
}
