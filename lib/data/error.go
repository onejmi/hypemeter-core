package data

type ResponseErr struct {
	Status string `json:"status"`
}

func (err ResponseErr) Error() string {
	return err.Status
}
