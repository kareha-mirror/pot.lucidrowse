# ルシドロウズ (Lucidrowse)

生成AIでキャラクタを作り、1日に1回だけ短い言葉でキャラクタの行動を指定して、夢の世界と関わって行く、絵本日記作成ゲームのようなものです。

* デモ: https://lucid.kareha.org/
* 紹介ページ: https://kareha.org/view/lucidrowse/index

# ビルドと実行

フロントエンドはFlutterで、バックエンドはGoで出来てます。
フロントエンドはFlutter Webを想定していて、ネイティブアプリは想定していません。
GoのバックエンドはデータベースにPostgreSQLを使います。
また、Goのバックエンドは生成AIを呼び出しますが、今のところOpenAI APIにだけ対応しています。

## フロントエンド

client ディレクトリで run-local.sh を実行すると、デバッグモードで起動します。
この際設定は env/local.json に従います。

client ディレクトリで build-web.sh を実行すると、リリースモードでビルドします。
この際設定は env/demo.json に従います。
client/build/web をウェブサーバに転送してください。

## バックエンド

server ディレクトリで task run を実行するとバックエンドサーバが起動します。
バックエンドサーバのバイナリをあらかじめ作っておくには task を実行します。
バイナリを実行するには server server.yaml のように引数で設定ファイルを指定します。
設定ファイルは server.yaml.example を元にして書いてください。
