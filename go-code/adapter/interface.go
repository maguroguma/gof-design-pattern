package adapter

type Getter interface {
	Get(key string) string
}
