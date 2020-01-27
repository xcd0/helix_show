# helix_show

helixキーボード向けにつくったキーマップのヘッダーファイルをいい感じに見やすく表示します。

また[QMK configurator](https://config.qmk.fm/#/helix/LAYOUT)で生成した`.json`ファイルを渡すと、
キーマップのヘッダーファイルとして出力します。  
QMK configuratorについては https://bit-trade-one.co.jp/selfmadekb/softwaremanual/ が詳しいです。

## 使い方

[release](https://github.com/xcd0/helix_show/releases)にバイナリがあるので落として使ってください。  

引数をひとつ取ります。  
入力として受け取るファイルの拡張子で2つ(.h,.json)に処理が分岐します。

### ヘッダーファイルの場合

keymap.cからキーマップの定義の部分だけ別ファイルとしてインクルードするように  
keymap.cを書き換えます。私の場合、`5.h`としました。  
インクルードするようにしたヘッダーファイルにキーマップを書きます。  

	#if HELIX_ROWS == 5
	#include "5.h"
	#else

keymap.cから取り出したキーマップを別ファイルに保存します。  
これを入力とします。私の場合、`5.h`としました。  

	const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
		[0] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, KC_Q, KC_W, KC_E, KC_R, KC_T, KC_Y, KC_U, KC_I, KC_O, KC_P, KC_BSPC, KC_LCTL, KC_A, KC_S, KC_D, KC_F, KC_G, KC_H, KC_J, KC_K, KC_L, KC_EQL, KC_ENT, TO(8), KC_Z, KC_X, KC_C, KC_V, KC_B, TG(3), TG(3), KC_N, KC_M, KC_COMM, KC_DOT, KC_SLSH, TO(8), KC_LALT, RESET, TO(10), KC_LGUI, KC_LANG2, KC_SPC, TG(4), TG(5), KC_SPC, KC_LANG1, KC_RGUI, TG(2), DF(1), DF(0)),
		[1] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, KC_Q, KC_W, KC_D, KC_F, KC_G, KC_Y, KC_S, KC_T, KC_R, KC_P, KC_BSPC, KC_LCTL, KC_A, KC_O, KC_E, KC_U, KC_I, KC_H, KC_J, KC_K, KC_L, KC_EQL, KC_ENT, TO(9), KC_Z, KC_X, KC_C, KC_V, KC_B, TG(3), TG(3), KC_N, KC_M, KC_COMM, KC_DOT, KC_SLSH, TO(9), KC_LALT, RESET, TO(10), KC_LGUI, KC_LANG2, KC_SPC, TG(4), TG(5), KC_SPC, KC_LANG1, KC_RGUI, TG(2), DF(1), DF(0)),
		[2] = LAYOUT(KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NLCK, KC_PSLS, KC_PAST, KC_PMNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_P7, KC_P8, KC_P9, KC_PPLS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_P4, KC_P5, KC_P6, KC_PENT, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_P1, KC_P2, KC_P3, KC_PENT, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_P0, KC_P0, KC_PDOT, KC_TRNS, DF(1), DF(0)),
		[3] = LAYOUT(KC_F1, KC_F2, KC_F3, KC_F4, KC_F5, KC_F6, KC_F7, KC_F8, KC_F9, KC_F10, KC_F11, KC_F12, KC_NO, RGB_M_P, RGB_M_X, RGB_M_G, RGB_SAD, RGB_SAI, KC_NO, KC_BTN1, KC_BTN3, KC_BTN2, KC_WH_U, KC_NO, KC_NO, RGB_M_SN, RGB_M_B, RGB_M_R, RGB_HUD, RGB_HUI, KC_MS_L, KC_MS_D, KC_MS_U, KC_MS_R, KC_WH_D, KC_NO, ANY(AG_NORM), RGB_M_SW, RGB_RMOD, RGB_MOD, RGB_VAD, RGB_VAI, KC_TRNS, KC_TRNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, ANY(AG_SWAP), KC_NO, KC_NO, KC_NO, RGB_SPD, RGB_SPI, RGB_TOG, RESET, EEP_RST, KC_NO, KC_NO, KC_NO, DF(1), DF(0)),
		[4] = LAYOUT(KC_NO, KC_P1, KC_P2, KC_P3, KC_P4, KC_P5, KC_P6, KC_P7, KC_P8, KC_P9, KC_P0, KC_DEL, KC_NO, KC_NO, KC_UP, KC_NO, KC_HOME, KC_PGUP, KC_LPRN, KC_RPRN, KC_LT, KC_GT, KC_AMPR, KC_BSPC, KC_NO, KC_LEFT, KC_DOWN, KC_RGHT, KC_END, KC_PGDN, KC_LCBR, KC_RCBR, KC_LBRC, KC_RBRC, KC_PIPE, KC_ENT, KC_NO, KC_NO, KC_NO, KC_NO, KC_PSCR, KC_INS, KC_TRNS, KC_TRNS, KC_AT, KC_CIRC, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, EEP_RST, RESET, KC_TRNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, DF(1), DF(0)),
		[5] = LAYOUT(KC_NO, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_NO, KC_TILD, KC_UNDS, KC_DQUO, KC_QUOT, KC_GRV, KC_LPRN, KC_RPRN, KC_GT, KC_LT, KC_AMPR, KC_BSPC, KC_LCTL, KC_PLUS, KC_MINS, KC_ASTR, KC_EQL, KC_COLN, KC_LCBR, KC_RCBR, KC_LBRC, KC_RBRC, KC_PIPE, KC_ENT, KC_LSFT, KC_EXLM, KC_QUES, KC_HASH, KC_DLR, KC_PERC, KC_NO, KC_NO, KC_AT, KC_CIRC, KC_NO, KC_NO, KC_NO, KC_LSFT, KC_LALT, KC_NO, KC_NO, KC_LGUI, KC_LANG2, KC_SPC, KC_TRNS, KC_TRNS, KC_SPC, KC_LANG1, KC_RGUI, KC_NO, DF(1), DF(0)),
		[6] = LAYOUT(KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, DF(1), DF(0)),
		[7] = LAYOUT(KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, DF(1), DF(0)),
		[8] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_TRNS, KC_TRNS, KC_0, KC_DEL, KC_TAB, LSFT(KC_Q), LSFT(KC_W), LSFT(KC_E), LSFT(KC_R), LSFT(KC_T), LSFT(KC_Y), LSFT(KC_U), LSFT(KC_I), LSFT(KC_O), LSFT(KC_P), KC_BSPC, KC_LCTL, LSFT(KC_A), LSFT(KC_S), LSFT(KC_D), LSFT(KC_F), LSFT(KC_G), LSFT(KC_H), LSFT(KC_J), LSFT(KC_K), LSFT(KC_L), LSFT(KC_EQL), KC_ENT, KC_TRNS, LSFT_T(KC_Z), LSFT(KC_X), LSFT(KC_C), LSFT(KC_V), LSFT(KC_B), KC_NO, KC_NO, LSFT(KC_N), LSFT(KC_M), LSFT(KC_COMM), LSFT(KC_DOT), LSFT(KC_SLSH), KC_TRNS, KC_LALT, KC_NO, KC_NO, LSFT(KC_LGUI), KC_LANG2, LSFT(KC_SPC), KC_NO, KC_NO, LSFT(KC_SPC), KC_LANG1, LSFT(KC_RGUI), KC_NO, DF(1), DF(0)),
		[9] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_TRNS, KC_DEL, KC_TAB, LSFT(KC_Q), LSFT(KC_W), LSFT(KC_D), LSFT(KC_F), LSFT(KC_G), LSFT(KC_Y), LSFT(KC_S), LSFT(KC_T), LSFT(KC_R), LSFT(KC_P), KC_BSPC, KC_LCTL, LSFT(KC_A), LSFT(KC_O), LSFT(KC_E), LSFT(KC_U), LSFT(KC_I), LSFT(KC_H), LSFT(KC_J), LSFT(KC_K), LSFT(KC_L), LSFT(KC_EQL), KC_ENT, KC_TRNS, LSFT(KC_Z), LSFT(KC_X), LSFT(KC_C), LSFT(KC_V), LSFT(KC_B), KC_NO, KC_NO, LSFT(KC_N), LSFT(KC_M), LSFT(KC_COMM), LSFT(KC_DOT), LSFT(KC_SLSH), KC_TRNS, KC_LALT, KC_NO, KC_NO, LSFT(KC_LGUI), KC_LANG2, LSFT(KC_SPC), KC_NO, KC_NO, LSFT(KC_SPC), KC_LANG1, LSFT(KC_RGUI), KC_NO, DF(1), DF(0)),
		[10] = LAYOUT(KC_NO, DF(0), DF(1), DF(2), DF(3), DF(4), DF(5), DF(6), DF(7), DF(8), DF(9), KC_NO, KC_NO, DF(0), DF(0), DF(0), DF(0), DF(0), DF(0), KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_TRNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, DF(1), DF(0))
	};

この状態でビルドが正常に通るなら正しく設定できています。  
この5.hを本プログラムに第一引数として渡します。拡張子は`.h`に限ります。
これで標準出力に成形されたキーマップが表示されます。
注意として<font color="red">入力とするヘッダーファイルはコメントに対応していません。</font>  
入っているとエラーになります。  

	$ ./helix_show.exe yj3/5.h  
	const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
		/*
			layer 0
			+---------+---------+---------+---------+---------+---------+         +         +---------+---------+---------+---------+---------+---------+
			|   Esc   |    1    |    2    |    3    |    4    |    5    |                   |    6    |    7    |    8    |    9    |    0    | Delete  |
			|---------+---------+---------+---------+---------+---------+         +         +---------+---------+---------+---------+---------+---------|
			|   Tab   |    q    |    w    |    e    |    r    |    t    |                   |    y    |    u    |    i    |    o    |    p    |   Bk    |
			|---------+---------+---------+---------+---------+---------+         +         +---------+---------+---------+---------+---------+---------|
			|  LCtrl  |    a    |    s    |    d    |    f    |    g    |                   |    h    |    j    |    k    |    l    |  =  +   |  Enter  |
			|---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------|
			|  TO.8   |    z    |    x    |    c    |    v    |    b    |  TG.3   |  TG.3   |    n    |    m    |  ,  <   |  .  >   |  /  ?   |  TO.8   |
			|---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------|
			|  LAlt   |  Reset  |  TO.15  |  LGUI   |  Eisu   |   Spc   |  TG.4   |  TG.5   |   Spc   |  Kana   |  RGUI   |  TG.2   |  DF.1   |  DF.0   |
			+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+---------+
		*/
		[0] = LAYOUT(KC_ESC,KC_1,KC_2,KC_3,KC_4,KC_5,KC_6,KC_7,KC_8,KC_9,KC_0,KC_DEL,KC_TAB,KC_Q,KC_W,KC_E,KC_R,KC_T,KC_Y,KC_U,KC_I,KC_O,KC_P,KC_BSPC,KC_LCTL,KC_A,KC_S,KC_D,KC_F,KC_G,KC_H,KC_J,K
	C_K,KC_L,KC_EQL,KC_ENT,TO(8,KC_Z,KC_X,KC_C,KC_V,KC_B,TG(3,TG(3,KC_N,KC_M,KC_COMM,KC_DOT,KC_SLSH,TO(8,KC_LALT,RESET,TO(15,KC_LGUI,KC_LANG2,KC_SPC,TG(4,TG(5,KC_SPC,KC_LANG1,KC_RGUI,TG(2,DF(1,D
	F(0,),

		/*
			layer 1
	(以降省略)
	*/

Cのコメント形式にしているのでそのまま標準出力をヘッダーとして出力するのもいいと思います。

	./helix_show 5.h >> 5.h

### jsonの場合

拡張機能としてQMK configratorで生成した.jsonファイルも使用できます。  
生成した.jsonのファイルを第一引数に渡します。拡張子は`.json`に限ります。  

	./helix_show keymap.json

とすると、同じディレクトリに与えたファイルと同じ名前でヘッダーファイルが生成されます。  

	./helix_show keymap.json; mv keymap.h 5.h

のようにすれば一気にコメント付きで出力できます。


