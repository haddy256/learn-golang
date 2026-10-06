package export

// Hello is uppercase → public (other packages can use it)
func Hello() string {
	return "Hi!"
}

// bye is lowercase → private (hidden outside this package)
func bye() string {
	return "Bye!"
}
