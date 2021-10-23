package defaults

import (
	"fmt"
	"strconv"
)

// String checks if data is nil, returns value instead
func String(data string, value string) string {
	if data == "" {
		return value
	}

	return data
}

// Int checks if data is nil, returns value instead
func Int(data int, value int) int {
	if data == 0 {
		return value
	}

	return data
}

// Float checks if data is nil, returns value instead
func Float(data float64, value float64) float64 {
	if data == 0 {
		return value
	}

	return data
}

func ForceInt(key interface{}) int {
	isInt := fmt.Sprintf("%v", key)
	v, err := strconv.Atoi(isInt)
	if err != nil {
		return 0
	}
	return v
}

func ForceBool(key interface{}) bool {
	isBool := fmt.Sprintf("%v", key)
	v, err := strconv.ParseBool(isBool)
	if err != nil {
		return false
	}
	return v
}
