package util

import "fmt"

func JsonError(message string) string {
	return fmt.Sprint(`"error":"%v"`, message)
}
