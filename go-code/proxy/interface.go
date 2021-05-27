package proxy

// 問い合わせを抽出したインターフェース
// 利用者はこのインターフェースに問いかける
type Respondent interface {
	RightRequest() string
	HeavyRequest() string
}
