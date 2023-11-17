package stk

type CoreDrum interface {
	Drum
	JumpBeat() JumpCoreBeat
	PostgresBeat() PostgresCoreBeat
	ClusterBeat() ClusterCoreBeat
}
