package logrz

//SetLevel allow user to personalize logging level
func SetLevel(level Level) {
	printLevel = level
}

//DisableColor allow user to remove log color
func DisableColor(set bool) {
	disableColor = set
}
