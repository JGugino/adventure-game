package engine

import "fmt"

func LogInfo(message string) {
	fmt.Printf("INFO: %s\n", message)
}

func LogError(message string) {
	fmt.Printf("ERROR: %s\n", message)
}

func LogWarn(message string) {
	fmt.Printf("WARN: %s\n", message)
}
