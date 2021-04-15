# yatter-backend-go
<!--
[![CircleCI](https://cci.dmm.com/gh/bootcamp-2020/yatter-backend-go.svg?style=svg)](https://cci.dmm.com/gh/bootcamp-2020/yatter-backend-go)
-->

## develop
### requirements
* docker
* GHE access rights

### start
```
git clone https://git.dmm.com/bootcamp-2020/yatter-backend-go
cd ./yatter-backend-go
docker-compose up -d
```

### modify DB model
Do NOT use any migration tools, for simplification.
`ddl/ddl.sql` will be automatically executed per `docker-compose up` (regenerate container).

Using `app/domain/object` even as DTO for DB.


## structure
[Architecture Design](https://git.dmm.com/cto-tech//wiki/Architecture-design) をベースに短期間で実装するために簡素化

```
.
├── app      ----> application core codes
│   ├── app      ----> collection of dependency injected
│   ├── config   ----> config
│   ├── domain   ----> domain layer, core business logics
│   ├── handler  ----> (interface layer & application layer), request handlers
│   └── dao      ----> (infrastructure layer), implementation of domain/repository
│
└── ddl      ----> DB definition master
```

※ 本来は usecase レイヤーが必要となるが、handler レイヤーへ合体

