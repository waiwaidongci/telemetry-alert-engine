package devicequery

func ShouldRetry(status int) bool { return status >= 400 }
