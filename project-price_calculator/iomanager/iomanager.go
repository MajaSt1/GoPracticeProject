package iomanager

type IOManager interface {
	Readlines() ([]string, error)
	WriteResult(data interface{}) error
}