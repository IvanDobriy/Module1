package user

type Module1 interface {
	Version() string
}

type MyModule struct{}

func (m MyModule) Version() string {
	return "1.0.0"
}
