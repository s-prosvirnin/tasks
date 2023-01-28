package main

import "errors"

func main() {
	println(bar())
}

func bar() (r int) {
	defer func() {
		r += 4
		if err := recover(); err != nil {
			r += 8
		}
	}()

	var f func()
	defer f()
	f = func() {
		r += 2
	}

	return 1
}

func makeRecoveryError(recoveryMessage interface{}) error {
	var err error
	switch x := recoveryMessage.(type) {
	case string:
		err = errors.New("recovered panic! " + x)
	case error:
		err = errors.New("recovered panic!" + x.Error())
	default:
		err = errors.New("recovered panic! unknown")
	}

	return err
}
