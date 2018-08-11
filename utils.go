package logrz

//padRight pad string on the right on the str if len(str)< length
func padRight(str, pad string, length int) string {
	if len(str) >= length {
		return str
	}
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}
