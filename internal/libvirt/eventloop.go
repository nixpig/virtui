package libvirt

import (
	"log"

	"libvirt.org/go/libvirt"
)

// StartEventLoop initializes and runs the libvirt event loop in a goroutine.
// This function should be called once at the application startup.
func StartEventLoop() error {
	if err := libvirt.EventRegisterDefaultImpl(); err != nil {
		return err
	}
	go func() {
		if err := libvirt.EventRunDefaultImpl(); err != nil {
			log.Printf("libvirt event loop error: %v", err)
		}
	}()
	return nil
}
