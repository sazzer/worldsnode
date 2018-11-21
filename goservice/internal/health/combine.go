package health

// combineHealth will work out the new healthcheck status based on the current and new values to combine
func combineHealth(old string, new string) string {
	if new == Fail || old == Fail {
		return Fail
	} else if new == Warn || old == Warn {
		return Warn
	} else {
		return Pass
	}
}
