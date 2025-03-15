package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

func (e *notFoundError) Error() string {
	return e.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "ID is required"}
	}

	if id != "123" {
		return &notFoundError{Message: "Data not found"}
	}

	return nil

}

func main() {
	err := SaveData("123", nil)
	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error", validationError.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error", notFoundError.Error())
		// } else {
		// 	fmt.Println("Unknown error", err.Error())
		// }

		switch err.(type) {
		case *validationError:
			fmt.Println("validation error", err.Error())
		case *notFoundError:
			fmt.Println("not found error", err.Error())
		default:
			fmt.Println("Unknown error", err.Error())
		}

	} else {
		fmt.Println("Data saved successfully")
	}
}
