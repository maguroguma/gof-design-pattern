# Command

- Stateのように、「コマンド」という形無いものをクラスと見立てている。
  - **なのでStateと同じくちょっと複雑で難しいし、ピンポイントなパターンだと思う。**
- マクロコマンドもコマンドのインターフェースを実装しておくことで、「コマンドの集合」も1つのコマンドとして扱うことが出来る。
  - これに関してはCompositeパターンが応用されている。
- 結城先生の本のサンプルはちょっと複雑なので、[monochromeganeさんのサンプル](https://github.com/monochromegane/go_design_pattern/tree/master/command)を写経してみる。
