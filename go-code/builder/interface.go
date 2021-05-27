package builder

// 何らかを作成するために必要な共通の振る舞いを抽出したインターフェース
type builder interface {
	opening() string
	closing() string
}
