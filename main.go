package main

import (
	"fmt"
	"log"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	// 結合したいファイルリスト
	files := []string{
		"pdf-file-1.pdf",
		"pdf-file-2.pdf",
		"pdf-file-3.pdf",
	}

	// 出力ファイル名
	outputFile := "merged_result.pdf"

	// PDFを結合（第2引数は設定オブジェクト、nilでデフォルト）
	err := api.MergeCreateFile(files, outputFile, false, nil)
	if err != nil {
		log.Fatal("❌ 結合に失敗しました:", err)
	}

	fmt.Printf("✅ 正常に結合されました: %s\n", outputFile)
}
