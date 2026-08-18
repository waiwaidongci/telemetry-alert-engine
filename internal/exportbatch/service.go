package exportbatch

func ValidateExport(t *Tracker, invalid bool) error {
	r := t.Acquire()
	if invalid {
		return ErrBusiness
	}
	r.Close()
	return nil
}
