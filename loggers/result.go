package loggers

import "fmt"

func LogResult(result interface{}) {
	fmt.Printf("The result is: %v\n", result)
}
