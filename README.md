# 📚 Go Cli Base

**`go-cli-base`** は、Go言語でコマンドラインインターフェース (CLI) アプリケーションを迅速に構築するための、**`spf13/cobra`** ベースの共通基盤を提供するパッケージです。

このパッケージを利用することで、全てのサブコマンドで利用可能な共通フラグの定義や、アプリケーション全体での初期化処理（例：設定ファイルの読み込み、詳細出力の有効化）を標準化し、boilerplateコードを削減できます。

## ✨ 特徴

  * **`cobra`** ベース: 強力なCLI構築ライブラリ **`spf13/cobra`** を基盤としています。
  * **共通フラグの提供**: `verbose` (`-v`) と `config` (`-c`) の2つの永続フラグを標準で提供します。
  * **初期化処理の抽象化**: 全てのコマンド実行前に共通処理（例：詳細モードの有効化、設定ファイルの読み込み）を挿入するための構造を提供します。
  * **シンプルなエントリポイント**: アプリケーションの実行を簡潔な `Execute` 関数にカプセル化します。

-----

## 🛠️ インストール

標準のGoコマンドでインストールできます。

```bash
go get github.com/shouni/go-cli-base
```

-----

## 🚀 使用方法

### 1\. ルートコマンドの初期化と実行

アプリケーションのエントリポイントとなる **`main.go`** で、`clibase.Execute` 関数を使用してCLIを起動します。

```go
// main.go
package main

import (
	"fmt"
	"os"

	"github.com/shouni/go-cli-base"
	"github.com/spf13/cobra"
)

// アプリケーション固有のサブコマンドを定義
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a greeting message.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from the CLI app!")
		// 共通フラグの値にアクセス
		if clibase.Flags.Verbose {
			fmt.Printf("Config file path: %s\n", clibase.Flags.ConfigFile)
		}
	},
}

func main() {
	// アプリケーション名とサブコマンドを渡して実行
	clibase.Execute("my-awesome-cli", helloCmd)
}
```

### 2\. 共通フラグへのアクセス

定義された共通フラグの値は、**`clibase.Flags`** グローバル変数を通じて、アプリケーションのどこからでもアクセス可能です。

```go
// アプリケーションの任意の場所で
if clibase.Flags.Verbose {
    // ... 詳細モードのロギング ...
}
if clibase.Flags.ConfigFile != "" {
    // ... 設定ファイルの読み込み ...
}
```

### 3\. 提供される共通機能

`clibase.NewRootCmd` 関数が返すルートコマンドには、以下の永続フラグが自動的に設定されます。

| フラグ | ショートカット | 型 | 説明 |
| :--- | :--- | :--- | :--- |
| `--verbose` | `-v` | `bool` | 詳細な出力 (デバッグ情報など) を有効にします。 |
| `--config` | `-c` | `string` | 設定ファイルのパスを指定します。 |

また、`clibase.NewRootCmd` 内の **`PersistentPreRun`** では、`--verbose` フラグが有効な場合にシンプルなメッセージを出力する処理が含まれており、ここにロギングライブラリの初期化や設定ファイルの読み込みロジックを追加・カスタマイズすることができます。

-----

## ⚙️ 内部構造

`go-cli-base` パッケージの主要なコンポーネントは以下の通りです。

### `GlobalFlags` 構造体

全てのコマンドで利用できる共通フラグの値を保持します。

```go
type GlobalFlags struct {
	Verbose    bool
	ConfigFile string
}
// clibase.Flags として公開されており、アプリケーションからアクセス可能
var Flags GlobalFlags
```

### `NewRootCmd` 関数

アプリケーション名に基づき、共通フラグを設定した **`*cobra.Command`** を生成します。

### `Execute` 関数

CLIアプリケーションのエントリポイントとして機能し、アプリケーション固有のサブコマンドをルートコマンドに追加した後、**`rootCmd.Execute()`** を呼び出します。

-----

## 📜 ライセンス (License)

このプロジェクトは [MIT License](https://opensource.org/licenses/MIT) の下で公開されています。

