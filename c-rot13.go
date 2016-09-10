package main

type rot13 struct {
	input string
}

func (r rot13) encrypt() (string, error) {
	c := caesar{input: r.input, n: 13}
	return c.encrypt()
}

func (r rot13) decrypt() (string, error) {
	c := caesar{input: r.input, n: 13}
	return c.encrypt()
}
