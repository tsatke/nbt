package nbt

type tagBase struct {
	name string
}

func (t tagBase) Name() string {
	return t.name
}

func (t *tagBase) SetName(name string) {
	t.name = name
}
