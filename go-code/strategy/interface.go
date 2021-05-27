package strategy

// strategy: 戦略の振る舞いを定義するインタフェース
type strategy interface {
	prepare() string
	execute() string
}
