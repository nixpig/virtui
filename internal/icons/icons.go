package icons

var Icons = fonts{
	VM{
		Off:     "󰔂",
		Running: "󰻏",
		Paused:  "󰾉",
		Stopped: "󰾊",
		Blocked: "󰠻",
		New:     "󱄐",
		Menu:    "󰔃",
	},
}

type fonts struct {
	VM
}

type VM struct {
	// 󰔂 md_television
	Off string

	// 󰻏 md_television_play
	Running string

	// 󰾉 md_television_pause
	Paused string

	// 󰾊 md_television_stop
	Stopped string

	// 󰠻 md_television_off
	Blocked string

	// 󱄐 md_television_shimmer
	New string

	// 󰔃 md_television_guide
	Menu string
}
