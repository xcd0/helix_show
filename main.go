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

	for i := 0; i < LayerNum; i++ {
		fmt.Printf("layer %v\n", i)
		for j := 0; j < RowNum; j++ {
			for k := 0; k < ColumnNum; k++ {
				fmt.Printf("|")
				fmt.Printf(" %8v ", keymap[i][j][k])
				if k == ColumnNum-1 {
					fmt.Printf("|\n")
				}
			}
		}
		fmt.Printf("\n")
	}
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
					// helix にないキー
					*t = "xx"
				} else {
					*t = keys[count]
					count++
				}
				// main processing
				// LSFT() とかの処理
				if strings.Contains(*t, "LSFT(") {
					tmp := (*t)[5:len(*t)]
					if _, ok := KEYMAP[tmp]; ok {
						tmp = KEYMAP[tmp]
					}
					*t = "(S)" + tmp
				}
				if strings.Contains(*t, "RSFT(") {
					tmp := (*t)[5:len(*t)]
					if _, ok := KEYMAP[tmp]; ok {
						tmp = KEYMAP[tmp]
					}
					*t = "(S)" + tmp
				}
				if strings.Contains(*t, "LCTL(") {
					tmp := (*t)[5:len(*t)]
					if _, ok := KEYMAP[tmp]; ok {
						tmp = KEYMAP[tmp]
					}
					*t = "(C)" + tmp
				}
				if strings.Contains(*t, "RCTL(") {
					tmp := (*t)[5:len(*t)]
					if _, ok := KEYMAP[tmp]; ok {
						tmp = KEYMAP[tmp]
					}
					*t = "(C)" + tmp
				}
				// LSFT_T(KC_A)
				if strings.Contains(*t, "LSFT_T(") {
					tmp := (*t)[7:len(*t)]
					if _, ok := KEYMAP[tmp]; ok {
						tmp = KEYMAP[tmp]
					}
					*t = "(S)" + tmp
				}
				if strings.Contains(*t, "LT(") {
					// LT( 5, KC_NO ) とかは LT(5とKC_NOに別れる
					// 半角空白やコンマは入らない
					layerNum := (*t)[3:len(*t)]
					tmp := keys[count]
					// これでtmpに数値だけ入る
					// keys[count]が次のキーコードになる
					if _, ok := KEYMAP[keys[count]]; ok {
						tmp = KEYMAP[keys[count]]
					}
					*t = "(l" + layerNum + ")" + tmp
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

func divNewLine(in string) []string { // {{{
	lines := strings.Split(in, "\n")
	length := len(lines)
	cc := make([]string, length)

	for i, line := range lines {
		tmp := line[strings.LastIndex(line, "LAYOUT(")+len("LAYOUT("):]
		cc[i] = tmp
	}
	return cc
}

// }}}

func cut(in string) string { // {{{
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
}

// }}}

func newLayer(numRow, numColumn int) [][]string { // {{{

	rs := make([][]string, numRow)
	for i := 0; i < len(rs); i++ {
		rs[i] = make([]string, numColumn)
	}
	return rs
}

// }}}

func ReadText(path string) string { // {{{

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
} // }}}

func ConvertNewline(str, nlcode string) string { // {{{
	// 改行コードを統一する
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
} // }}}

/* var input = ` // {{{
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
	[0] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, KC_Q, KC_W, KC_D, KC_F, KC_G, KC_Y, KC_S, KC_T, KC_R, KC_P, KC_BSPC, KC_LCTL, KC_A, KC_O, KC_E, KC_U, KC_I, KC_H, KC_J, KC_K, KC_L, KC_SCLN, KC_ENT, TO(1), KC_Z, KC_X, KC_C, KC_V, KC_B, TO(2), TO(2), KC_N, KC_M, KC_SLSH, KC_RO, KC_DOT, TO(1), KC_LALT, KC_NO, KC_NO, KC_LGUI, KC_LANG2, KC_SPC, TO(3), TO(4), KC_SPC, KC_HANJ, KC_NO, DF(0), DF(5), KC_NO),
	[1] = LAYOUT(LSFT(KC_ESC), KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, LSFT(KC_DEL), LSFT(KC_TAB), LSFT(KC_Q), LSFT(KC_W), LSFT(KC_D), LSFT(KC_F), LSFT(KC_G), LSFT(KC_Y), LSFT(KC_S), LSFT(KC_T), LSFT(KC_R), LSFT(KC_P), LSFT(KC_BSPC), LSFT(KC_LCTL), LSFT_T(KC_A), LSFT(KC_O), LSFT(KC_E), LSFT(KC_U), LSFT(KC_I), LSFT(KC_H), LSFT(KC_J), LSFT(KC_K), LSFT(KC_L), KC_NO, LSFT(KC_ENT), KC_TRNS, LSFT(KC_Z), LSFT(KC_X), LSFT(KC_C), LSFT(KC_V), LSFT(KC_B), KC_NO, KC_NO, LSFT(KC_N), LSFT(KC_M), KC_NO, KC_NO, KC_NO, KC_TRNS, LSFT(KC_LALT), KC_NO, KC_NO, LSFT(KC_LGUI), KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO),
	[2] = LAYOUT(KC_F1, KC_F2, KC_F3, KC_F4, KC_F5, KC_F6, KC_F7, KC_F8, KC_F9, KC_F10, KC_F11, KC_F12, RESET, KC_NO, KC_NO, KC_NO, KC_NO, RGB_TOG, KC_NO, KC_NO, KC_BTN1, KC_BTN2, KC_WH_U, KC_NO, EEP_RST, RGB_MOD, RGB_SPI, RGB_VAI, RGB_HUI, RGB_SAI, KC_MS_L, KC_MS_D, KC_MS_U, KC_MS_R, KC_WH_D, KC_NO, AG_NORM, RGB_RMOD, RGB_SPD, RGB_VAD, RGB_HUD, RGB_SAD, KC_TRNS, KC_TRNS, KC_LEFT, KC_DOWN, KC_UP, KC_RGHT, KC_NO, KC_NO, AG_SWAP, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_TRNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO),
	[3] = LAYOUT(KC_F1, KC_F2, KC_F3, KC_F4, KC_F5, KC_F6, KC_F7, KC_F8, KC_F9, KC_F10, KC_F11, KC_F12, RESET, KC_NO, KC_NO, KC_NO, KC_NO, RGB_TOG, KC_NO, KC_NO, KC_HOME, KC_END, KC_NO, KC_NO, EEP_RST, RGB_MOD, RGB_SPI, RGB_VAI, RGB_HUI, RGB_SAI, KC_LEFT, KC_DOWN, KC_UP, KC_RGHT, KC_NO, KC_NO, AG_NORM, RGB_RMOD, RGB_SPD, RGB_VAD, RGB_HUD, RGB_SAD, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, AG_SWAP, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_TRNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO),
	[4] = LAYOUT(KC_GRV, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, JP_EXLM, JP_QUES, JP_DQT, JP_QUOT, JP_HASH, JP_AMPR, JP_PIPE, JP_CIRC, JP_COLN, KC_NO, KC_BSPC, KC_LCTL, JP_MINS, JP_PLUS, JP_LPRN, JP_RPRN, JP_ASTR, JP_TILD, JP_LCBR, JP_RCBR, JP_COLN, KC_NO, KC_ENT, KC_NO, JP_EQL, JP_AT, JP_LBRC, JP_RBRC, JP_PERC, LCTL(KC_INS), KC_NO, JP_DLR, JP_LT, JP_GT, JP_UNDS, KC_COMM, KC_NO, KC_LALT, KC_NO, KC_NO, KC_LGUI, EISU, KC_SPC, LSFT(KC_INS), KC_TRNS, KC_SPC, KANA, KC_LGUI, KC_NO, KC_NO, KC_NO),
	[5] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, KC_Q, KC_W, KC_E, KC_R, KC_T, KC_Y, KC_U, KC_I, KC_O, KC_P, KC_BSPC, KC_LCTL, KC_A, KC_S, KC_D, KC_F, KC_G, KC_H, KC_J, KC_K, KC_L, KC_SCLN, KC_ENT, KC_LSFT, KC_Z, KC_X, KC_C, KC_V, KC_B, TO(2), TO(2), KC_N, KC_M, KC_SLSH, KC_RO, KC_DOT, KC_RSFT, KC_LALT, KC_NO, KC_NO, KC_LGUI, EISU, KC_SPC, TO(3), TO(4), KC_SPC, KANA, KC_NO, DF(0), KC_TRNS, KC_NO)
};
`
}}} */
