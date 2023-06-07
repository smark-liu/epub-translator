package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/smark-d/epub-translator/common"
	"github.com/smark-d/epub-translator/parser"
	"github.com/spf13/cobra"
)

func Execute() {
	var filePath, translationEngine, sourceLanguage, targetLanguage string
	var keepOrigin string

	var rootCmd = &cobra.Command{
		Use:   "translator",
		Short: "A command-line epub translation tool",
		Run: func(cmd *cobra.Command, args []string) {
			keepOriginBool, _ := strconv.ParseBool(keepOrigin)
			parser.GetParser(common.GetParserByName(filePath), filePath, sourceLanguage, targetLanguage, translationEngine, keepOriginBool).Parse()
		},
	}

	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "File path")
	rootCmd.Flags().StringVarP(&translationEngine, "engine", "e", "google", "Translation engine (google, openai)")
	rootCmd.Flags().StringVarP(&sourceLanguage, "source", "s", "en", "Source language")
	rootCmd.Flags().StringVarP(&targetLanguage, "target", "t", "zh-CN", "Target language")
	rootCmd.Flags().StringVarP(&keepOrigin, "keep", "k", "true", "Keep the original text")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
