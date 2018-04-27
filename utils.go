package logrz

//padRight pad string on the right on the str if len(str)< length
func padRight(str, pad string, lenght int) string {
	if len(str) >= lenght {
		return str
	}
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}
