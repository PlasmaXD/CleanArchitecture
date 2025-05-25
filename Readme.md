# Go Clean Architecture サンプルアプリ (TODO) with Delete

以下は、Gin を用いたシンプルな TODO 管理アプリに「削除機能 (Delete)」を加えた例です。

```
.
├── cmd
│   └── app
│       └── main.go
└── internal
    ├── domain
    │   ├── todo.go
    │   └── todo_repository.go
    ├── usecase
    │   └── todo_usecase.go
    ├── interface
    │   └── web
    │       ├── handler
    │       │   └── todo_handler.go
    │       └── router.go
    └── infrastructure
        └── persistence
            └── todo_repository.go
```

---



### 起動手順
※Go 1.18 以上がインストールされている前提

```bash
# 1. リポジトリをクローン（なければローカルにディレクトリを用意）
git clone https://github.com/yourname/clean-arch-go-app.git
cd clean-arch-go-app

# 2. Go モジュールを初期化／依存取得
go mod tidy

# 3. アプリを「起動」または「ビルド＆実行」
#   a) そのまま起動する場合
go run ./cmd/app

#   b) ビルドして実行ファイルを作る場合
go build -o todo-app ./cmd/app
./todo-app
```

起動に成功するとデフォルトで `:8080` 番号の HTTP サーバが立ち上がります。

---

### 動作確認例

1. **TODO 作成**

```bash
curl -X POST http://localhost:8080/api/todos \
     -H "Content-Type: application/json" \
     -d '{"title":"最初のタスク"}'
```

2. **TODO 一覧取得**

```bash
curl http://localhost:8080/api/todos
```

以上で、Clean Architecture 構成の TODO アプリが動作するはずです。
もしポート番号やログ出力、DB 接続先を変えたい場合は `cmd/app/main.go` や環境変数で調整してください。
