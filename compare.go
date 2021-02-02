package nscript

func CompareInt64(op string, a, b int64) bool {
	switch op {
	case "<":
		return a < b
	case "<=":
		return a <= b
	case ">":
		return a > b
	case ">=":
		return a >= b
	case "==", "=":
		return a == b
	case "!=", "<>":
		return a != b
	}
	return false
}
