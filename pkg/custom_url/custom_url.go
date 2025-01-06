package custom_url

func QueryOrEmpty(s string) string {
	if len(s) == 0 {
		return s
	}
	return "?" + s
}
