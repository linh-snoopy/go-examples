package usecases

type Item struct {
	Id    int
	Name  string
	Value float64
}

type Logger interface {
	Log(message string) error
}
