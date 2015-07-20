package phineas

type Resource interface {
	Build() error
	BuildPath() string
}
