package utils

func BetweenClose(str string, min, max int) bool {
	sz := len(str)
	return sz >= min && sz <= max
}

func BetweenOpen(str string, min, max int) bool {
	sz := len(str)
	return sz > min && sz < max
}