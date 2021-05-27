package composite

// 木構造の各ノードは抽象化される
// 抽象化することで、実際には異なるもの同士を含む混合物が形成できる
type Entry interface {
	GetName() string
	GetSize() int
	Add(Entry) (error, Entry)
	PrintListDefault() string
	PrintList(prefix string) string
	String() string
}
