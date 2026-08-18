package deliveryflow

type Status string

const (
	Queued   Status = "queued"
	Retrying Status = "retrying"
	Sent     Status = "sent"
	Failed   Status = "failed"
)

type Delivery struct {
	ID     string
	Status Status
}

func (s Status) Terminal() bool { return s == Sent || s == Failed }
