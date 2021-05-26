package strategy

// Player: 戦略を実行する人（戦略の振る舞いを活用する人）
type Player struct {
	st Strategy
}

func NewPlayer(st Strategy) *Player {
	p := new(Player)
	p.st = st
	return p
}

// DIを利用したStrategyパターン
// Template Methodパターンの気持ちも入っている
func (p *Player) Do() string {
	return p.st.Prepare() + " -> " + p.st.Execute()
}

// 関数オブジェクトを利用したStrategyパターン
// 振る舞いが単一で済むのならば、この方法が簡易的
func (p *Player) AdLib(st func() string) string {
	return "戦略: " + st()
}

// Strategy: 戦略の振る舞いを定義するインタフェース
type Strategy interface {
	Prepare() string
	Execute() string
}

// CarefulStrategy: 注意深い戦略
type CarefulStrategy struct{}

func NewCarefulStrategy() *CarefulStrategy {
	return new(CarefulStrategy)
}
func (cs *CarefulStrategy) Prepare() string {
	return "宿題は計画的に"
}
func (cs *CarefulStrategy) Execute() string {
	return "初日に終わった"
}

// ShallowStrategy: 浅はかな戦略
type ShallowStrategy struct{}

func NewShallowStrategy() *ShallowStrategy {
	return new(ShallowStrategy)
}
func (ss *ShallowStrategy) Prepare() string {
	return "宿題は最後の一週間に"
}
func (ss *ShallowStrategy) Execute() string {
	return "助けて"
}
