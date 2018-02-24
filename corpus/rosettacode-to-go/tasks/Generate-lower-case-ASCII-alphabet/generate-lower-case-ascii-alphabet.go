func loweralpha() string {
	p := make([]byte, 26)
	for i := range p {
		p[i] = 'a' + byte(i)
	}
	return string(p)
}

//\Generate-lower-case-ASCII-alphabet\generate-lower-case-ascii-alphabet.go
