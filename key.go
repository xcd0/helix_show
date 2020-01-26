package main

var KEYMAP_FUNC = map[string]string{
	"LSFT(":   "S ",
	"RSFT(":   "S ",
	"LCTL(":   "C ",
	"RCTL(":   "C ",
	"LALT(":   "A ",
	"RALT(":   "A ",
	"LSFT_T(": "S ",
	"DF(":     "DF.",
	"TO(":     "TO.",
	"TG(":     "TG.",
	"TT(":     "TT.",
	"MO(":     "MO.",
	"ANY(":    "",
}

var KEYMAP = map[string]string{
	// QMKの機能など {{{

	"KC_TRANSPARENT": "Trans", // Use the next lowest non-transparent key
	"KC_TRNS":        "Trans", // Use the next lowest non-transparent key
	"_______":        "Trans", // Use the next lowest non-transparent key

	"KC_NO":   "--",      // NOOP
	"XXXXXXX": "--",      // NOOP
	"RESET":   "Reset",   // リセット
	"RGB_TOG": "LED SW",  // LEDのON/OFF
	"EEP_RST": "E Reset", // EEPROM のリセット

	"RGB_MOD":  "Effect+", // カラーエフェクト
	"RGB_RMOD": "Effect-", // カラーエフェクト
	"RGB_SPI":  "Eff Sp+", // カラーエフェクトの速さ
	"RGB_SPD":  "Eff Sp-", // カラーエフェクトの速さ
	"RGB_VAI":  "Shodo+",  // 照度
	"RGB_VAD":  "Shodo-",  // 照度
	"RGB_HUI":  "Meido+",  // 明度
	"RGB_HUD":  "Meido-",  // 明度
	"RGB_SAI":  "Saido+",  // 彩度
	"RGB_SAD":  "Saido-",  // 彩度

	"AG_NORM": "mac", // 英数かな とかのためのOS切り替え mac
	"AG_SWAP": "win", // 英数かな とかのためのOS切り替え win

	// https://beta.docs.qmk.fm/features/feature_mouse_keys
	// rules.mk で MOUSEKEY_ENABLE = yes にする必要がある
	"KC_MS_UP":       "Mouse U",       // マウス ↑
	"KC_MS_U":        "Mouse U",       // マウス ↑
	"KC_MS_DOWN":     "Mouse D",       // マウス ↓
	"KC_MS_D":        "Mouse D",       // マウス ↓
	"KC_MS_RIGHT":    "Mouse R",       // マウス →
	"KC_MS_R":        "Mouse R",       // マウス →
	"KC_MS_LEFT":     "Mouse L",       // マウス ←
	"KC_MS_L":        "Mouse L",       // マウス ←
	"KC_MS_WH_UP":    "Wheel U",       // マウス ホイール↑
	"KC_WH_U":        "Wheel U",       // マウス ホイール↑
	"KC_WH_WH_DOWN":  "Wheel D",       // マウス ホイール↓
	"KC_WH_D":        "Wheel D",       // マウス ホイール↓
	"KC_MS_WH_LEFT":  "Wheel L",       // Move wheel left
	"KC_WH_L":        "Wheel L",       // Move wheel left
	"KC_MS_WH_RIGHT": "Wheel R",       // Move wheel right
	"KC_WH_R":        "Wheel R",       // Move wheel right
	"KC_MS_BTN1":     "Mouse 1",       // Press button 1
	"KC_BTN1":        "L Click",       // マウス 左クリック "KC_BTN1":"Mouse 1", // Press button 1
	"KC_MS_BTN2":     "Mouse 2",       // Press button 2
	"KC_BTN2":        "R Click",       // マウス 右クリック "KC_BTN2":"Mouse 2", // Press button 2
	"KC_MS_BTN3":     "Mouse 3",       // Press button 3
	"KC_BTN3":        "C Click",       // マウス 中クリック "KC_BTN3":"Mouse 3",// Press button 3
	"KC_MS_BTN4":     "Mouse 4",       // Press button 4
	"KC_BTN4":        "Mouse 4",       // Press button 4
	"KC_MS_BTN5":     "Mouse 5",       // Press button 5
	"KC_BTN5":        "Mouse 5",       // Press button 5
	"KC_MS_ACCEL0":   "Mouse Speed 0", // Set speed to 0
	"KC_ACL0":        "Mouse Speed 0", // Set speed to 0
	"KC_MS_ACCEL1":   "Mouse Speed 1", // Set speed to 1
	"KC_ACL1":        "Mouse Speed 1", // Set speed to 1
	"KC_MS_ACCEL2":   "Mouse Speed 2", // Set speed to 2
	"KC_ACL2":        "Mouse Speed 2", // Set speed to 2

	// }}}

	// JP {{{
	// keymap_jp.hをインクルードしてるとき使える
	// qmk_firmware/quantum/keymap_extras/keymap_jp.h にある
	// Aliases for shifted symbols
	"JP_ZHTG":  "半/全",   // KC_GRV  // hankaku/zenkaku|kanzi
	"JP_YEN":   "\\  |", // KC_INT3 // yen, |
	"JP_CIRC":  "^  ~",  // KC_EQL  // ^, ~
	"JP_AT":    "@  `",  // KC_LBRC // @, `
	"JP_LBRC":  "[  {",  // KC_RBRC // [, {
	"JP_COLN":  ":  *",  // KC_QUOT // :, *
	"JP_RBRC":  "]  }",  // KC_NUHS // ], }
	"JP_BSLS":  "\\  _", // KC_INT1 // \, _
	"JP_MHEN":  "無変換",   // KC_INT5 // muhenkan
	"JP_HENK":  "変換",    // KC_INT4 // henkan
	"JP_KANA":  "カひロ",   // KC_INT2 // katakana/hiragana|ro-mazi
	"JP_MKANA": "かな",    // KC_LANG1 //kana on MacOSX
	"JP_MEISU": "英数",    // KC_LANG2 //eisu on MacOSX
	"JP_DQT":   "\"",    // LSFT(KC_2)    // "
	"JP_AMPR":  "&",     // LSFT(KC_6)    // &
	"JP_QUOT":  "'",     // LSFT(KC_7)    // '
	"JP_LPRN":  "(",     // LSFT(KC_8)    // (
	"JP_RPRN":  ")",     // LSFT(KC_9)    // )
	"JP_EQL":   "=",     // LSFT(KC_MINS) // =
	"JP_TILD":  "~",     // LSFT(JP_CIRC) // ~
	"JP_PIPE":  "|",     // LSFT(JP_YEN)  // |
	"JP_GRV":   "`",     // LSFT(JP_AT)   // `
	"JP_LCBR":  "{",     // LSFT(JP_LBRC) // {
	"JP_PLUS":  "+",     // LSFT(KC_SCLN) // +
	"JP_ASTR":  "*",     // LSFT(JP_COLN) // *
	"JP_RCBR":  "}",     // LSFT(JP_RBRC) // }
	"JP_UNDS":  "_",     // LSFT(JP_BSLS) // _
	// These symbols are correspond to US101-layout.
	"JP_MINS": "-", // KC_MINS // -
	"JP_SCLN": ";", // KC_SCLN // ;
	"JP_COMM": ",", // KC_COMM // ,
	"JP_DOT":  ".", // KC_DOT  // .
	"JP_SLSH": "/", // KC_SLSH // /
	// shifted
	"JP_EXLM": "!", // KC_EXLM // !
	"JP_HASH": "#", // KC_HASH // #
	"JP_DLR":  "$", // KC_DLR  // $
	"JP_PERC": "%", // KC_PERC // %
	"JP_LT":   "<", // KC_LT   // <
	"JP_GT":   ">", // KC_GT   // >
	"JP_QUES": "?", // KC_QUES // ?
	// }}}

	// 基本 {{{1
	// アルファベット {{{
	"KC_A": "a",
	"KC_B": "b",
	"KC_C": "c",
	"KC_D": "d",
	"KC_E": "e",
	"KC_F": "f",
	"KC_G": "g",
	"KC_H": "h",
	"KC_I": "i",
	"KC_J": "j",
	"KC_K": "k",
	"KC_L": "l",
	"KC_M": "m",
	"KC_N": "n",
	"KC_O": "o",
	"KC_P": "p",
	"KC_Q": "q",
	"KC_R": "r",
	"KC_S": "s",
	"KC_T": "t",
	"KC_U": "u",
	"KC_V": "v",
	"KC_W": "w",
	"KC_X": "x",
	"KC_Y": "y",
	"KC_Z": "z",
	// }}}

	// 数字 {{{
	"KC_1": "1",
	"KC_2": "2",
	"KC_3": "3",
	"KC_4": "4",
	"KC_5": "5",
	"KC_6": "6",
	"KC_7": "7",
	"KC_8": "8",
	"KC_9": "9",
	"KC_0": "0",
	// }}}

	// エンターとか {{{2
	"KC_ENTER":  "Enter",
	"KC_ENT":    "Enter", // "Return (Enter)",
	"KC_ESCAPE": "Esc",
	"KC_ESC":    "Esc", // "Escape",
	"KC_BSPACE": "Bk",
	"KC_BSPC":   "Bk", // "Delete (Backspace)",
	"KC_TAB":    "Tab",
	"KC_SPACE":  "Spc",
	"KC_SPC":    "Spc", // "Spacebar",

	"KC_LCTRL": "LCtrl", // "Left Control",
	"KC_LCTL":  "LCtrl",
	"KC_RCTRL": "RCtrl", // "Right Control",
	"KC_RCTL":  "RCtrl",

	"KC_LSHIFT": "LShft", // "Left Shift",
	"KC_LSFT":   "LShift",
	"KC_RSHIFT": "RShft", // "Right Shift",
	"KC_RSFT":   "RShift",

	"KC_LALT": "LAlt", // "Left Alt",
	"KC_ALGR": "RAlt", //  (AltGr)",
	"KC_RALT": "RAlt", // "Right Alt (AltGr)",

	"KC_LGUI": "LGUI", // "Left GUI (Windows/Command/Meta key)",
	"KC_LCMD": "LGUI", // (Windows/Command/Meta key)",
	"KC_LWIN": "LGUI", // (Windows/Command/Meta key)",
	"KC_RGUI": "RGUI", // "Right GUI (Windows/Command/Meta key)",
	"KC_RCMD": "RGUI", //  (Windows/Command/Meta key)",
	"KC_RWIN": "RGUI", //  (Windows/Command/Meta key)",
	// }}}2

	// 記号{{{2
	"KC_MINS":       "-  _",
	"KC_MINUS":      "-  _",
	"KC_EQUAL":      "=  +",
	"KC_LBRACKET":   "[  {",
	"KC_LBRC":       "[  {",
	"KC_RBRACKET":   "]  }",
	"KC_RBRC":       "]  }",
	"KC_LPRN":       "(",
	"KC_RPRN":       ")",
	"KC_LT":         "<",
	"KC_GT":         ">",
	"KC_AMPR":       "&",
	"KC_LCBR":       "{",
	"KC_RCBR":       "}",
	"KC_PIPE":       "|",
	"KC_EQL":        "=  +",
	"KC_BSLASH":     "\\  |",
	"KC_BSLS":       "\\  |",
	"KC_NONUS_HASH": "#  ~",
	"KC_NUHS":       "#  ~",
	"KC_SCOLON":     ";  :",
	"KC_SCLN":       ";  :",
	"KC_QUOTE":      "'  \"",
	"KC_QUOT":       "'  \"",
	"KC_GRV":        "`  ~", // 半/全
	"KC_GRAVE":      "`  ~", // 半/全
	"KC_ZKHK":       "全/半",
	"KC_COMMA":      ",  <", // 小なり 「<」 「ぐ」じゃない
	"KC_COMM":       ",  <", // 小なり 「<」 「ぐ」じゃない
	"KC_DOT":        ".  >",
	"KC_SLSH":       "/  ?",
	"KC_SLASH":      "/  ?",
	// }}}2

	// Function {{{
	"KC_F1":  "F1",
	"KC_F2":  "F2",
	"KC_F3":  "F3",
	"KC_F4":  "F4",
	"KC_F5":  "F5",
	"KC_F6":  "F6",
	"KC_F7":  "F7",
	"KC_F8":  "F8",
	"KC_F9":  "F9",
	"KC_F10": "F10",
	"KC_F11": "F11",
	"KC_F12": "F12",
	"KC_F13": "F13",
	"KC_F14": "F14",
	"KC_F15": "F15",
	"KC_F16": "F16",
	"KC_F17": "F17",
	"KC_F18": "F18",
	"KC_F19": "F19",
	"KC_F20": "F20",
	"KC_F21": "F21",
	"KC_F22": "F22",
	"KC_F23": "F23",
	"KC_F24": "F24",
	// }}}
	// }}}1

	// その他 {{{1

	// テンキー {{{2
	"KC_NUMLOCK":        "_NL", // "Keypad Num Lock / Clear",
	"KC_KP_SLASH":       "_/",
	"KC_KP_ASTERISK":    "_*",
	"KC_KP_MINUS":       "_-",
	"KC_KP_PLUS":        "_+",
	"KC_KP_ENTER":       "_Enter",
	"KC_KP_1":           "_1_End",
	"KC_P1":             "_1_End",
	"KC_KP_2":           "_2_Down",
	"KC_P2":             "_2_Down",
	"KC_KP_3":           "_3_PgD",
	"KC_P3":             "_3_PgD",
	"KC_KP_4":           "_4_Left",
	"KC_P4":             "_4_Left",
	"KC_KP_5":           "_5",
	"KC_P5":             "_5",
	"KC_KP_6":           "_6_Right",
	"KC_P6":             "_6_Right",
	"KC_KP_7":           "_7_Home",
	"KC_P7":             "_7_Home",
	"KC_KP_8":           "_8_Up",
	"KC_P8":             "_8_Up",
	"KC_KP_9":           "_9_PgU",
	"KC_P9":             "_9_PgU",
	"KC_KP_0":           "_0_Ins",
	"KC_P0":             "_0_Ins",
	"KC_KP_DOT":         "_ . Dlt",
	"KC_PDOT":           "_ . Dlt",
	"KC_KP_EQUAL":       "_ =",
	"KC_PEQL":           "_ =",
	"KC_KP_EQUAL_AS400": "_ = on AS/400 keyboards",
	"KC_KP_COMMA":       "_ ,",
	"KC_PCMM":           "_,",
	"KC_NLCK":           "_NLClear",
	"KC_PSLS":           "_ /",
	"KC_PAST":           "_ *",
	"KC_PMNS":           "_ -",
	"KC_PPLS":           "_ +",
	"KC_PENT":           "_ Enter",
	// }}}2
	"KC_INSERT": "Insert",
	"KC_INS":    "Insert",
	"KC_DELETE": "Delete", // "Forward Delete",
	"KC_DEL":    "Delete",
	"KC_HOME":   "Home",
	"KC_END":    "End",
	"KC_PGUP":   "PageU", // Page Up
	"KC_PGDOWN": "PageD", // Page Down
	"KC_PGDN":   "PageD", // Page Down

	"KC_RIGHT":        "Right", // Right Arrow",
	"KC_RGHT":         "Right",
	"KC_LEFT":         "Left",  // Left Arrow",
	"KC_DOWN":         "Down",  // Down Arrow",
	"KC_UP":           "Up",    // Up Arrow",
	"KC_NONUS_BSLASH": "\\  |", // "Non-US \\ / |",
	"KC_NUBS":         "\\  |", // Non-US \ and |

	"KC_SYSTEM_POWER": "Power", // "System Power Down",
	"KC_POWER":        "Power", // "Power (macOS)",
	"KC_PWR":          "Power",
	"KC_SYSTEM_SLEEP": "Sleep", // "System Sleep",
	"KC_SLEP":         "Sleep",
	"KC_SYSTEM_WAKE":  "Wake", // "System Wake",
	"KC_WAKE":         "Wake",

	"KC_CAPSLOCK": "CapsLck", // Caps Lock
	"KC_CLCK":     "CapsLk",  // Caps Lock
	"KC_CAPS":     "CapsLk",  // Caps Lock

	"KC_PSCREEN": "PrntScn", // Print Screen
	"KC_PSCR":    "PrntScn", // Print Screen

	"KC_SCROLLLOCK": "ScL BrD", // Scroll Lock, (Brightness Down)
	"KC_SLCK":       "ScrllLk", // Scroll Lock, (Brightness Down)

	"KC_INT1": "\\  _", // "JIS \\ / _",
	"KC_RO":   "\\  _",
	"KC_INT2": "カひ",    // "JIS Katakana/Hiragana",
	"KC_KANA": "カひ",    // 全角文字はずれるかもなのでローマ字のほうがいいかも
	"KC_INT3": "\\  |", // "JIS \\ / |",
	"KC_JYEN": "\\  |",
	"KC_INT4": "HenKan", // "JIS Henkan",
	"KC_HENK": "HenKan",
	"KC_INT5": "MuHen", // "JIS Muhenkan",
	"KC_MHEN": "MuHen",
	"KC_INT6": "J Numpad ,", // "JIS Numpad ,",

	// }}}2

	// 使わなさそう {{{2
	"KC_INT7": "International 7",
	"KC_INT8": "International 8",
	"KC_INT9": "International 9",

	"KC_LANG1": "Kana",
	"KC_LANG2": "Eisu",
	"KC_HANJ":  "Hanja",          // ハングルと英数を切り替えるようなやつっぽい
	"KC_HAEN":  "Hangul/English", // ハングルでの漢字入力っぽい
	"KC_LANG3": "JIS Katakana",
	"KC_LANG4": "JIS Hiragana",
	"KC_LANG5": "JIS Zenkaku/Hankaku",
	"KC_LANG6": "Language 6",
	"KC_LANG7": "Language 7",
	"KC_LANG8": "Language 8",
	"KC_LANG9": "Language 9",

	"KC_MEDIA_NEXT_TRACK":   "Next Track",
	"KC_MNXT":               "Next Track",
	"KC_MEDIA_FAST_FORWARD": "Next Track (macOS)",
	"KC_MFFD":               "Next Track",
	"KC_MEDIA_PREV_TRACK":   "Previous Track",
	"KC_MPRV":               "Previous Track",
	"KC_MEDIA_REWIND":       "Previous Track (macOS)",
	"KC_MRWD":               "Previous Track",
	"KC_MEDIA_STOP":         "Stop Track (Windows)",
	"KC_MSTP":               "Stop Track (Windows)",
	"KC_MEDIA_PLAY_PAUSE":   "Play/Pause Track",
	"KC_MPLY":               "Play/Pause Track",
	"KC_AUDIO_MUTE":         "Mute",
	"KC_MUTE":               "Mute",
	"KC__MUTE":              "Mute", // Mute (macOS)

	"KC_AUDIO_VOL_UP":   "VolUp",   // Volume Up
	"KC__VOLUP":         "VolUp",   // Volume Up (macOS)
	"KC_VOLU":           "VolUp",   // Volume Up
	"KC_AUDIO_VOL_DOWN": "VolDown", // Volume Down
	"KC__VOLDOWN":       "VolDown", // Volume Down (macOS)
	"KC_VOLD":           "VolDown", // Volume Down

	"KC_BRIGHTNESS_UP": "BrightU",  // Brightness Up
	"KC_BRIU":          "BrightU",  // Brightness Up
	"KC_PAUSE":         "Pau BrU",  // Pause, Brightness Up
	"KC_PAUS":          "Pau  BrU", // Pause, Brightness Up (macOS)
	"KC_BRK":           "Pau  BrU", // Pause, Brightness Up (macOS)
	"KC_BRMU":          "Pau  BrU", // Pause, Brightness Up (macOS)

	"KC_BRIGHTNESS_DOWN": "BrightD", // Brightness Down
	"KC_BRID":            "BrightD", // Brightness Down
	"KC_BRMD":            "BrightD", // Brightness Down (macOS)

	"KC_APPLICATION": "Win", // Application (Windows Menu Key)",
	"KC_APP":         "APP", // "Application (Windows Menu Key)",

	"KC_MEDIA_SELECT":  "Launch Media Player (Windows)",
	"KC_MSEL":          "Launch Media Player (Windows)",
	"KC_MEDIA_EJECT":   "Eject (macOS)",
	"KC_EJCT":          "Eject (macOS)",
	"KC_MAIL":          "Launch Mail (Windows)",
	"KC_CALCULATOR":    "Launch Calculator (Windows)",
	"KC_CALC":          "Launch Calculator (Windows)",
	"KC_MY_COMPUTER":   "Launch My Computer (Windows)",
	"KC_MYCM":          "Launch My Computer (Windows)",
	"KC_WWW_SEARCH":    "Browser Search (Windows)",
	"KC_WSCH":          "Browser Search (Windows)",
	"KC_WWW_HOME":      "Browser Home (Windows)",
	"KC_WHOM":          "Browser Home (Windows)",
	"KC_WWW_BACK":      "Browser Back (Windows)",
	"KC_WBAK":          "Browser Back (Windows)",
	"KC_WWW_FORWARD":   "Browser Forward (Windows)",
	"KC_WFWD":          "Browser Forward (Windows)",
	"KC_WWW_STOP":      "Browser Stop (Windows)",
	"KC_WSTP":          "Browser Stop (Windows)",
	"KC_WWW_REFRESH":   "Browser Refresh (Windows)",
	"KC_WREF":          "Browser Refresh (Windows)",
	"KC_WWW_FAVORITES": "Browser Favorites (Windows)",
	"KC_WFAV":          "Browser Favorites (Windows)",
	// }}}2

	// なにこれら {{{2

	"KC_LOCKING_CAPS":   "LCpsLck", // Locking CapsLock   謎のキー 何に使うのか...
	"KC_LCAP":           "LCpsLck", // Locking CapsLock   謎のキー 何に使うのか...
	"KC_LOCKING_NUM":    "LNumLck", // Locking NumLock    謎のキー 何に使うのか...
	"KC_LNUM":           "LNumLck", // Locking NumLock    謎のキー 何に使うのか...
	"KC_LOCKING_SCROLL": "LScrLck", // Locking ScreenLock 謎のキー 何に使うのか...
	"KC_LSCR":           "LScrLck", // Locking ScreenLock 謎のキー 何に使うのか...

	"KC_EXECUTE": "Execute",
	"KC_EXEC":    "Execute",
	"KC_HELP":    "Help",
	"KC_MENU":    "Menu",
	"KC_SELECT":  "Select",
	"KC_SLCT":    "Select",
	"KC_STOP":    "Stop",
	"KC_AGAIN":   "Again",
	"KC_AGIN":    "Again",
	"KC_UNDO":    "Undo",
	"KC_CUT":     "Cut",
	"KC_COPY":    "Copy",
	"KC_PASTE":   "Paste",
	"KC_PSTE":    "Paste",
	"KC_FIND":    "Find",

	"KC_ALT_ERASE":   "Alternate Erase",
	"KC_ERAS":        "Alternate Erase",
	"KC_SYSREQ":      "SysReq/Attention",
	"KC_CANCEL":      "Cancel",
	"KC_CLEAR":       "Clear",
	"KC_CLR":         "Clear",
	"KC_PRIOR":       "Prior",
	"KC_RETURN":      "Return",
	"KC_SEPARATOR":   "Separator",
	"KC_OUT":         "Out",
	"KC_OPER":        "Oper",
	"KC_CLEAR_AGAIN": "Clear/Again",
	"KC_CRSEL":       "CrSel/Props",
	"KC_EXSEL":       "ExSel",

	// }}}2

	// }}}1
}
