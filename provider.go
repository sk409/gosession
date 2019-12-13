package gosession

type Provider interface {
	Start() (Session, error)
	Stop(string) error
	Get(string) (Session, error)
	F() []string
}
