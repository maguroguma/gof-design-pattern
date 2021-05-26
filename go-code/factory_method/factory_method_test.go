package factory_method

import (
	"reflect"
	"testing"
)

func TestFacotryMethod(t *testing.T) {
	// 据え置き機のファクトリ
	sgf := NewStationaryGameFactory()
	fm := NewFactoryManager(sgf)

	p1 := fm.FactoryMethod()
	p2 := fm.FactoryMethod()
	p3 := fm.FactoryMethod()

	expected := []string{
		"[Start] Stationary game No.1",
		"[Start] Stationary game No.2",
		"[Start] Stationary game No.3",
	}

	// 保持しているプロダクトの振る舞いから確認
	behaviors := []string{p1.use(), p2.use(), p3.use()}
	if !reflect.DeepEqual(behaviors, expected) {
		t.Errorf("Error")
	}

	// データベースを確認
	dbResults := []string{}
	for _, p := range sgf.record {
		dbResults = append(dbResults, p.use())
	}
	if !reflect.DeepEqual(dbResults, expected) {
		t.Errorf("Error")
	}

	// ポータブル機のファクトリ（変数が使い回せる）
	pgf := NewPortableGameFactory()
	fm = NewFactoryManager(pgf)

	p1 = fm.FactoryMethod()
	p2 = fm.FactoryMethod()
	p3 = fm.FactoryMethod()

	expected = []string{
		"[Start] Portable game No.1",
		"[Start] Portable game No.2",
		"[Start] Portable game No.3",
	}

	// 保持しているプロダクトの振る舞いから確認
	behaviors = []string{p1.use(), p2.use(), p3.use()}
	if !reflect.DeepEqual(behaviors, expected) {
		t.Errorf("Error")
	}

	// データベースを確認
	dbResults = []string{}
	for _, p := range pgf.record {
		dbResults = append(dbResults, p.use())
	}
	if !reflect.DeepEqual(dbResults, expected) {
		t.Errorf("Error")
	}
}
