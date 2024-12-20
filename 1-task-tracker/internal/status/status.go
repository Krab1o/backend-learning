package status

type Status int

const (
	ToDo Status = iota
	InProgress
	Done
)

func (status Status) String() string {
	switch status {
	case ToDo:
		return "to do"
	case InProgress:
		return "in progress"
	case Done:
		return "done"
	default:
		return "unknown"
	}
}