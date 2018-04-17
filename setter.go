package logrz

//SetLevel allow user to personalize logging level
func SetLevel(level Level) {
	printLevel = level
}

//DisableColor allow user to remove log color
func DisableColor(set bool) {
	disableColor = set
}

//ForceFullColor allow user to force color on the entire log
func ForceFullColor(set bool) {
	forceFullColor = set
}

//SetUseFile set file usage for log
func SetUseFile(use, flush bool, set string) {
	useFile = use
	if useFile {
		filePath = set
		ensureFile()
		if flush {
			flushFile()
		}
		startFileLog()
	} else {
		filePath = ""
	}
}
