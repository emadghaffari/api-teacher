package course

var (
	// Service var
	Service courses = &course{}
)

type courses interface{}

type course struct{}
