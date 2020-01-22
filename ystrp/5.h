#include"keymap_jp.h"
#include QMK_KEYBOARD_H

const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
	[0] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, KC_Q, KC_W, KC_D, KC_F, KC_G, KC_Y, KC_S, KC_T, KC_R, KC_P, KC_BSPC, KC_LCTL, KC_A, KC_O, KC_E, KC_U, KC_I, KC_H, KC_J, KC_K, KC_L, KC_SCLN, KC_ENT, TO(1), KC_Z, KC_X, KC_C, KC_V, KC_B, TO(2), TO(2), KC_N, KC_M, KC_SLSH, KC_RO, KC_DOT, TO(1), KC_LALT, KC_NO, KC_NO, KC_LGUI, KC_LANG2, KC_SPC, TO(3), TO(4), KC_SPC, KC_HANJ, KC_NO, DF(0), DF(5), KC_NO),
	[1] = LAYOUT(LSFT(KC_ESC), KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, LSFT(KC_DEL), LSFT(KC_TAB), LSFT(KC_Q), LSFT(KC_W), LSFT(KC_D), LSFT(KC_F), LSFT(KC_G), LSFT(KC_Y), LSFT(KC_S), LSFT(KC_T), LSFT(KC_R), LSFT(KC_P), LSFT(KC_BSPC), LSFT(KC_LCTL), LSFT_T(KC_A), LSFT(KC_O), LSFT(KC_E), LSFT(KC_U), LSFT(KC_I), LSFT(KC_H), LSFT(KC_J), LSFT(KC_K), LSFT(KC_L), KC_NO, LSFT(KC_ENT), KC_TRNS, LSFT(KC_Z), LSFT(KC_X), LSFT(KC_C), LSFT(KC_V), LSFT(KC_B), KC_NO, KC_NO, LSFT(KC_N), LSFT(KC_M), KC_NO, KC_NO, KC_NO, KC_TRNS, LSFT(KC_LALT), KC_NO, KC_NO, LSFT(KC_LGUI), KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO),
	[2] = LAYOUT(KC_F1, KC_F2, KC_F3, KC_F4, KC_F5, KC_F6, KC_F7, KC_F8, KC_F9, KC_F10, KC_F11, KC_F12, RESET, KC_NO, KC_NO, KC_NO, KC_NO, RGB_TOG, KC_NO, KC_NO, KC_BTN1, KC_BTN2, KC_WH_U, KC_NO, EEP_RST, RGB_MOD, RGB_SPI, RGB_VAI, RGB_HUI, RGB_SAI, KC_MS_L, KC_MS_D, KC_MS_U, KC_MS_R, KC_WH_D, KC_NO, AG_NORM, RGB_RMOD, RGB_SPD, RGB_VAD, RGB_HUD, RGB_SAD, KC_TRNS, KC_TRNS, KC_LEFT, KC_DOWN, KC_UP, KC_RGHT, KC_NO, KC_NO, AG_SWAP, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_TRNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO),
	[3] = LAYOUT(KC_F1, KC_F2, KC_F3, KC_F4, KC_F5, KC_F6, KC_F7, KC_F8, KC_F9, KC_F10, KC_F11, KC_F12, RESET, KC_NO, KC_NO, KC_NO, KC_NO, RGB_TOG, KC_NO, KC_NO, KC_HOME, KC_END, KC_NO, KC_NO, EEP_RST, RGB_MOD, RGB_SPI, RGB_VAI, RGB_HUI, RGB_SAI, KC_LEFT, KC_DOWN, KC_UP, KC_RGHT, KC_NO, KC_NO, AG_NORM, RGB_RMOD, RGB_SPD, RGB_VAD, RGB_HUD, RGB_SAD, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, AG_SWAP, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_TRNS, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO, KC_NO),
	[4] = LAYOUT(KC_GRV, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, JP_EXLM, JP_QUES, JP_DQT, JP_QUOT, JP_HASH, JP_AMPR, JP_PIPE, JP_CIRC, JP_COLN, KC_NO, KC_BSPC, KC_LCTL, JP_MINS, JP_PLUS, JP_LPRN, JP_RPRN, JP_ASTR, JP_TILD, JP_LCBR, JP_RCBR, JP_COLN, KC_NO, KC_ENT, KC_NO, JP_EQL, JP_AT, JP_LBRC, JP_RBRC, JP_PERC, LCTL(KC_INS), KC_NO, JP_DLR, JP_LT, JP_GT, JP_UNDS, KC_COMM, KC_NO, KC_LALT, KC_NO, KC_NO, KC_LGUI, EISU, KC_SPC, LSFT(KC_INS), KC_TRNS, KC_SPC, KANA, KC_LGUI, KC_NO, KC_NO, KC_NO),
	[5] = LAYOUT(KC_ESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_DEL, KC_TAB, KC_Q, KC_W, KC_E, KC_R, KC_T, KC_Y, KC_U, KC_I, KC_O, KC_P, KC_BSPC, KC_LCTL, KC_A, KC_S, KC_D, KC_F, KC_G, KC_H, KC_J, KC_K, KC_L, KC_SCLN, KC_ENT, KC_LSFT, KC_Z, KC_X, KC_C, KC_V, KC_B, TO(2), TO(2), KC_N, KC_M, KC_SLSH, KC_RO, KC_DOT, KC_RSFT, KC_LALT, KC_NO, KC_NO, KC_LGUI, EISU, KC_SPC, TO(3), TO(4), KC_SPC, KANA, KC_NO, DF(0), KC_TRNS, KC_NO)
};