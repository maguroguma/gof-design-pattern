package prototype

// 特定の共通メソッドを抽出したClonableなモノ
type ClonableObject interface {
	Use(num int) string
	CreateClone() ClonableObject // Prototypeパターンたらしめる最たるもの
}
