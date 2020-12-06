package student

var (
	// Service var
	Service students = &student{}
)

type students interface{}

type student struct{}
