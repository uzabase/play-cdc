# play-cdc

Gaugeを使用したCDCをシンプルに、かつ低い労力で作成できるようにすることで、CDCを楽しめるようにするためのプロジェクトです。

## java-library

WireMockのスタブ定義用JSONファイルを直接契約書として扱うことで、プロバイダー側のCDCテストを簡潔に書けるようにするJavaライブラリ。  
このアプローチは一通り探索できたため、現在は次のGaugeプラグインを中心に探索（開発）を進めています。

## gauge-plugin

WireMockのスタブ定義を契約書の原型として扱い、プロバイダー側用の契約書（=specファイル）を自動生成するGaugeプラグイン。
現在探索中。

## tools

### ccc.mjs

carbon-copy contract。
Javaライブラリを使用する際に、コマンドラインの対話形式で、契約を表すJSONファイルのシンボリックリンクを作ることができるツール。
