package dog

import "strings"

func Bark(name string) string {
	return strings.ToUpper(name) + " does Woof! Woof!"
}
