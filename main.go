package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitly/go-simplejson"
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

func hasKey(j, k int) bool {
	if j < 3 && (k == 6 || k == 7) {
		// キーがない場所
		return false
	}
	return true
}

var LayerNum int

var keymap [][][]string

func main() {
	log.SetFlags(log.Llongfile)
	flag.Parse()
	arg := flag.Arg(0)
	ext := filepath.Ext(arg)
	switch ext {
	case ".json":
		readJson(arg)
	case ".h":

		input := readText(arg)
		c := cutHeader(input)   // 不要部分削除 前後と半角空白や括弧閉じなど
		layers := divNewLine(c) // レイヤーごとに改行で分割
		readHeader(&layers)
	default:
		fmt.Println(ext)
	}
	printKeymap()
}

func printKeymap() { // {{{
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
} // }}}

func readHeader(layers *[]string) { // {{{
	// keymap保存場所を作る
	LayerNum = len(*layers)
	keymap = make([][][]string, LayerNum)
	for i := 0; i < LayerNum; i++ {
		keymap[i] = newLayer(RowNum, ColumnNum)
	}

	// つめる
	for i := 0; i < LayerNum; i++ {
		// keysをいい感じに参照するためのカウンタ
		count := 0
		for j := 0; j < RowNum; j++ {
			keys := strings.Split((*layers)[i], ",")
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

/*
type keymapJson struct {
	//keyboard string     `json:"keyboard"`
	//keymap   string     `json:"keymap"`
	//layout   string     `json:"layout"`
	layers interface{} `json:"layers"`
	//author   string     `json:"author"`
	//notes    string     `json:"notes"`
}
*/

func StringDArray(j *Json) ([]string, error) {

	if a, ok := (j.data).([]interface{}); !ok {
	}
	// a [][]interface{} <- [][]string{}

	arr, err := j.Array()
	if err != nil {
		return nil, err
	}

	retArrArr := make([][]string, 0, len(arr))
	for _, a := range arr {
		if a == nil {
			retArr := make([]string, 0, 100)
			retArr = append(retArr, retArr)
			continue
		}
		s, ok := a.(string)
		if !ok {
			return nil, errors.New("type assertion to []string failed")
		}
		retArr = append(retArr, s)
	}
	return retArr, nil
}

func readJson(path string) { // {{{

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイル %v が読み込めません\n", path)
		log.Println(err)
		panic(err)
	}

	json, err := simplejson.NewJson(b)
	keymaps, _ := json.Get("layers").StringArray()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	fmt.Printf("%v\n", keymaps)

	for i, km := range keymaps {
		for j := 0; j < len(km); j++ {
			fmt.Printf("%d\n", i)
			fmt.Printf("%v,\n", keymaps[i][j])
		}
	}

	return

	layers := []string{"", ""}

	// 前処理

	// keymap保存場所を作る
	LayerNum := len(layers)
	keymap = make([][][]string, LayerNum)
	for i := 0; i < LayerNum; i++ {
		keymap[i] = newLayer(RowNum, ColumnNum)
	}

	// つめる
	//for i, l := range p.layers {
	for i := 0; i < LayerNum; i++ {
		//fmt.Printf("%d : %s\n", i, l)
		// keysをいい感じに参照するためのカウンタ
		// keyがない場所では進めない
		count := 0
		for j := 0; j < RowNum; j++ {
			keys := strings.Split(layers[i], ",")
			//keys := l
			for k := 0; k < ColumnNum; k++ {
				t := &keymap[i][j][k] // 代入先
				if hasKey(j, k) {
					*t = keys[count]
					count++
				} else {
					*t = "xx" // helix にキーがない場所j=3行目以内でk=6,7列目
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

func cutJson(in string) string { // {{{2
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
