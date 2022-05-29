# Gof Design Pattern

2022-05-08追記: 次第にGoFに限らず、設計全般について学んだこと、を集積する場所になってきた。

## 前置き

- ここに記した知見は、すべて本ドキュメント最下部の「参考、引用元」の資料を起点としています。
- このリポジトリをまとめている人はあまりプログラミング経験がない（体感、圧縮して5年ぐらい）。
- デザインパターンがしっくりこない、というか考えるたびに憂鬱になるので少しずつ考えをまとめたい。
- 他サイトの説明は、どうしても「1を聞いて10わかってね」という説明に感じてしまう。
  - めんどいからしょうがないことだと思うけど。
  - なので、自分の考えのもと知識を整理する。

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
  - 可視範囲（パッケージ内のみで参照可能かどうか）はかなり雑かもしれない。
- コードを書いてみて思ったが、デザインパターンを意識する際は、振る舞いに注意すべくTDDに忠実になると良いかもしれない。
  - 公開すべきメソッドや関数を強く意識しながら練習できるため。
- 結城先生の本でも丁寧に説明されるような、各クラスやインターフェースの役割の名前を厳密に覚えるようなことは「決してしない」。
  - 受験勉強がしたいわけではなく、実際に（例えばリファクタリングなどで）使えることが目的であるため。
  - それよりは「ここは振る舞いを抽出したinterfaceで、それを実装した具象クラス（構造体）がこれで、interfaceを保有する（composeする）ユーザクラス（構造体）はこれで...」というふうに、感覚で理解したい。

## 参考、引用元

- [Java言語で学ぶデザインパターン入門](https://www.sbcr.jp/product/4797327030/)
  - 聖典。
- [fukabori.fm 48.GoFデザインパターンとDI（前編）](https://fukabori.fm/episode/48)
- [fukabori.fm 49.GoFデザインパターンとDI＋リファクタリング（後編）](https://fukabori.fm/episode/49)
  - めちゃくちゃ勉強になった。twadaさんのこれまでの経験を元にスラスラと語ってくれるので非常に聞きやすく、また学びも多かった。
- [Design Patterns 15 Years Later: An Interview with Erich Gamma, Richard Helm, and Ralph Johnson](https://www.informit.com/articles/article.aspx?p=1404056)
  - 上記のpodcastで紹介された記事。これもさらに改定する必要があるとは思うが、デザインパターンを学ぶ上で時間配分を考えられるので、ありがたい。
- [monochromeganeさんのGoによるデザインパターン](https://github.com/monochromegane/go_design_pattern)
  - 特に、Goにおける構造体・インターフェースの埋め込みによる実装パターンの勉強にもなる。
- [TECHSCOREさんのデザインパターン](https://www.techscore.com/tech/DesignPattern/index.html/)
  - 別の例でイメージをふくらませるために活用させていただいた。
- [GoFデザインパターンチートシート](https://qiita.com/tanakahisateru/items/df03d2558f9499d1a64a)
  - 示唆に富んでいる。
- [フロントエンドのデザインパターン](https://zenn.dev/morinokami/books/learning-patterns-1)
  - オリジンは[この英語のページ](https://www.patterns.dev/)
