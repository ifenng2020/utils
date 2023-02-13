package tools

func SubString(s string, start, end int) string {
	ns := []rune(s)
	if start < 0 || start > len(ns) {
		return ""
	}
	if end > len(ns) {
		end = len(ns)
	}
	return string(ns[start:end])
}
