package exportbatch

func ValidateExport(t *Tracker, invalid bool) error {
	r := t.Acquire()
	defer r.Close()
	if invalid {
		return ErrBusiness
	}
	return nil
}
