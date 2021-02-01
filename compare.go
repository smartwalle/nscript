package nscript

import "github.com/smartwalle/conv4go"

func CompareInt64(op string, a, b interface{}) bool {
	var aValue = conv4go.Int64(a)
	var bValue = conv4go.Int64(b)
	switch op {
	case "<":
		return aValue < bValue
	case "<=":
		return aValue <= bValue
	case ">":
		return aValue > bValue
	case ">=":
		return aValue >= bValue
	case "==":
		return aValue == bValue
	case "!=":
		return aValue != bValue
	}
	return false
}
