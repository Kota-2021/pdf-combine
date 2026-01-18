package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	// 1. pdfフォルダのファイルを取得
	entries, err := os.ReadDir("./pdf")
	if err != nil {
		log.Fatal(err)
	}

	var files []string
	for _, entry := range entries {
		// 2. 拡張子が .pdf のものだけを抽出（大文字小文字を区別しない）
		if !entry.IsDir() && strings.ToLower(filepath.Ext(entry.Name())) == ".pdf" {
			// 出力用ファイル自身を除外するためのチェック
			if entry.Name() != "merged_output.pdf" {
				files = append(files, "./pdf/"+entry.Name())
			}
		}
	}

	// 3. ファイル名でソート（pdf-file-1, 2, 13...の順になるように）
	sort.Strings(files)

	if len(files) < 2 {
		fmt.Println("結合には2枚以上のPDFファイルが必要です。")
		return
	}

	fmt.Printf("以下のファイルを結合します:\n %v\n", files)

	// 4. 結合実行
	err = api.MergeCreateFile(files, "merged_output.pdf", false, nil)
	if err != nil {
		log.Fatalf("❌ エラーが発生しました: %v", err)
	}

	fmt.Println("✅ 作成完了: merged_output.pdf")
}
