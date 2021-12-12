# dotfile-cli
dotfiles のインストールスクリプトを Go で CLI にしてみたもの。  
go の promptui を使ってみたかったため作成。   
\+ Bazel を使ってビルド環境を作成してみる。
## BUILD 方法
```
# 依存パッケージのリポジトリ情報を WORKSPACE に書き出す
$ bazel run //:gazelle -- update-repos -from_file ./go.mod -to_macro=deps.bzl%go_dependencies

# build
$ bazel build //...

# 実行
# bazel run //:dotfile-cli
```