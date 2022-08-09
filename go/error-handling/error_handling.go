package erratum

import "errors"

func Use(opener ResourceOpener, input string) (err error) {
	res, err := tryOpen(opener)
	if err != nil {
		return
	}
	defer func() { _ = res.Close() }()
	defer func() {
		if r := recover(); r != nil {
			if frobError, ok := r.(FrobError); ok {
				res.Defrob(frobError.defrobTag)
			}
			err = errors.New("meh")
			return
		}
	}()

	res.Frob(input)

	return nil
}

func tryOpen(opener ResourceOpener) (res Resource, err error) {
	run := true
	for run {
		res, err = opener()
		if _, ok := err.(TransientError); !ok {
			run = false
		}
	}
	return
}
