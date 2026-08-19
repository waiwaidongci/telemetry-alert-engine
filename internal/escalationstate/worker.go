package escalationstate

type Worker struct {
	machine Machine
}

func NewWorker(machine Machine) *Worker {
	return &Worker{machine: machine}
}

func (w *Worker) Recover(incident Incident, deliver func() error) (Incident, error) {
	retrying, err := w.machine.Move(incident, StateRetrying)
	if err != nil {
		return incident, err
	}
	if err := deliver(); err != nil {
		failed, transitionErr := w.machine.Move(retrying, StateFailed)
		if transitionErr != nil {
			return retrying, transitionErr
		}
		return failed, err
	}
	return retrying, nil
}
