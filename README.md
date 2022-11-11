# round-robin-time-table

総当たり式組み合わせのタイムテーブルを生成する CLI ツール環境

## Docker 環境構築

### 必要ツール

- docker
- docker-compose
- envsubst

### 環境変数設定

```sh
# e.g.
echo "export GITHUB_TOKEN=token" >> ~/.bashrc
echo "export GITHUB_USER_NAME=name" >> ~/.bashrc
echo "export GITHUB_EMAIL=email" >> ~/.bashrc
source ~/.bashrc
```

### コンテナ起動

以下コマンドでコンテナを立ち上げ、`/repo`を開く

```sh
  make up
```

## 使い方

### 人物リストファイルを用意

任意の場所にテキストファイルで人物名リストを作成する。人物は３人以上である必要があります。

```persons.txt
一郎
二郎
三郎
```

### 実行

`-o`オプションで出力する CSV ファイルのパスを指定できます。

```
go run main.go -i 人物リストファイルのパス -o 出力ファイルのパス
```

### 組み合わせを Miro で付箋化

以下の手順のようにスプレッドシートを仲介してコピペすることで、各セルを付箋として貼り付けることができます。

1. 出力した CSV ファイルをスプレッドシートにインポート
2. スプレッドシートでセルを範囲選択し、コピーする
3. Miro でペーストする
   https://help.miro.com/hc/article_attachments/4809489511442/cells_as_sticky_notes_in_Miro.gif
