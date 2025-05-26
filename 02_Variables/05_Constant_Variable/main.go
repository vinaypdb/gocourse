package main

import "fmt"

func main() {
	const appName = "GoTutorial"
	const version = 1.0
	const isReleased = true

	fmt.Println("App Name:", appName)
	fmt.Println("Version:", version)
	fmt.Println("Released:", isReleased)

	// appName = "NewApp" // ❌ This will cause a compile-time error
}
