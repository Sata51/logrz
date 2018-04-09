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
