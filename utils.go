package gologrz

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

//padLeft pad string on the left on the str if len(str)< length
func padLeft(str, pad string, lenght int) string {
	if len(str) >= lenght {
		return str
	}
	for {
		str = pad + str
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}
