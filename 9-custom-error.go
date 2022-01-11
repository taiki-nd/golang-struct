package main

import "fmt"

type UserNotFound struct {
	username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("user not found: %v", e.username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{username: "mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
