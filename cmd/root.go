package cmd

import (
	"fmt"
	"github.com/smark-d/epub-translator/parser"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	var filePath, translationEngine, sourceLanguage, targetLanguage string
	var rootCmd = &cobra.Command{
		Use:   "translator",
		Short: "A command-line epub translation tool",
		Run: func(cmd *cobra.Command, args []string) {
			parser.GetParser("epub", filePath, sourceLanguage, targetLanguage, translationEngine).Parse()
		},
	}

	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")
	rootCmd.Flags().StringVarP(&translationEngine, "engine", "e", "google", "Translation engine (google, openai)")
	rootCmd.Flags().StringVarP(&sourceLanguage, "source", "s", "en", "Source language")
	rootCmd.Flags().StringVarP(&targetLanguage, "target", "t", "zh-CN", "Target language")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
