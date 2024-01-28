package Error

import "fmt"

// cutome error
type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func saveData(id string) error {
	if id != "james" {
		return &notFoundError{Message: "data not found"}
	}
	return nil
}

func mainError() {
	err := saveData("")

	if err != nil {
		if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("no found error: ", notFoundError.Message)
		} else {
			fmt.Println("unknown error: ", err.Error())
		}
	}
}
