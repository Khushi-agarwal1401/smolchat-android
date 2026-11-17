package main

import (
	gioapp "gioui.org/app"
	"gioui.org/io/event"

	"github.com/smollabs/smolchat-android/app"
)

// platformEvent records the Android View (used for haptic feedback).
func platformEvent(e event.Event) {
	if v, ok := e.(gioapp.AndroidViewEvent); ok {
		app.SetAndroidView(v.View)
	}
}
