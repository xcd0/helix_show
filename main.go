package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xcd0/go-nkf"
)

/* helixキーボードでは // {{{
keymapの
(*,0,6), (*,0,7)
(*,1,6), (*,1,7)
(*,2,6), (*,2,7)
の6個が使えない
5 * 14 - 6 = 64
*/
// }}}

const (
	rowNum      = 5  // 行数
	columnNum   = 14 // 列数
	maxLayerNum = 16 // レイヤーの最大数
	keyNum      = 64 // キーの数 rowNum * columnNum - 無効なキーの数
)

func hasKey(j, k int) bool {
	if j < 3 && (k == 6 || k == 7) {
		// キーがない場所
		return false
	}
	return true
}

var keymap [][][]string

func main() {
	log.SetFlags(log.Llongfile)
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("エラー : 引数が与えられていません。")
	}
	arg := flag.Arg(0)
	apath, _ := filepath.Abs(flag.Arg(0))
	outputDir := filepath.Dir(apath)
	ext := filepath.Ext(arg)
	switch ext {
	case ".json":
		input := readText(arg)
		createdHeader := readJson(input)

		c := cutHeader(*createdHeader)
		layers := divNewLine(c) // レイヤーごとに改行で分割
		output := readHeader(&layers)

		outputName := filepath.Base(apath[:len(apath)-5] + ".h")
		outputFilePath := filepath.Join(outputDir, outputName)
		var outputFile *os.File
		var err error
		outputFile, err = os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			// Openエラー処理
			log.Println("エラー : 出力先ファイルが開けません。")
			log.Println("       : 他のプログラムでファイルを開いていませんか？")

			t := time.Now().Local()
			outputName = filepath.Base(apath) + "_" + fmt.Sprintf(t.Format("2006-01-02-15-04-05")) + ".h"
			outputFilePath = filepath.Join(outputDir, outputName)
			outputFile, err = os.OpenFile(outputFilePath, os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				log.Fatal(fmt.Sprintf("エラーメッセージ : %v\nエラー終了します。\n", err))
			}
			log.Println("       : 別ファイル名で保存します。")
		}
		defer outputFile.Close()
		outputFile.Write(([]byte)(*output))

		//fmt.Println(*output)
	case ".h":
		input := readText(arg)
		c := cutHeader(input)   // 不要部分削除 前後と半角空白や括弧閉じなど
		layers := divNewLine(c) // レイヤーごとに改行で分割
		output := readHeader(&layers)

		// もともとheaderファイルが入力された場合標準出力に出力するだけにする。
		// jsonの時はheaderファイルを生成するのでいいかんじに埋め込む
		fmt.Println(*output)
	default:
		log.Fatalf("Error: 拡張子 %s が不正です。", ext)
	}
}

func readHeader(layers *[]string) *string { // {{{
	// keymap保存場所を作る
	layerNum := len(*layers)
	keymap = make([][][]string, layerNum)
	for i := 0; i < layerNum; i++ {
		keymap[i] = newLayer(rowNum, columnNum)
	}
	// keymapにつめる
	for i := 0; i < layerNum; i++ {
		count := 0 // keysをいい感じに参照するためのカウンタ helixキーボードにないところを飛ばす
		for j := 0; j < rowNum; j++ {
			keys := strings.Split((*layers)[i], ",") // コンマで分割する これでレイヤー1層分のキーコードが入る
			for k := 0; k < columnNum; k++ {
				// 一文字づつ
				t := &keymap[i][j][k] // ここに保存する
				if j < 3 && (k == 6 || k == 7) {
					// helix にキーがない場所
					*t = "xx"
				} else {
					*t = keys[count]
					count++
				}
				// main processing
				// LSFT() とか MO() とかの処理
				for k, v := range KEYMAP_FUNC { // LT()以外はここで処理する
					if strings.Contains(*t, k) {
						tmp := (*t)[len(k):len(*t)]
						if _, ok := KEYMAP[tmp]; ok {
							tmp = KEYMAP[tmp]
						}
						// LSFT(KC_1)とかを(S)1に変える
						*t = v + tmp
					}
				}
				if strings.Contains(*t, "LT(") {
					// LT( 5, KC_NO ) とかは LT(5と KC_NOに別れる 半角空白やコンマは入らない
					l := (*t)[3:len(*t)]
					tmp := keys[count]
					// これでtmpに数値だけ入る keys[count]が次のキーコードになる
					if _, ok := KEYMAP[keys[count]]; ok {
						tmp = KEYMAP[keys[count]]
					}
					*t = "LT" + l + "," + tmp
					count++
				}
				if _, ok := KEYMAP[*t]; ok {
					*t = KEYMAP[*t]
				}
				//fmt.Printf("%3v,%3v,%3v : %10v\n", i, j, k, t)
			}
		}
	}

	output := ""
	output += "const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {\n"
	for i := 0; i < layerNum; i++ {
		output += "\t/*\n"
		output += fmt.Sprintf("\t\tlayer %v\n", i)
		output += "\t\t+---------+---------+---------+---------+---------+---------+         +         +---------+---------+---------+---------+---------+---------+\n\t\t"
		for j := 0; j < rowNum; j++ {
			for k := 0; k < columnNum; k++ {
				if j < 3 && k == 7 {
					output += " "
				} else {
					output += "|"
				}
				if j < 3 && (k == 6 || k == 7) {
					// helixにない場所
					output += "         "
				} else {
					//fmt.Printf(" %8v ", keymap[i][j][k])
					num := len(keymap[i][j][k])
					// m + num + n = 9にする
					// numが6文字ならm=1,n=1
					// numが5文字ならm=1,n=2
					m := (9 - num) / 2
					n := 9 - m - num
					for l := 0; l < m; l++ {
						output += " "
					}
					output += fmt.Sprintf("%v", keymap[i][j][k])
					for l := 0; l < n; l++ {
						output += " "
					}
				}
				if k == columnNum-1 {
					output += "|\n"
					if j < 2 {
						// これは上のほう
						output += "\t\t|---------+---------+---------+---------+---------+---------+         +         +---------+---------+---------+---------+---------+---------|\n\t\t"
					} else if j == rowNum-1 {
						// これが最後
						output += "\t\t+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+\n"
					} else {
						// これは下のほう
						output += "\t\t|---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------|\n\t\t"
					}
				}
			}
		}

		output += "\t*/\n"
		// ここに元のレイヤーの記述を書く
		if i != rowNum-1 {
			output += fmt.Sprintf("\t[%d] = LAYOUT(%s),\n", i, (*layers)[i])
		} else {
			output += fmt.Sprintf("\t[%d] = LAYOUT(%s)\n", i, (*layers)[i])
		}
		output += "\n"
	}
	output += "};\n\n"
	return &output
}

// }}}

func readJson(jsonText string) *string { // {{{

	countC := 0 // ,の数カウント
	countP := 0 // []の数カウント
	countL := 0 // layer番号 countC := 0 // ,の数カウント
	keymap := make([]string, maxLayerNum)

	rs := []rune(jsonText)
	//fmt.Printf("%s", string(rs))
	i := strings.Index(jsonText, "\"layers\":")
	if i == -1 {
		log.Fatal("Error: jsonの形式が正しくありません。異常終了します。")
	}
	for i += len("\"layers\":"); i < len(rs); i++ { // "layers":まで進んでいる

		// ここはlayersのなか 1文字づつ処理する
		switch rs[i] {
		case ':': // あったらおかしい
			log.Fatal("Error: jsonの形式が正しくありません。異常終了します。")
		case ',':
			countC += 1
			//fmt.Printf("%d ", countC)
		case '[':
			countP += 1
			//fmt.Printf("countP : %d\n", countP)
		case ']':
			countP -= 1 // []の層の深さをひとつ戻す
			if countP == 0 {
				// 終わり
				i = len(rs)
				break
			}
			countL += 1 // layerを次のレイヤーに
			countC = -1 // ,のカウンタをリセット // この後にすぐ,があるので
		case '"':
			// ここはキーコードの手前のダブルクオート
			// 次の"まで読み込む
			i++      // ダブルクオートの次の文字を指すようにする
			pre := i // この位置を保存
			for rs[i] != '"' {
				i++ // ダブルクオートがない間進める
			}
			literal := string(rs[pre:i]) // ダブルクオートの次の文字からダブルクオートの手前まで切り取る つまりキーコードが入る
			//fmt.Printf("|%v,%d ", literal, countL)
			if strings.Contains(literal, "ANY(") { // QMK configuratorでは許容されているが標準ではないマクロ #define ANY(X) X とかでいい
				fmt.Println("")
				tmp := literal[len("ANY(") : len(literal)-1]
				// ANY(KC_1)とかをKC_1に変える つまり中身に置き換える
				literal = tmp
			}
			if countC+1 == keyNum {
				keymap[countL] += literal // これでこのレイヤーは終わり
			} else {
				keymap[countL] += literal + ", " // これでliteralにキーコードが入った
			}
		default: // 無視
		}
	}

	output := ""
	output += "#define ANY(X) (X)\n"
	output += "const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {\n"
	for i = 0; i < countL; i++ {
		if i != countL-1 {
			output += fmt.Sprintf("\t[%d] = LAYOUT(%s),\n", i, keymap[i])
		} else {
			output += fmt.Sprintf("\t[%d] = LAYOUT(%s)\n", i, keymap[i])
		}
	}
	output += "};"
	return &output
}

// }}}

// other {{{1
func divNewLine(in string) []string { // {{{2
	lines := strings.Split(in, "\n")
	length := len(lines)
	cc := make([]string, length)

	for i, line := range lines {
		tmp := line[strings.LastIndex(line, "LAYOUT(")+len("LAYOUT("):]
		cc[i] = tmp
	}
	return cc
} // }}}2

func cutHeader(in string) string { // {{{2
	rs := []rune(in)
	pre := '\n'
	out := ""
	flag := false
	for _, r := range rs {
		if pre == '{' && r == '\n' {
			flag = true
			pre = r
			continue
		}
		if flag {
			switch r {
			case ' ': // 捨てる
				continue
			case ')': // 捨てる
				continue
			case '}': // 捨てる
				return out[:len(out)-1]
			default:
				out += string(r)
			}
		}
		pre = r
	}
	return out[:len(out)-1]
} // }}}2

func newLayer(numRow, numColumn int) [][]string { // {{{2

	rs := make([][]string, numRow)
	for i := 0; i < len(rs); i++ {
		rs[i] = make([]string, numColumn)
	}
	return rs
} // }}}2

func readText(path string) string { // {{{2

	// 与えられたパスの文字列について
	// そのパスにあるファイルをテキストファイルとして読み込む

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイル%vが読み込めません\n", path)
		log.Println(err)
		panic(err)
		return ""
	}
	// ファイルの文字コード変換
	charset, err := nkf.CharDet(b)
	if err != nil {
		/*
			fmt.Fprintf(os.Stderr, "文字コード変換に失敗しました\nutf8を使用してください\n")
			log.Println(err)
			panic(err)
			return ""
		*/
		return convertNewline(string(b), "\n")
	}

	str, err := nkf.ToUtf8(string(b), charset)

	str = convertNewline(str, "\n")

	return str
} // }}}2

func convertNewline(str, nlcode string) string { // {{{2
	// 改行コードを統一する
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
} // }}}2
// }}}1
