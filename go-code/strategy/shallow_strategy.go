package strategy

// shallowStrategy: 浅はかな戦略
type shallowStrategy struct{}

func NewShallowStrategy() *shallowStrategy {
	return new(shallowStrategy)
}
func (ss *shallowStrategy) prepare() string {
	return "宿題は最後の一週間に"
}
func (ss *shallowStrategy) execute() string {
	return "助けて"
}
