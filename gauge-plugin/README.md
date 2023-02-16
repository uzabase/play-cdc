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
  - `cdc_output_base_path`: 契約の出力先パス（ディレクトリ）
  - `cdc_consumer_name`: コンシューマ（=E2E自体のテスト対象）の名前
3. WireMockでスタブしているプロバイダーそれぞれについて、`env/default/playcdc.properties`に以下2つのプロパティを設定する。（`[n]`は1以上の整数）
  - `cdc_provider_endpoint_[n]`: WireMockのエンドポイント
  - `cdc_provider_name_[n]`: スタブしているプロバイダーの名前

### 実行

1. スペック（シナリオ）を実行する。
  - なお、WireMockのリクエスト履歴はExecution hooksなどで予めクリア（削除）しておくこと。
2. すべてのシナリオが成功した場合、設定しておいた出力先パスに、プロバイダごとに成立した契約を記述したspecファイルが出力される。
  - リクエストボディがある場合は、fixtures以下にリクエストボディのファイルも出力される。

### デバッグモード

- `env/default/playcdc.properties`に`cdc_debug=true`を設定すると、デバッグ用のログを出力できます。
