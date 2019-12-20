package main

import (
	"fmt"
	"strings"
)

// {{{
/*
helixキーボードでは
keymapの
(*,0,6), (*,0,7)
(*,1,6), (*,1,7)
(*,2,6), (*,2,7)
の6個が使えない
5 * 14 - 6 = 64
*/

const (
	RowNum    = 5
	ColumnNum = 14
)

var LayerNum int

var keymap [][][]string

// }}}

func main() {
	read(input)
}

// {{{

func read(in string) {
	// 前処理
	c := cut(in)            // 不要部分削除
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
				if j < 3 && (k == 6 || k == 7) {
					keymap[i][j][k] = "xx"
				} else {
					keymap[i][j][k] = keys[count]
					count++
				}
				if _, ok := KEYMAP[keymap[i][j][k]]; ok {
					keymap[i][j][k] = KEYMAP[keymap[i][j][k]]
				}
				//fmt.Printf("%3v,%3v,%3v : %10v\n", i, j, k, keymap[i][j][k])
			}
		}
	}

	for i := 0; i < LayerNum; i++ {
		fmt.Printf("layer %v\n", i)
		for j := 0; j < RowNum; j++ {
			for k := 0; k < ColumnNum; k++ {
				fmt.Printf("|")
				fmt.Printf(" %10v ", keymap[i][j][k])
				if k == ColumnNum-1 {
					fmt.Printf("|\n")
				}
			}
		}
		fmt.Printf("\n")
	}

}

func divNewLine(in string) []string {
	cc := make([]string, 6, 6)
	lines := strings.Split(in, "\n")

	for i, line := range lines {
		index := strings.LastIndex(line, "LAYOUT(")
		tmp := line[index+len("LAYOUT(") : len(line)]
		rs := []rune(tmp)
		fmt.Printf("%c\n", rs)
		n := len(rs)
		if rs[n-1] == ')' {
			cc[i] = string(rs[:n-1])
		} else if rs[n-2] == ')' && rs[n-1] == ',' {
			cc[i] = string(rs[:n-2])
		} else {
			panic("書式エラー : " + tmp)
		}
	}
	return cc
}

func cut(in string) string {
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
			if r == ' ' {
				continue
			}
			if r == '}' {
				return out[:len(out)-1]
			}
			out += string(r)
		}
		pre = r
	}
	return out[:len(out)-1]
}

func newLayer(numRow, numColumn int) [][]string {

	rs := make([][]string, numRow)
	for i := 0; i < len(rs); i++ {
		rs[i] = make([]string, numColumn)
	}
	return rs
}

var input = `
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
	[0] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, KC_Q, KC_W, KC_D, KC_F, KC_G, KC_Y, KC_S, KC_T, KC_R, KC_P, KC_BSPC, KC_LCTL, KC_A, KC_O, KC_E, KC_U, KC_I, KC_H, KC_J, KC_K, KC_L, KC_SCLN, KC_ENT, TO(1), KC_Z, KC_X, KC_C, KC_V, KC_B, TO(2), TO(2), KC_N, KC_M, KC_SLSH, KC_RO, KC_DOT, TO(1), KC_LALT, KC_NO, KC_NO, KC_LGUI, KC_LANG2, KC_SPC, TO(3), TO(4), KC_SPC, KC_HANJ, KC_NO, DF(0), DF(5), KC_NO),
	[1] = LAYOUT(LSFT(KC_ESC), KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, LSFT(KC_DEL), LSFT(KC_TAB), LSFT(KC_Q), LSFT(KC_W), LSFT(KC_D), LSFT(KC_F), LSFT(KC_G), LSFT(KC_Y), LSFT(KC_S), LSFT(KC_T), LSFT(KC_R), LSFT(KC_P), LSFT(KC_BSPC), LSFT(KC_LCTL), LSFT_T(KC_A), LSFT(KC_O), LSFT(KC_E), LSFT(KC_U), LSFT(KC_I), LSFT(KC_H), LSFT(KC_J), LSFT(KC_K), LSFT(KC_L), KC_NO, LSFT(KC_ENT), KC_TRNS, LSFT(KC_Z), LSFT(KC_X), LSFT(KC_C), LSFT(KC_V), LSFT(KC_B), KC_NO, KC_NO, LSFT(KC_N), LSFT(KC_M), KC_NO, KC_NO, KC_NO, KC_TRNS, LSFT(KC_LALT), KC_NO, KC_NO, LSFT(KC_LGUI), KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO),
	[2] = LAYOUT(KC_F1, KC_F2, KC_F3, KC_F4, KC_F5, KC_F6, KC_F7, KC_F8, KC_F9, KC_F10, KC_F11, KC_F12, RESET, KC_NO, KC_NO, KC_NO, KC_NO, RGB_TOG, KC_NO, KC_NO, KC_BTN1, KC_BTN2, KC_WH_U, KC_NO, EEP_RST, RGB_MOD, RGB_SPI, RGB_VAI, RGB_HUI, RGB_SAI, KC_MS_L, KC_MS_D, KC_MS_U, KC_MS_R, KC_WH_D, KC_NO, AG_NORM, RGB_RMOD, RGB_SPD, RGB_VAD, RGB_HUD, RGB_SAD, KC_TRNS, KC_TRNS, KC_LEFT, KC_DOWN, KC_UP, KC_RGHT, KC_NO, KC_NO, AG_SWAP, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_TRNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO),
	[3] = LAYOUT(KC_F1, KC_F2, KC_F3, KC_F4, KC_F5, KC_F6, KC_F7, KC_F8, KC_F9, KC_F10, KC_F11, KC_F12, RESET, KC_NO, KC_NO, KC_NO, KC_NO, RGB_TOG, KC_NO, KC_NO, KC_HOME, KC_END, KC_NO, KC_NO, EEP_RST, RGB_MOD, RGB_SPI, RGB_VAI, RGB_HUI, RGB_SAI, KC_LEFT, KC_DOWN, KC_UP, KC_RGHT, KC_NO, KC_NO, AG_NORM, RGB_RMOD, RGB_SPD, RGB_VAD, RGB_HUD, RGB_SAD, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, AG_SWAP, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_TRNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO),
	[4] = LAYOUT(KC_GRV, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, JP_EXLM, JP_QUES, JP_DQT, JP_QUOT, JP_HASH, JP_AMPR, JP_PIPE, JP_CIRC, JP_COLN, KC_NO, KC_BSPC, KC_LCTL, JP_MINS, JP_PLUS, JP_LPRN, JP_RPRN, JP_ASTR, JP_TILD, JP_LCBR, JP_RCBR, JP_COLN, KC_NO, KC_ENT, KC_NO, JP_EQL, JP_AT, JP_LBRC, JP_RBRC, JP_PERC, LCTL(KC_INS), KC_NO, JP_DLR, JP_LT, JP_GT, JP_UNDS, KC_COMM, KC_NO, KC_LALT, KC_NO, KC_NO, KC_LGUI, EISU, KC_SPC, LSFT(KC_INS), KC_TRNS, KC_SPC, KANA, KC_LGUI, KC_NO, KC_NO, KC_NO),
	[5] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, KC_Q, KC_W, KC_E, KC_R, KC_T, KC_Y, KC_U, KC_I, KC_O, KC_P, KC_BSPC, KC_LCTL, KC_A, KC_S, KC_D, KC_F, KC_G, KC_H, KC_J, KC_K, KC_L, KC_SCLN, KC_ENT, KC_LSFT, KC_Z, KC_X, KC_C, KC_V, KC_B, TO(2), TO(2), KC_N, KC_M, KC_SLSH, KC_RO, KC_DOT, KC_RSFT, KC_LALT, KC_NO, KC_NO, KC_LGUI, EISU, KC_SPC, TO(3), TO(4), KC_SPC, KANA, KC_NO, DF(0), KC_TRNS, KC_NO)
};
`

var KEYMAP = map[string]string{
	"JP_ZHTG":  "半/全",
	"JP_YEN":   "\\ / |",
	"JP_CIRC":  "^ / ~",
	"JP_AT":    "@ / `",
	"JP_LBRC":  "[ / {",
	"JP_COLN":  ": / *",
	"JP_RBRC":  "] / }",
	"JP_BSLS":  "\\ / _",
	"JP_MHEN":  "無変換",
	"JP_HENK":  "変換",
	"JP_KANA":  "カひロ",
	"JP_MKANA": "かな",
	"JP_MEISU": "英数",
	"JP_DQT":   "\"",
	"JP_AMPR":  "&",
	"JP_QUOT":  "'",
	"JP_LPRN":  "(",
	"JP_RPRN":  ")",
	"JP_EQL":   ":",
	"JP_TILD":  "~",
	"JP_PIPE":  "|",
	"JP_GRV":   "`",
	"JP_LCBR":  "{",
	"JP_PLUS":  "+",
	"JP_ASTR":  "*",
	"JP_RCBR":  "}",
	"JP_UNDS":  "_",
	"JP_MINS":  "-",
	"JP_SCLN":  ";",
	"JP_COMM":  ",",
	"JP_DOT":   ".",
	"JP_SLSH":  "/",
	"JP_EXLM":  "!",
	"JP_HASH":  "#",
	"JP_DLR":   "$",
	"JP_PERC":  "%",
	"JP_LT":    "<",
	"JP_GT":    ">",
	"JP_QUES":  "?",

	"KC_A":                     "a / A",
	"KC_B":                     "b / B",
	"KC_C":                     "c / C",
	"KC_D":                     "d / D",
	"KC_E":                     "e / E",
	"KC_F":                     "f / F",
	"KC_G":                     "g / G",
	"KC_H":                     "h / H",
	"KC_I":                     "i / I",
	"KC_J":                     "j / J",
	"KC_K":                     "k / K",
	"KC_L":                     "l / L",
	"KC_M":                     "m / M",
	"KC_N":                     "n / N",
	"KC_O":                     "o / O",
	"KC_P":                     "p / P",
	"KC_Q":                     "q / Q",
	"KC_R":                     "r / R",
	"KC_S":                     "s / S",
	"KC_T":                     "t / T",
	"KC_U":                     "u / U",
	"KC_V":                     "v / V",
	"KC_W":                     "w / W",
	"KC_X":                     "x / X",
	"KC_Y":                     "y / Y",
	"KC_Z":                     "z / Z",
	"KC_1":                     "1 / !",
	"KC_2":                     "2 / @",
	"KC_3":                     "3 / #",
	"KC_4":                     "4 / $",
	"KC_5":                     "5 / %",
	"KC_6":                     "6 / ^",
	"KC_7":                     "7 / &",
	"KC_8":                     "8 / *",
	"KC_9":                     "9 / (",
	"KC_0":                     "0 / )",
	"KC_ENTER":                 "Enter",
	"KC_ESCAPE":                "Escape",
	"KC_BSPACE":                "Backspace",
	"KC_TAB":                   "Tab",
	"KC_SPACE":                 "Space",
	"KC_MINUS":                 "- / _",
	"KC_EQUAL":                 "= / +",
	"KC_LBRACKET":              "[ / {",
	"KC_RBRACKET":              "] / }",
	"KC_BSLASH":                "\\ / |",
	"KC_NONUS_HASH":            "Non-US # / ~",
	"KC_SCOLON":                "; / :",
	"KC_QUOTE":                 "' / \"",
	"KC_GRAVE":                 "` / ~ (半 / 全)",
	"KC_COMMA":                 ", / <",
	"KC_DOT":                   ". / >",
	"KC_SLASH":                 "/ / ?",
	"KC_CAPSLOCK":              "Caps Lock",
	"KC_F1":                    "F1",
	"KC_F2":                    "F2",
	"KC_F3":                    "F3",
	"KC_F4":                    "F4",
	"KC_F5":                    "F5",
	"KC_F6":                    "F6",
	"KC_F7":                    "F7",
	"KC_F8":                    "F8",
	"KC_F9":                    "F9",
	"KC_F10":                   "F10",
	"KC_F11":                   "F11",
	"KC_F12":                   "F12",
	"KC_PSCREEN":               "Print Screen",
	"KC_SCROLLLOCK":            "Scroll Lock, (Brightness Down)",
	"KC_PAUSE":                 "Pause, (Brightness Up)",
	"KC_INSERT":                "Insert",
	"KC_HOME":                  "Home",
	"KC_PGUP":                  "Page Up",
	"KC_DELETE":                "Forward Delete",
	"KC_END":                   "End",
	"KC_PGDOWN":                "Page Down",
	"KC_RIGHT":                 "Right Arrow",
	"KC_LEFT":                  "Left Arrow",
	"KC_DOWN":                  "Down Arrow",
	"KC_UP":                    "Up Arrow",
	"KC_NUMLOCK":               "Keypad Num Lock / Clear",
	"KC_KP_SLASH":              "Keypad /",
	"KC_KP_ASTERISK":           "Keypad *",
	"KC_KP_MINUS":              "Keypad -",
	"KC_KP_PLUS":               "Keypad +",
	"KC_KP_ENTER":              "Keypad Enter",
	"KC_KP_1":                  "Keypad 1 / End",
	"KC_KP_2":                  "Keypad 2 / Down Arrow",
	"KC_KP_3":                  "Keypad 3 / Page Down",
	"KC_KP_4":                  "Keypad 4 / Left Arrow",
	"KC_KP_5":                  "Keypad 5 ",
	"KC_KP_6":                  "Keypad 6 / Right Arrow",
	"KC_KP_7":                  "Keypad 7 / Home",
	"KC_KP_8":                  "Keypad 8 / Up Arrow",
	"KC_KP_9":                  "Keypad 9 / Page Up",
	"KC_KP_0":                  "Keypad 0 / Insert",
	"KC_KP_DOT":                "Keypad . / Delete",
	"KC_NONUS_BSLASH":          "Non-US \\ / |",
	"KC_APPLICATION":           "Application (Windows Menu Key)",
	"KC_POWER":                 "Power (macOS)",
	"KC_KP_EQUAL":              "Keypad =",
	"KC_F13":                   "F13",
	"KC_F14":                   "F14",
	"KC_F15":                   "F15",
	"KC_F16":                   "F16",
	"KC_F17":                   "F17",
	"KC_F18":                   "F18",
	"KC_F19":                   "F19",
	"KC_F20":                   "F20",
	"KC_F21":                   "F21",
	"KC_F22":                   "F22",
	"KC_F23":                   "F23",
	"KC_F24":                   "F24",
	"KC_EXECUTE":               "Execute",
	"KC_HELP":                  "Help",
	"KC_MENU":                  "Menu",
	"KC_SELECT":                "Select",
	"KC_STOP":                  "Stop",
	"KC_AGAIN":                 "Again",
	"KC_UNDO":                  "Undo",
	"KC_CUT":                   "Cut",
	"KC_COPY":                  "Copy",
	"KC_PASTE":                 "Paste",
	"KC_FIND":                  "Find",
	"KC__MUTE":                 "Mute (macOS)",
	"KC__VOLUP":                "Volume Up (macOS)",
	"KC__VOLDOWN":              "Volume Down (macOS)",
	"KC_LOCKING_CAPS":          "Caps Lock",
	"KC_LOCKING_NUM":           "Num Lock",
	"KC_LOCKING_SCROLL":        "Scroll Lock",
	"KC_KP_COMMA":              "Keypad ,",
	"KC_KP_EQUAL_AS400":        "Keypad = on AS/400 keyboards",
	"KC_INT1":                  "JIS \\ / _",
	"KC_INT2":                  "JIS Katakana/Hiragana",
	"KC_INT3":                  "JIS \\ / |",
	"KC_INT4":                  "JIS Henkan",
	"KC_INT5":                  "JIS Muhenkan",
	"KC_INT6":                  "JIS Numpad ,",
	"KC_INT7":                  "International 7",
	"KC_INT8":                  "International 8",
	"KC_INT9":                  "International 9",
	"KC_LANG1":                 "Hangul/English",
	"KC_LANG2":                 "Hanja",
	"KC_LANG3":                 "JIS Katakana",
	"KC_LANG4":                 "JIS Hiragana",
	"KC_LANG5":                 "JIS Zenkaku/Hankaku",
	"KC_LANG6":                 "Language 6",
	"KC_LANG7":                 "Language 7",
	"KC_LANG8":                 "Language 8",
	"KC_LANG9":                 "Language 9",
	"KC_ALT_ERASE":             "Alternate Erase",
	"KC_SYSREQ":                "SysReq/Attention",
	"KC_CANCEL":                "Cancel",
	"KC_CLEAR":                 "Clear",
	"KC_PRIOR":                 "Prior",
	"KC_RETURN":                "Return",
	"KC_SEPARATOR":             "Separator",
	"KC_OUT":                   "Out",
	"KC_OPER":                  "Oper",
	"KC_CLEAR_AGAIN":           "Clear/Again",
	"KC_CRSEL":                 "CrSel/Props",
	"KC_EXSEL":                 "ExSel",
	"KC_LCTRL":                 "Left Control",
	"KC_LSHIFT":                "Left Shift",
	"KC_LALT":                  "Left Alt",
	"KC_LGUI":                  "Left GUI (Windows/Command/Meta key)",
	"KC_RCTRL":                 "Right Control",
	"KC_RSHIFT":                "Right Shift",
	"KC_RALT":                  "Right Alt (AltGr)",
	"KC_RGUI":                  "Right GUI (Windows/Command/Meta key)",
	"KC_SYSTEM_POWER":          "System Power Down",
	"KC_SYSTEM_SLEEP":          "System Sleep",
	"KC_SYSTEM_WAKE":           "System Wake",
	"KC_AUDIO_MUTE":            "Mute",
	"KC_AUDIO_VOL_UP":          "Volume Up",
	"KC_AUDIO_VOL_DOWN":        "Volume Down",
	"KC_MEDIA_NEXT_TRACK":      "Next Track",
	"KC_MEDIA_PREV_TRACK":      "Previous Track",
	"KC_MEDIA_STOP":            "Stop Track (Windows)",
	"KC_MEDIA_PLAY_PAUSE":      "Play/Pause Track",
	"KC_MEDIA_SELECT":          "Launch Media Player (Windows)",
	"KC_MEDIA_EJECT":           "Eject (macOS)",
	"KC_MAIL":                  "Launch Mail (Windows)",
	"KC_CALCULATOR":            "Launch Calculator (Windows)",
	"KC_MY_COMPUTER":           "Launch My Computer (Windows)",
	"KC_WWW_SEARCH":            "Browser Search (Windows)",
	"KC_WWW_HOME":              "Browser Home (Windows)",
	"KC_WWW_BACK":              "Browser Back (Windows)",
	"KC_WWW_FORWARD":           "Browser Forward (Windows)",
	"KC_WWW_STOP":              "Browser Stop (Windows)",
	"KC_WWW_REFRESH":           "Browser Refresh (Windows)",
	"KC_WWW_FAVORITES":         "Browser Favorites (Windows)",
	"KC_MEDIA_FAST_FORWARD":    "Next Track (macOS)",
	"KC_MEDIA_REWIND":          "Previous Track (macOS)",
	"KC_BRIGHTNESS_UP":         "Brightness Up",
	"KC_BRIGHTNESS_DOWN":       "Brightness Down",
	"KC_ENT":                   "Return (Enter)",
	"KC_ESC":                   "Escape",
	"KC_BSPC":                  "Delete (Backspace)",
	"KC_SPC":                   "Spacebar",
	"KC_MINS":                  "- / _",
	"KC_EQL":                   "= / +",
	"KC_LBRC":                  "[ / {",
	"KC_RBRC":                  "] / }",
	"KC_BSLS":                  "\\ / |",
	"KC_NUHS":                  "Non-US # / ~",
	"KC_SCLN":                  "; / :",
	"KC_QUOT":                  "' / \"",
	"KC_GRV, KC_ZKHK":          "` / ~, JIS Zenkaku/Hankaku",
	"KC_COMM":                  ", / <",
	"KC_SLSH":                  "/ / ?",
	"KC_CLCK, KC_CAPS":         "Caps Lock",
	"KC_PSCR":                  "Print Screen",
	"KC_SLCK, KC_BRMD":         "Scroll Lock, Brightness Down (macOS)",
	"KC_PAUS, KC_BRK, KC_BRMU": "Pause, Brightness Up (macOS)",
	"KC_INS":                   "Insert",
	"KC_DEL":                   "Forward Delete",
	"KC_PGDN":                  "Page Down",
	"KC_RGHT":                  "Right Arrow",
	"KC_NLCK":                  "Keypad Num Lock / Clear",
	"KC_PSLS":                  "Keypad /",
	"KC_PAST":                  "Keypad *",
	"KC_PMNS":                  "Keypad -",
	"KC_PPLS":                  "Keypad +",
	"KC_PENT":                  "Keypad Enter",
	"KC_P1":                    "Keypad 1 and End",
	"KC_P2":                    "Keypad 2 and Down Arrow",
	"KC_P3":                    "Keypad 3 and Page Down",
	"KC_P4":                    "Keypad 4 and Left Arrow",
	"KC_P5":                    "Keypad 5",
	"KC_P6":                    "Keypad 6 and Right Arrow",
	"KC_P7":                    "Keypad 7 and Home",
	"KC_P8":                    "Keypad 8 and Up Arrow",
	"KC_P9":                    "Keypad 9 and Page Up",
	"KC_P0":                    "Keypad 0 and Insert",
	"KC_PDOT":                  "Keypad . and Delete",
	"KC_NUBS":                  "Non-US \\ / |",
	"KC_APP":                   "Application (Windows Menu Key)",
	"KC_PEQL":                  "Keypad =",
	"KC_EXEC":                  "Execute",
	"KC_SLCT":                  "Select",
	"KC_AGIN":                  "Again",
	"KC_PSTE":                  "Paste",
	"KC_LCAP":                  "Locking Caps Lock",
	"KC_LNUM":                  "Locking Num Lock",
	"KC_LSCR":                  "Locking Scroll Lock",
	"KC_PCMM":                  "Keypad ,",
	"KC_RO":                    "JIS \\ / _",
	"KC_KANA":                  "JIS Katakana/Hiragana",
	"KC_JYEN":                  "JIS \\ / |",
	"KC_HENK":                  "JIS Henkan",
	"KC_MHEN":                  "JIS Muhenkan",
	"KC_HAEN":                  "Hangul/English",
	"KC_HANJ":                  "Hanja",
	"KC_ERAS":                  "Alternate Erase",
	"KC_CLR":                   "Clear",
	"KC_LCTL":                  "Left Control",
	"KC_LSFT":                  "Left Shift",
	"KC_LCMD, KC_LWIN":         "Left GUI (Windows/Command/Meta key)",
	"KC_RCTL":                  "Right Control",
	"KC_RSFT":                  "Right Shift",
	"KC_ALGR":                  "Right Alt (AltGr)",
	"KC_RCMD, KC_RWIN":         "Right GUI (Windows/Command/Meta key)",
	"KC_PWR":                   "System Power Down",
	"KC_SLEP":                  "System Sleep",
	"KC_WAKE":                  "System Wake",
	"KC_MUTE":                  "Mute",
	"KC_VOLU":                  "Volume Up",
	"KC_VOLD":                  "Volume Down",
	"KC_MNXT":                  "Next Track",
	"KC_MPRV":                  "Previous Track",
	"KC_MSTP":                  "Stop Track (Windows)",
	"KC_MPLY":                  "Play/Pause Track",
	"KC_MSEL":                  "Launch Media Player (Windows)",
	"KC_EJCT":                  "Eject (macOS)",
	"KC_CALC":                  "Launch Calculator (Windows)",
	"KC_MYCM":                  "Launch My Computer (Windows)",
	"KC_WSCH":                  "Browser Search (Windows)",
	"KC_WHOM":                  "Browser Home (Windows)",
	"KC_WBAK":                  "Browser Back (Windows)",
	"KC_WFWD":                  "Browser Forward (Windows)",
	"KC_WSTP":                  "Browser Stop (Windows)",
	"KC_WREF":                  "Browser Refresh (Windows)",
	"KC_WFAV":                  "Browser Favorites (Windows)",
	"KC_MFFD":                  "N Track",
	"KC_MRWD":                  "P Track",
	"KC_BRIU":                  "Bri Up",
	"KC_BRID":                  "Bri Down",
	"KC_NO":                    "---",
}

//}}}}
