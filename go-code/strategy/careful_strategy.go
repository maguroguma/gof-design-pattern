package strategy

// carefulStrategy: 注意深い戦略
type carefulStrategy struct{}

func NewCarefulStrategy() *carefulStrategy {
	return new(carefulStrategy)
}
func (cs *carefulStrategy) prepare() string {
	return "宿題は計画的に"
}
func (cs *carefulStrategy) execute() string {
	return "初日に終わった"
}
