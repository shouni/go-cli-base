package clibase

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// GlobalFlags は、すべてのコマンドで利用できる共通フラグを保持する構造体です。
// アプリケーション側から clibase.Flags.Verbose のようにアクセスできます。
type GlobalFlags struct {
	Verbose    bool
	ConfigFile string
}

var Flags GlobalFlags

// CustomFlagFunc は、アプリケーション固有の永続フラグを追加するためのコールバック関数の型です。
type CustomFlagFunc func(rootCmd *cobra.Command)

// CustomPreRunEFunc は、アプリケーション固有の実行前チェック（エラーを返すことが可能）のためのコールバック関数の型です。
type CustomPreRunEFunc func(cmd *cobra.Command, args []string) error

// NewRootCmd は、指定されたアプリケーション名に基づいてルートコマンドの基盤を生成します。
// アプリケーション固有のフラグ追加や、PersistentPreRunE のロジックを注入できます。
func NewRootCmd(appName string, addFlags CustomFlagFunc, preRunE CustomPreRunEFunc) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   appName,
		Short: fmt.Sprintf("A CLI tool for %s.", appName),
		Long:  fmt.Sprintf("The CLI tool for %s. Use a subcommand to perform a task.", appName),

		// PersistentPreRunE: 全てのコマンド実行前に共通して行いたい処理を定義
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 1. clibase 共通の PersistentPreRun 処理
		if Flags.Verbose {
		// ロギングライブラリの初期化などをここで行う
		fmt.Println("Verbose mode enabled by clibase.")
	}
		// 設定ファイル読み込みロジックなどをここに記述

		// 2. アプリケーション固有の PersistentPreRunE 処理を実行
		if preRunE != nil {
		return preRunE(cmd, args)
	}
		return nil
	},
		// Run は通常、Help表示などに利用されます
		Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	}

	// 共通フラグの定義 (永続フラグとして定義することで、全てのサブコマンドで利用可能)
	rootCmd.PersistentFlags().BoolVarP(&Flags.Verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringVarP(&Flags.ConfigFile, "config", "c", "", "Config file path")

	// アプリケーション固有のフラグを追加
	if addFlags != nil {
		addFlags(rootCmd)
	}

	return rootCmd
}

// Execute は、CLIアプリケーションのエントリポイントです。
// アプリケーション固有のサブコマンドとカスタマイズ関数をルートコマンドに追加し、実行します。
func Execute(appName string, addFlags CustomFlagFunc, preRunE CustomPreRunEFunc, cmds ...*cobra.Command) {
	rootCmd := NewRootCmd(appName, addFlags, preRunE)
	rootCmd.AddCommand(cmds...)

	if err := rootCmd.Execute(); err != nil {
		// Execute()でエラーが発生した場合、cobraは既にエラーを出力しているため
		// os.Exit(1) のみで良いが、念のため fmt.Println(err) を残す
		fmt.Println(err)
		os.Exit(1)
	}
}
