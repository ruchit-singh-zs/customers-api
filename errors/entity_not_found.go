package errors

type NoEntity struct {
}

func (d NoEntity) Error() string {
	return "No entity"
}
