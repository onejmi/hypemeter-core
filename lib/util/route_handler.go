package util

type ResponseErr struct {
	Status string `json:"error"`
}

func (err ResponseErr) Error() string {
	return err.Status
}
