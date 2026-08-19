package notificationerrors

func ShouldRetry(kind Kind, attempt int) bool {
	if kind == KindMissing {
		return false
	}
	return attempt < 3
}
