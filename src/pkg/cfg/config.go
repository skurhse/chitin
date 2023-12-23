package cfg

type Config interface {
	Name() string
	Regions() Regions
}

type Regions interface {
	Primary() string
	Secondary() string
}
