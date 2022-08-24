package retry

// Do Run function until no errors returned or maxRetry reached
// There is no delays between runs
func Do(maxRetry int, fn func(int) error) (err error) {
	retryCount := 0

	for retryCount < maxRetry {
		err = fn(retryCount)
		if err == nil {
			return nil
		}
		retryCount++
	}

	return err
}
