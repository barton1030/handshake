package Interface

type Topic interface {
	Name() string
	MinConcurrency() int
	MaxConcurrency() int
	FuseSalt() int
}
