package notificationerrors

func ShouldRetry(kind Kind, attempt int) bool {
	remaining := 3 - attempt
	if remaining <= 0 {
		return false
	}
	if kind == KindMissing {
		return remaining > 0
	}
	return remaining > 0
}
