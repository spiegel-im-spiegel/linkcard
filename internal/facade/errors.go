package facade

type jsonOutputError struct {
	Payload string
}

func (e *jsonOutputError) Error() string {
	return "verification failed"
}
