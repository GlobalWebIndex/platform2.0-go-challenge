package helpers

func StringTernary(statement bool, a, b string) string {
	if statement {
		return a
	}
	return b
}
