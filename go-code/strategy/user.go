package strategy

// player: 戦略を実行する人（戦略の振る舞いを活用する人）
type player struct {
	st strategy
}

func NewPlayer(st strategy) *player {
	p := new(player)
	p.st = st
	return p
}

// DIを利用したStrategyパターン
// Template Methodパターンの気持ちも入っている
func (p *player) Do() string {
	return p.st.prepare() + " -> " + p.st.execute()
}

// 関数オブジェクトを利用したStrategyパターン
// 振る舞いが単一で済むのならば、この方法が簡易的
func (p *player) AdLib(st func() string) string {
	return "戦略: " + st()
}
