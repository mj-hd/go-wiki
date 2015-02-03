# go-wiki
タイトルは仮称。Markdownを使った、golangによるWikiシステム。

# インストール
config/config.goを作成し、config.go.exampleを参考にして設定する。
設定が終わったら、go build go-wiki.goを実行する。

init.dを使用するシステムの場合は、config/init_scriptをシンボリックリンクとして/etc/init.d/以下に貼る。
そして、/etc/init.d/go-wiki startを実行することで、起動できる。
