package libvirtui

import (
	"github.com/charmbracelet/log"
	"libvirt.org/go/libvirt"
)

// StartEventLoop initializes and runs the libvirt event loop in a goroutine.
// This function should be called once at the application startup.
func StartEventLoop() error {
	if err := libvirt.EventRegisterDefaultImpl(); err != nil {
		return err
	}

	go func() {
		for {
			if err := libvirt.EventRunDefaultImpl(); err != nil {
				log.Error("libvirt event loop error", "err", err)
			}
		}
	}()

	return nil
}
