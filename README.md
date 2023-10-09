# go-layered-architecture

## 概要
レイヤードアーキテクチャについての実装です。
DDDの他のアーキテクチャとの比較を把握するためにrepoistoryに依存した実装にしています。

## 利用方法
```sh
docker-compose up -d # コンテナを立ち上げた段階でテスト用のDB作成とテストデータの生成を行っています
go run cmd/user/main.go
curl -i http://localhost:8080/user/1
```