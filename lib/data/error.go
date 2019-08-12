package data

type ResponseErr struct {
	Status string
}

func (err ResponseErr) Error() string {
	return err.Status
}
