package icons

type vmIcons struct {
	Running string
	Paused  string
	Blocked string
	Off     string
}

var Vm = vmIcons{
	Running: "▶",
	Paused:  "⏸",
	Blocked: "⛔",
	Off:     "■",
}
