package utils

import "fmt"

// MapAliasString ...
func MapAliasString(alias, src string) string {
	return fmt.Sprintf("%s -> %s", alias, src)
}
