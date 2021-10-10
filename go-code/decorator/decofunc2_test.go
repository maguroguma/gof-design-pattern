package decorator

import "testing"

type DecoFuncType func() string

// Pythonなら @RoundingDecoratorGenerator('<p>', '</p>') のように前置できる関数
func RoundingDecoratorGenerator(prefix, suffix string) func(DecoFuncType) DecoFuncType {
	decorator := func(wrapped DecoFuncType) DecoFuncType {
		decoratedFunc := func() string {
			return prefix + wrapped() + suffix
		}
		return decoratedFunc
	}
	return decorator
}

func TestAnotherDecoratedFunction(t *testing.T) {
	hello := func() string {
		return "hello"
	}
	bye := func() string {
		return "bye"
	}

	actual := RoundingDecoratorGenerator("<p>", "</p>")(hello)()
	expected := "<p>hello</p>"
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}

	actual = RoundingDecoratorGenerator("<h1>", "</h1>")(bye)()
	expected = "<h1>bye</h1>"
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
