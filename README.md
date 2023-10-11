# go-layered-architecture

## 概要
レイヤードアーキテクチャについての実装です。
「[ドメイン駆動設計 モデリング/実装ガイド](https://booth.pm/ja/items/1835632)」という本を参考にさせていただきました

## 利用方法
```sh
docker-compose up -d # コンテナを立ち上げた段階でテスト用のDB作成とテストデータの生成を行っています
go run cmd/user/main.go
curl -i http://localhost:8080/user/1
```
