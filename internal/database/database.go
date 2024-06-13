package database

type Wrapper interface {
	Set(token, link string) error
	Get(token string) (string, error)
	Close()
}

type UniqueError struct {
}

func (err UniqueError) Error() string {
	return "unique error"
}

type NotFoundError struct {
}

func (err NotFoundError) Error() string {
	return "not found error"
}
