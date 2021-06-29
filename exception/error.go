package exception

import "log"

func FatalErrorIfNeeded(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}
