# gqlgen-sample

## 手順

1. PostgreSQLコンテナを起動
    ```shell
     $ docker run --name postgres \
            -e POSTGRES_PASSWORD=password \
            -e POSTGRES_INITDB_ARGS="--encoding=UTF8 --no-locale" \
            -e TZ=Asia/Tokyo \
            -v postgresdb:/var/lib/postgresql/data \
            -p 5432:5432 \
            -d postgres
     ```
2. データベースを作成
    ```shell
    $ docker exec -it postgres /bin/bash
    /# psql -U postgres
    postgres=# create database gqlgen_sample
    ```
3. スキーマからテーブルを生成
    - schema.sqlをインポートする
4. モデルを生成
   以下を実行することで`models/`にファイルが生成される
   ```shell
   $ go generate
   ```
5. gqlgenコマンドを実行
    ```shell
    go run github.com/99designs/gqlgen
    ```
6. サーバーを起動

## 参考記事

- [【GraphQL × Go】gqlgenの基本構成とオーバーフェッチを防ぐmodel resolverの実装](https://tech.layerx.co.jp/entry/2021/10/22/171242)
