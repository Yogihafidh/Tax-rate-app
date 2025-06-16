package iomanager

type IOManager interface {
	ReadFile() ([]string, error)
	WriteResult(data any) error
}
