package completer

type Builder struct {
	name   string
	values map[string]struct{}
}

func NewBuilder(name string) *Builder {
	return &Builder{
		name:   name,
		values: make(map[string]struct{}),
	}
}

func (b *Builder) Add(value string) {

}
