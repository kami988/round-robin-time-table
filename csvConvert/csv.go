package csvconvert

import (
	"encoding/csv"
	"os"
)

func OutputCSV(filename string, combiArray [][][]int, persons []string) error {
	records := cnvRecords(combiArray, persons)
	f, err := os.Create(filename) // 書き込む先のファイル
	if err != nil {
		return err
	}

	if err = csv.NewWriter(f).WriteAll(records); err != nil {
		return err
	}
	return nil
}

func cnvRecords(combiArray [][][]int, persons []string) [][]string {
	// あらかじめ配列長確保（ヘッダー分１行追加）
	records := make([][]string, len(combiArray)+1)
	for i := range records {
		records[i] = make([]string, len(persons))
	}

	// ヘッダー追加
	copy(records[0], persons)

	for ci, combiLine := range combiArray {
		for pi := range persons {
			for _, combi := range combiLine {
				if sortedCombi := includePerson(combi, pi); sortedCombi != nil {
					records[ci+1][pi] = persons[sortedCombi[0]] + " - " + persons[sortedCombi[1]]
				}
			}
		}
	}
	return records
}

func includePerson(combi []int, personIndex int) []int {
	for i, combiElement := range combi {
		// 組み合わせに対象者が存在したとき
		if combiElement == personIndex {
			// 対象者が先になるよう返却
			if i == 0 {
				return []int{combi[0], combi[1]}
			}
			return []int{combi[1], combi[0]}
		}
	}
	return nil
}
