# Gof Design Pattern

## 前置き

- このリポジトリをまとめている人はあまりプログラミング経験がない（体感、圧縮して5年ぐらい）。
- デザインパターンがしっくりこない、というか考えるたびに憂鬱になるので少しずつ考えをまとめたい。

## アイデア

※語尾にはすべて「らしい、そうだ」などの伝聞・推定を付け加えてください。

- 「GoFのデザインパターンは古い」というのは、熟練のプログラマーにとっては共通認識であると思って良さそう。
  - Java言語よりも古くに生まれたものだから、自然ではある。
- 一方で「いくつかは現在でも生き残っておりちゃんと有用である」というのも共通認識であると思って良さそう。
- パターンの中で、不要と断じられやすいもの。
  - Singleton
  - 継承による差分プログラミングを要求するもの全般？
    - 「継承によって想定されるデメリットは、想定されるメリットよりもはるかに大きい。」
- 「最初からデザインパターンを適用しよう」と頑張るのではなく、「リファクタリングをしているときに適用できそうな部分を見つけたら適用する」ぐらいのスタンスが正しい。
- いくつかのパターンは、言語の機能の中に組み込まれている。
  - 典型はPythonやTypeScriptのジェネレータ（＝イテレータパターン）？
- 「新たに追加すべきパターン」というのもある。
  - Null Object
  - Type Object
  - Dependency Injection
  - Extension Object/Interface

## 学習の戦略

- GoFによる再編にあるように、重要なものから順番に押さえていく。特にCore, Creationalから。
  - Core: Composite, Strategy, State, Command, Iterator, Proxy, Template Method, Facade
  - Creational: Factory, Prototype, Builder, Dependency Injection
- 例は自分で考えてみる、そしてコードを書いてみる。
  - 結局、他の人の感性の元で例示されてもしっくりこない。
- 言語はGoで書いてみて、それでうまく表現できないものは切り捨てる。
  - 継承は忘れる。

## 参考

- [Java言語で学ぶデザインパターン入門](https://www.sbcr.jp/product/4797327030/)
- [fukabori.fm 48.GoFデザインパターンとDI（前編）](https://fukabori.fm/episode/48)
- [fukabori.fm 49.GoFデザインパターンとDI＋リファクタリング（後編）](https://fukabori.fm/episode/49)
- [Design Patterns 15 Years Later: An Interview with Erich Gamma, Richard Helm, and Ralph Johnson](https://www.informit.com/articles/article.aspx?p=1404056)
