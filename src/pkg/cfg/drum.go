package cfg

type Drum interface {
	Config() Config
	Tokens() []*string
}

type DefaultDrum struct {
	Config_ Config
	Tokens_ []*string
}
