package greeting

import "fmt"

func Hello(name string) string {
	if name == "" {
		return "Hello, Guest"
	}
	return fmt.Sprintf("Hello, %s", name)
}

func Hi(name string) string {
	if name == "" {
		return "Hello, Guest"
	}
	return fmt.Sprintf("Hi, %s", name)
}
