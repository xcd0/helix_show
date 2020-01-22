
bool TOG_STATUS = false;
bool process_record_user(uint16_t keycode, keyrecord_t *record) {
	switch (keycode) {
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
