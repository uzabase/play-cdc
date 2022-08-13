# Play CDC (beta)

## Install

```
$ make install
```

## Usage

### 準備

1. `manifest.json`の`Plugins`に`play-cdc`を追加する。
2. `env/default/playcdc.properties`ファイルを作り、WireMockでスタブしているAPIそれぞれについて、以下3つのプロパティを設定する。（`[n]`は1以上の整数）
  - `cdc_api_endpoint_[n]`: WireMockのエンドポイント
  - `cdc_api_name_[n]`: スタブしているAPIの名前
  - `cdc_output_path_[n]`: 契約の出力先パス（ディレクトリ）

### 実行

1. スペック（シナリオ）を実行する。
  - なお、WireMockのリクエスト履歴はExecution hooksなどで予めクリア（削除）しておくこと。
2. すべてのシナリオが成功した場合、設定しておいた出力先パスに成立した契約を記述したspecファイルが出力される。
