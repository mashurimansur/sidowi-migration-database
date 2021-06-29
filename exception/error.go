package exception

import "log"

// FatalErrorIfNeeded this function
func FatalErrorIfNeeded(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
