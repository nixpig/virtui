package tui

var icons = fonts{
	vm{
		off:     "󰔂",
		running: "󰻏",
		paused:  "󰾉",
		stopped: "󰾊",
		blocked: "󰠻",
		new:     "󱄐",
		menu:    "󰔃",
	},
}

type fonts struct {
	vm
}

type vm struct {
	// 󰔂 md_television
	off string

	// 󰻏 md_television_play
	running string

	// 󰾉 md_television_pause
	paused string

	// 󰾊 md_television_stop
	stopped string

	// 󰠻 md_television_off
	blocked string

	// 󱄐 md_television_shimmer
	new string

	// 󰔃 md_television_guide
	menu string
}
