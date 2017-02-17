package utils

import "fmt"

// MapAliasString ...
func MapAliasString(alias, src string) string {
	return fmt.Sprintf("%s -> %s", alias, src)
}

// ConfirmAction ...
func ConfirmAction(message string) bool {
	fmt.Printf("%s? [y/n]: ", message)
	var response string
	fmt.Scanf("%s", &response)
	for response != "y" && response != "n" {
		fmt.Printf("Please answer with 'y' or 'n': ")
		fmt.Scanf("%s", response)
	}
	return response == "y"
}
