package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	jmb := NewJaMailBuilder()
	director := NewDirector(jmb)
	actual := director.Construct("お元気ですか。私は元気です。")
	if actual != "拝啓、\nお元気ですか。私は元気です。\n敬具。" {
		t.Errorf("Error")
	}

	emb := NewEnMailBuilder()
	director = NewDirector(emb)
	actual = director.Construct("How are you? I'm fine.")
	if actual != "Hello,\nHow are you? I'm fine.\nBest Regards." {
		t.Errorf("Error")
	}
}
