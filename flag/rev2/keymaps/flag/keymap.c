#include QMK_KEYBOARD_H
#include "bootloader.h"
#ifdef PROTOCOL_LUFA
#include "lufa.h"
#include "split_util.h"
#endif


extern uint8_t is_master;

// Each layer gets a name for readability, which is then used in the keymap matrix below.
// The underscores don't mean anything - you can have a layer called STUFF or any other name.
// Layer names don't all need to be of the same length, obviously, and you can also skip them
// entirely and just use numbers.
enum layer_number {
    _QWERTY = 0,
	_YSTRP,
	_LEFT,
	_RIGHT,
	_ADJUST,
	_SYMBOL,
};

enum custom_keycodes {
  QWERTY = SAFE_RANGE,
  YSTRP,
  LEFT,
  RIGHT,
  SYMBOL,
  ADJUST,
  EISU,
  KANA
};

enum macro_keycodes {
  KC_SAMPLEMACRO,
};

//Macros
#define M_SAMPLE M(KC_SAMPLEMACRO)

#include"keymap_jp.h"
#define JP_BTN1 KC_BTN1
#define JP_BTN2 KC_BTN2
#define JP_BTN3 KC_BTN3
#define JP_WH_U KC_WH_U
#define JP_WH_D KC_WH_D
#define JP_MS_L KC_MS_L
#define JP_MS_U KC_MS_U
#define JP_MS_D KC_MS_D
#define JP_MS_R KC_MS_R
#define JP_LEFT KC_LEFT
#define JP_DOWN KC_DOWN
#define JP_UP   KC_UP
#define JP_RGHT KC_RGHT
#define JP_BSPC KC_BSPC
#define JP_DEL  KC_EQL
#define JP_ENT  KC_ENT
#define JP_SPC  KC_SPC
#define JP_HOME KC_HOME
#define JP_END  KC_END
#define JP_INS  KC_INS
#define JP_ESC  KC_ESC
#define JP_TAB  KC_TAB
#define JP_LALT KC_LALT
#define JP_LCTL KC_LCTL
#define JP_LSFT KC_LSFT
#define JP_LGUI KC_LGUI
#define JP_LEFT KC_LEFT
#define JP_DOWN KC_DOWN
#define JP_UP   KC_UP
#define JP_RGHT KC_RGHT

#include "5.h"

// define variables for reactive RGB
bool TOG_STATUS = false;
int RGB_current_mode;

void persistent_default_layer_set(uint16_t default_layer) {
  eeconfig_update_default_layer(default_layer);
  default_layer_set(default_layer);
}

// Setting ADJUST layer RGB back to default
void update_tri_layer_RGB(uint8_t layer1, uint8_t layer2, uint8_t layer3) {
  if (IS_LAYER_ON(layer1) && IS_LAYER_ON(layer2)) {
    #ifdef RGBLIGHT_ENABLE
      //rgblight_mode(RGB_current_mode);
    #endif
    layer_on(layer3);
  } else {
    layer_off(layer3);
  }
}

bool process_record_user(uint16_t keycode, keyrecord_t *record) {
  switch (keycode) {
    case QWERTY:
      if (record->event.pressed) {
        persistent_default_layer_set(1UL<<_QWERTY);
      }
      return false;
      break;
    case YSTRP:
      if (record->event.pressed) {
        persistent_default_layer_set(1UL<<_YSTRP);
      }
      return false;
      break;
    case LEFT:
      if (record->event.pressed) {
          //not sure how to have keyboard check mode and set it to a variable, so my work around
          //uses another variable that would be set to true after the first time a reactive key is pressed.
        if (TOG_STATUS) { //TOG_STATUS checks is another reactive key currently pressed, only changes RGB mode if returns false
        } else {
          TOG_STATUS = !TOG_STATUS;
        }
        layer_on(_LEFT);
        update_tri_layer_RGB(_LEFT, _RIGHT, _ADJUST);
      } else {
        TOG_STATUS = false;
        layer_off(_LEFT);
        update_tri_layer_RGB(_LEFT, _RIGHT, _ADJUST);
      }
      return false;
      break;
    case RIGHT:
      if (record->event.pressed) {
        //not sure how to have keyboard check mode and set it to a variable, so my work around
        //uses another variable that would be set to true after the first time a reactive key is pressed.
        if (TOG_STATUS) { //TOG_STATUS checks is another reactive key currently pressed, only changes RGB mode if returns false
        } else {
          TOG_STATUS = !TOG_STATUS;
        }
        layer_on(_RIGHT);
        update_tri_layer_RGB(_LEFT, _RIGHT, _ADJUST);
      } else {
        layer_off(_RIGHT);
        TOG_STATUS = false;
        update_tri_layer_RGB(_LEFT, _RIGHT, _ADJUST);
      }
      return false;
      break;
    case ADJUST:
        if (record->event.pressed) {
          layer_on(_ADJUST);
        } else {
          layer_off(_ADJUST);
        }
        return false;
        break;
      //led operations - RGB mode change now updates the RGB_current_mode to allow the right RGB mode to be set after reactive keys are released
    case EISU:
      if (record->event.pressed) {
        if(keymap_config.swap_lalt_lgui==false){
          register_code(KC_LANG2);
        }else{
          SEND_STRING(SS_LALT("`"));
        }
      } else {
        unregister_code(KC_LANG2);
      }
      return false;
      break;
    case KANA:
      if (record->event.pressed) {
        if(keymap_config.swap_lalt_lgui==false){
          register_code(KC_LANG1);
        }else{
          SEND_STRING(SS_LALT("`"));
        }
      } else {
        unregister_code(KC_LANG1);
      }
      return false;
      break;
  }
  return true;
}

void matrix_init_user(void) {}

