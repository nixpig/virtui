package domain

import "github.com/digitalocean/go-libvirt"

const (
	UNKNOWN_STATE_PRESENTABLE   = "Unknown"
	NO_STATE_PRESENTABLE        = "None"
	RUNNING_STATE_PRESENTABLE   = "Running"
	BLOCKED_STATE_PRESENTABLE   = "Blocked"
	PAUSED_STATE_PRESENTABLE    = "Paused"
	SHUTDOWN_STATE_PRESENTABLE  = "Shutdown"
	CRASHED_STATE_PRESENTABLE   = "Crashed"
	SUSPENDED_STATE_PRESENTABLE = "Suspended"
	SHUTOFF_STATE_PRESENTABLE   = "Shutoff"
)

var domainState = map[libvirt.DomainState]string{
	libvirt.DomainNostate:  NO_STATE_PRESENTABLE,
	libvirt.DomainRunning:  RUNNING_STATE_PRESENTABLE,
	libvirt.DomainBlocked:  BLOCKED_STATE_PRESENTABLE,
	libvirt.DomainPaused:   PAUSED_STATE_PRESENTABLE,
	libvirt.DomainShutdown: SHUTDOWN_STATE_PRESENTABLE,
	libvirt.DomainCrashed:  CRASHED_STATE_PRESENTABLE,
	libvirt.DomainShutoff:  SHUTOFF_STATE_PRESENTABLE,
}

func PresentableState(state libvirt.DomainState) string {
	if s, ok := domainState[state]; !ok {
		return ""
	} else {
		return s
	}
}
