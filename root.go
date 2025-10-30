// github.com/shouni/go-cli-base/root.go

package clibase

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// GlobalFlags は、すべてのコマンドで利用できる共通フラグを保持する構造体です。
// アプリケーション側から clibase.Flags.Verbose のようにアクセスできます。
type GlobalFlags struct {
	Verbose    bool
	ConfigFile string
}

var Flags GlobalFlags

// NewRootCmd は、指定されたアプリケーション名に基づいてルートコマンドの基盤を生成します。
func NewRootCmd(appName string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   appName,
		Short: fmt.Sprintf("A CLI tool for %s.", appName),
		// PersistentPreRun: 全てのコマンド実行前に共通して行いたい処理を定義
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if Flags.Verbose {
				// ロギングライブラリの初期化などをここで行う
				fmt.Println("Verbose mode enabled by clibase.")
			}
			// 設定ファイル読み込みロジックなどをここに記述
		},
		// Run は通常、Help表示などに利用されます
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// 共通フラグの定義 (永続フラグとして定義することで、全てのサブコマンドで利用可能)
	rootCmd.PersistentFlags().BoolVarP(&Flags.Verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringVarP(&Flags.ConfigFile, "config", "c", "", "Config file path")

	return rootCmd
}

// Execute は、CLIアプリケーションのエントリポイントです。
// アプリケーション固有のサブコマンドをルートコマンドに追加し、実行します。
func Execute(appName string, cmds ...*cobra.Command) {
	rootCmd := NewRootCmd(appName)
	rootCmd.AddCommand(cmds...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
