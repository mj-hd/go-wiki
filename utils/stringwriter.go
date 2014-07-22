package utils

type StringWriter struct {
	String string
}

func (this *StringWriter) Write(bytes []byte) (int, error) {
	this.String += string(bytes)
	return len(bytes), nil
}
