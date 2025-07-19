package tui

var icons = fonts{
	vm{
		empty:   "󰔂",
		play:    "󰻏",
		pause:   "󰾉",
		stop:    "󰾊",
		off:     "󰠻",
		shimmer: "󱄐",
		guide:   "󰔃",
		light:   "󱍖",
	},
}

type fonts struct {
	vm
}

type vm struct {
	// 󰔂 md_television
	empty string

	// 󰻏 md_television_play
	play string

	// 󰾉 md_television_pause
	pause string

	// 󰾊 md_television_stop
	stop string

	// 󰠻 md_television_off
	off string

	// 󱄐 md_television_shimmer
	shimmer string

	// 󰔃 md_television_guide
	guide string

	// 󱍖 md_television_ambient_light
	light string
}
