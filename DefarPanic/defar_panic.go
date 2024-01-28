package defarpanic

import "fmt"

func logging() {
	fmt.Println("Done call function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func triggerPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic\n", r)
		}
	}()
	panic("International Panic!")
}

func DefarPanic() {

	// recover - panic
	triggerPanic()

}
