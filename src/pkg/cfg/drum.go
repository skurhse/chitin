package cfg

type Drum interface {
	Config() Config
	Tokens() []*string
}

type DefaultDrum

type Tokens interface {
}

type DefaultDrum struct {
	Config_ Config
	Tokens_ Tokens
}
