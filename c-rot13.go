package main

type rot13 struct {
	input string
}

func (this rot13) encrypt() (string, error) {
	c := caesar{input: this.input, n: 13}
	return c.encrypt()
}

func (this rot13) decrypt() (string, error) {
	c := caesar{input: this.input, n: 13}
	return c.encrypt()
}
