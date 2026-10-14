# Asista

Go で作る自作プログラミング言語 Asista のインタプリタと、それを Claude から使うための MCP サーバです。

## 構成

| パス | 役割 | 章 |
| --- | --- | --- |
| `token/` | トークン定義 | 第1章 |
| `lexer/` | 字句解析 | 第1章 |
| `ast/` | 抽象構文木 | 第2章 |
| `parser/` | 構文解析（Pratt パーサ） | 第2章 |
| `object/` | オブジェクトと環境 | 第3章 |
| `evaluator/` | 評価器 | 第3章 |
| `cmd/repl/` | 対話実行環境 | 第1章〜 |
| `cmd/mcp-server/` | MCP サーバ | 第7章〜 |
| `docs/spec.md` | 言語仕様書 | 第5章 |
| `docs/notes/` | 章ごとの学習メモ | 随時 |

## 使い方

```sh
go run ./cmd/repl      # REPL を起動
go test ./...          # すべてのテストを実行
```

## 進め方

- 各章はテストを先に書き、通ったらコミットする
- 章の終わりにタグを打つ（例: `git tag ch1-lexer`）
