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

	fmt.Println("--------------------------------")
	fmt.Println("START: PDF結合処理を開始します")

	// 1. pdfフォルダのファイルを取得
	dir := "./pdf"
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// グループ分け用のマップ (キー: "pdf-file-a", 値: [ファイル名リスト])
	groups := make(map[string][]string)

	for _, entry := range entries {

		name := filepath.Join(dir, entry.Name())

		// ディレクトリは除外
		if entry.IsDir() {
			continue
		}

		// 出力ファイル自身を除外
		if strings.HasSuffix(name, "-mix.pdf") {
			continue
		}

		if strings.ToLower(filepath.Ext(name)) == ".pdf" {
			// 最後の-のインデックスを取得。それをキーとしてグループ分け
			lastIndex := strings.LastIndex(name, "-")
			if lastIndex != -1 {
				groupKey := name[:lastIndex]
				groups[groupKey] = append(groups[groupKey], name)
			}
		}

	}

	// 2. グループごとに結合処理を実行
	for groupKey, files := range groups {
		// ファイル名でソート（1, 2...の順にするため）
		sort.Strings(files)

		outputFile := fmt.Sprintf("%s-mix.pdf", groupKey)

		fmt.Printf("グループ [%s] を結合中... -> %s\n", groupKey, files)

		// 結合実行
		err := api.MergeCreateFile(files, outputFile, false, nil)
		if err != nil {
			log.Printf("❌ エラーが発生しました: %v", err)
			continue
		}
		fmt.Printf("  ✅ 完了: %s をまとめました\n", groupKey)
	}

	fmt.Println("END: PDF結合処理を完了しました")
	fmt.Println("--------------------------------")
}
