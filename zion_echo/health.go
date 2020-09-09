package zion_echo

type ReadyCheck struct {
	Name  string
	Ready func() bool
}
