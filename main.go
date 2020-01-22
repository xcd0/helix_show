package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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
	RowNum    = 5
	ColumnNum = 14
)

var LayerNum int

var keymap [][][]string

func main() {
	flag.Parse()
	input := ReadText(flag.Arg(0))
	read(input)

	fmt.Println("/*")
	for i := 0; i < LayerNum; i++ {
		fmt.Printf("layer %v\n", i)
		fmt.Println("|--------+--------+--------+--------+--------+--------+        +        +--------+--------+--------+--------+--------+--------|")
		for j := 0; j < RowNum; j++ {
			for k := 0; k < ColumnNum; k++ {
				if j < 3 && k == 7 {
					fmt.Printf(" ")
				} else {
					fmt.Printf("|")
				}
				if j < 3 && (k == 6 || k == 7) {
					// helixにない場所
					fmt.Printf("        ")
				} else {
					//fmt.Printf(" %8v ", keymap[i][j][k])
					num := len(keymap[i][j][k])
					// m + num + n = 8にする
					// numが6文字ならm=1,n=1
					// numが5文字ならm=1,n=2
					m := (8 - num) / 2
					n := 8 - m - num
					for l := 0; l < m; l++ {
						fmt.Printf(" ")
					}
					fmt.Printf("%v", keymap[i][j][k])
					for l := 0; l < n; l++ {
						fmt.Printf(" ")
					}
				}
				if k == ColumnNum-1 {
					fmt.Printf("|\n")
					if j < 2 {
						fmt.Println("|--------+--------+--------+--------+--------+--------+        +        +--------+--------+--------+--------+--------+--------|")
					} else {
						fmt.Println("|--------+--------+--------+--------+--------+--------+--------+--------+--------+--------+--------+--------+--------+--------|")
					}
				}
			}
		}
		//fmt.Println("+--------+--------+--------+--------+--------+--------+--------+--------+--------+--------+--------+--------+--------+--------+")
		fmt.Printf("\n")
	}
	fmt.Println("*/")
}

func read(in string) { // {{{
	// 前処理

	// 不要部分削除 前後と半角空白や括弧閉じなど
	c := cut(in)
	layers := divNewLine(c) // レイヤーごとに改行で分割

	// keymap保存場所を作る
	LayerNum = len(layers)
	keymap = make([][][]string, LayerNum)
	for i := 0; i < LayerNum; i++ {
		keymap[i] = newLayer(RowNum, ColumnNum)
	}

	// つめる
	for i := 0; i < LayerNum; i++ {
		// keysをいい感じに参照するためのカウンタ
		count := 0
		for j := 0; j < RowNum; j++ {
			keys := strings.Split(layers[i], ",")
			for k := 0; k < ColumnNum; k++ {
				t := &keymap[i][j][k]
				if j < 3 && (k == 6 || k == 7) {
					// helix にキーがない場所
					*t = "xx"
				} else {
					*t = keys[count]
					count++
				}
				// main processing
				// LSFT() とかの処理
				for k, v := range KEYMAP_FUNC {
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
					// LT( 5, KC_NO ) とかは LT(5と KC_NOに別れる
					// 半角空白やコンマは入らない
					layerNum := (*t)[3:len(*t)]
					tmp := keys[count]
					// これでtmpに数値だけ入る
					// keys[count]が次のキーコードになる
					if _, ok := KEYMAP[keys[count]]; ok {
						tmp = KEYMAP[keys[count]]
					}
					*t = "LT" + layerNum + "," + tmp
					count++
				}
				if _, ok := KEYMAP[*t]; ok {
					*t = KEYMAP[*t]
				}
				//fmt.Printf("%3v,%3v,%3v : %10v\n", i, j, k, t)
			}
		}
	}
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

func cut(in string) string { // {{{2
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
			case ' ':
				continue
			case ')':
				continue
			case '}':
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

func ReadText(path string) string { // {{{2

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
		return ConvertNewline(string(b), "\n")
	}

	str, err := nkf.ToUtf8(string(b), charset)

	str = ConvertNewline(str, "\n")

	return str
} // }}}2

func ConvertNewline(str, nlcode string) string { // {{{2
	// 改行コードを統一する
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
} // }}}2
// }}}1
