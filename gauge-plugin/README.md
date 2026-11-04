# Play CDC (beta)

## Install

[Releases](https://github.com/uzabase/play-cdc/releases)から自分の環境に合うzipファイルをダウンロードして、
```
$ gauge install play-cdc -f [zipファイル名]
```

## Usage

### 準備

1. コンシューマ側のテストで、`manifest.json`の`Plugins`に`play-cdc`を追加する。
2. `env/default/playcdc.properties`ファイルを作り、以下のプロパティを設定する。
  - `cdc_output_base_path`: 契約の出力先パス（ディレクトリの**絶対パス**）
  - `cdc_consumer_name`: コンシューマ（=E2E自体のテスト対象）の名前
3. WireMockでスタブしているプロバイダーそれぞれについて、`env/default/playcdc.properties`に以下2つのプロパティを設定する。（`[n]`は1以上の整数）
  - `cdc_provider_endpoint_[n]`: WireMockのエンドポイント
  - `cdc_provider_name_[n]`: スタブしているプロバイダーの名前

### 実行

1. スペック（シナリオ）を実行する。
  - なお、WireMockのリクエスト履歴はExecution hooksなどで予めクリア（削除）しておくこと。
2. すべてのシナリオが成功した場合、設定しておいた出力先パスに、プロバイダごとに成立した契約を記述したspecファイルが出力される。
  - リクエストボディがある場合は、fixtures以下にリクエストボディのファイルも出力される。

### クエリ条件からのURL生成

- `url` が指定されている場合は、そのURLを使います。
- `urlPath` と `queryParameters` が指定されている場合は、`equalTo` の文字列をクエリ値として使います。`caseInsensitive` を含む条件も読み込めます。
- `absent: true` のパラメータはURLに含めません。
- 正規表現や複合条件など、値を決められない条件がある場合は、WireMockに記録された実リクエストのURLをそのまま使います。この場合は、URLが異なるリクエストごとに契約が生成されます。

### デバッグモード

- `env/default/playcdc.properties`に`cdc_debug=true`を設定すると、デバッグ用のログを出力できます。
