package Error

import (
	"errors"
	"fmt"
)

func CheckError(value int) (int, error) {
	if value == 0 {
		return 0, errors.New("value is zero")
	} else {
		return value, nil
	}
}

func MainError() {
	result, error := CheckError(0)
	if error == nil {
		fmt.Println("result: ", result)
	} else {
		fmt.Println("Error", error.Error())
	}
}
