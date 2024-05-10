# yatter-backend-go

## Table of Contents
- [Development Environment](#development-environment)
  - [Requirements](#requirements)
  - [Start](#start)
  - [Shutdown](#shutdown)
  - [Log](#log)
  - [Hot Reload](#hot-reload)
  - [Swagger UI](#swagger-ui)
  - [DB](#db)
- [Code](#code)
  - [Architecture](#architecture)
  - [Module Dependency](#module-dependency)
  - [Library](#library)
  - [Utilities](#utilities)

## Development Environment
開発環境を docker-compose で構築しています。

### Requirements
- Go
- docker / docker-compose

### Start
```
docker compose up -d
```

### Shutdown
```
docker compose down
```

### Log
```
# ログの確認
docker compose logs

# ストリーミング
docker compose logs -f

# webサーバonly
docker compose logs web
docker compose logs -f web
```

### Hot Reload
[air](https://github.com/cosmtrek/air)によるホットリロードをサポートしており、コードを編集・保存すると自動で反映されます。
読み込まれない場合は`docker compose restart`を実行してください。

### Swagger UI

API 仕様を Swagger UI で確認できます。
開発環境を立ち上げ、Web ブラウザで[localhost:8081](http://localhost:8081)にアクセスしてください。

#### Test
各 API 定義の"Try it out"から API の動作確認を行うことができます。

#### Authentication
鍵マークのついたエンドポイントは認証付きエンドポイントです。

`Authentication`というHTTPヘッダに`username ${ユーザー名}`を指定する単純な仕様です。

動作確認の際には画面上部の"Authorize"からヘッダの値の設定を行ってください。

### DB
マイグレーションツールの用意はありません。
初回起動時に`ddl/`以下にある SQL ファイルが実行されます。
再読み込みの際は`.data/mysql/`を削除し、DB を初期化してください。
```
docker compose down  # 開発環境が稼働中なら止める
rm -rfd .data/mysql  # `.data/mysql/`以下を削除
docker compose up -d # 再起動
```

## Code
### Architecture
```
.
├── app      ----> application core codes
│   ├── config   ----> config
│   ├── domain   ----> domain layer, core business logics
│   ├── handler  ----> (interface layer & application layer), request handlers
│   └── dao      ----> (infrastructure layer), implementation of domain/repository
│
└── ddl      ----> DB definition master
```

#### config
サーバーの設定をまとめたパッケージです。DB や listen するポートなどの設定を取得する API がまとめてあります。

#### domain
アプリケーションのモデルを扱う domain 層のパッケージです。

##### domain/object
ドメインに登場するデータ・モノの表現やその振る舞いを記述するパッケージです。
今回は簡単のため DTO の役割も兼ねています。

##### domain/repository
domain/object で扱う Entity の永続化に関する処理を抽象化し、インターフェースとして定義するパッケージです。
具体的な実装は dao パッケージで行います。

#### handler
HTTP リクエストのハンドラを実装するパッケージです。
リクエストからパラメータを読み取り、エンドポイントに応じた処理を行ってレスポンスを返します。
機能の提供のために必要になる各種処理の実装は別のパッケージに切り分け、handler は入出力に注力するケースも多いですが、今回は簡単のため統合しています。

#### dao
domain/repository に対する実装を提供するパッケージです。
DB など外部モジュールへアクセスし、データの保存・取得・更新などの処理を実装します。

### Module Dependency

![モジュールの依存関係](doc/module_dependency.png)

### Library
- HTTP
  - chi（[ドキュメント](https://pkg.go.dev/github.com/go-chi/chi/v5)）
- DB
  - sqlx（[ドキュメント](https://pkg.go.dev/github.com/jmoiron/sqlx)）

### Utilities
このテンプレートでは実装をサポートするユーティリティ関数を提供しています。

#### app/handler/auth
認証付きエンドポイントの実装のためのミドルウェア関数を提供しています。
`chi.Mux#Use`や`chi.Mux#With`を用いて利用できます。

- [chi ドキュメント](https://pkg.go.dev/github.com/go-chi/chi/v5)
ミドルウェアを埋め込んだエンドポイントでは`*http.Request`から`AccountOf`でリクエストと紐づくアカウントを取得できます。
```
// var r *http.Request
account := auth.AccountOf(r)
```
r.Group()により、特定のグループに対してミドルウェアを適用することができます。
グループに対して適用されたミドルウェアは、そのグループに属する全てのエンドポイントに対して適用されます。

```
r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		// リクエストの認証を行う
		r.Use(auth.Middleware(ar))
		h := &handler{ar}
		r.Post("/", h.Create)
	})
```
