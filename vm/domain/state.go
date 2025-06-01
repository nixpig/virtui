package domain

import "github.com/digitalocean/go-libvirt"

type PresentableState string

const (
	UNKNOWN_STATE_PRESENTABLE   PresentableState = "Unknown"
	NO_STATE_PRESENTABLE        PresentableState = "None"
	RUNNING_STATE_PRESENTABLE   PresentableState = "Running"
	BLOCKED_STATE_PRESENTABLE   PresentableState = "Blocked"
	PAUSED_STATE_PRESENTABLE    PresentableState = "Paused"
	SHUTDOWN_STATE_PRESENTABLE  PresentableState = "Shutdown"
	CRASHED_STATE_PRESENTABLE   PresentableState = "Crashed"
	SUSPENDED_STATE_PRESENTABLE PresentableState = "Suspended"
	SHUTOFF_STATE_PRESENTABLE   PresentableState = "Shutoff"
)

var domainState = map[libvirt.DomainState]PresentableState{
	libvirt.DomainNostate:  NO_STATE_PRESENTABLE,
	libvirt.DomainRunning:  RUNNING_STATE_PRESENTABLE,
	libvirt.DomainBlocked:  BLOCKED_STATE_PRESENTABLE,
	libvirt.DomainPaused:   PAUSED_STATE_PRESENTABLE,
	libvirt.DomainShutdown: SHUTDOWN_STATE_PRESENTABLE,
	libvirt.DomainCrashed:  CRASHED_STATE_PRESENTABLE,
	libvirt.DomainShutoff:  SHUTOFF_STATE_PRESENTABLE,
}

func ToPresentableState(state libvirt.DomainState) PresentableState {
	if s, ok := domainState[state]; !ok {
		return ""
	} else {
		return s
	}
}
