
bool TOG_STATUS = false;
bool process_record_user(uint16_t keycode, keyrecord_t *record) {
	switch (keycode) {
	case YSTRP:
		if (record->event.pressed) {
			persistent_default_layer_set(1UL<<_YSTRP);
		}
		return false;
		break;
	case SHIFT:
		if (record->event.pressed) {
			persistent_default_layer_set(1UL<<_SHIFT);
			layer_on(_SHIFT);
		} else {
			persistent_default_layer_set(1UL<<_YSTRP);
			layer_off(_SHIFT);
		}
		return false;
		break;
	case MOUSE:
		if (record->event.pressed) {
			persistent_default_layer_set(1UL<<_MOUSE);
			layer_on(_MOUSE);
		} else {
			persistent_default_layer_set(1UL<<_YSTRP);
			layer_off(_MOUSE);
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
		} else {
			layer_off(_LEFT);
			TOG_STATUS = false;
		}
		return false;
	case RIGHT:
		if (record->event.pressed) {
			//not sure how to have keyboard check mode and set it to a variable, so my work around
			//uses another variable that would be set to true after the first time a reactive key is pressed.
			if (TOG_STATUS) { //TOG_STATUS checks is another reactive key currently pressed, only changes RGB mode if returns false
			} else {
				TOG_STATUS = !TOG_STATUS;
			}
			layer_on(_RIGHT);
		} else {
			layer_off(_RIGHT);
			TOG_STATUS = false;
		}
		return false;
		break;

	
		case QWERTY:
			if (record->event.pressed) {
				persistent_default_layer_set(1UL<<_QWERTY);
			}
			return false;
			break;
/*
		case COLEMAK:
			if (record->event.pressed) {
				persistent_default_layer_set(1UL<<_COLEMAK);
			}
			return false;
			break;
		case DVORAK:
			if (record->event.pressed) {
				persistent_default_layer_set(1UL<<_DVORAK);
			}
			return false;
			break;
	case SYMBOL:
		if (record->event.pressed) {
			persistent_default_layer_set(1UL<<_SYMBOL);
			layer_on(_SYMBOL);
		} else {
			persistent_default_layer_set(1UL<<_YSTRP);
			layer_off(_SYMBOL);
		}
		return false;
		break;
*/
		case RGB_MOD:
			#ifdef RGBLIGHT_ENABLE
				if (record->event.pressed) {
					rgblight_mode(RGB_current_mode);
					rgblight_step();
					RGB_current_mode = rgblight_config.mode;
				}
			#endif
			return false;
			break;
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
		case RGBRST:
			#ifdef RGBLIGHT_ENABLE
				if (record->event.pressed) {
					eeconfig_update_rgblight_default();
					rgblight_enable();
					RGB_current_mode = rgblight_config.mode;
				}
			#endif
			break;
	}
	return true;
}
