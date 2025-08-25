### screens roadmap

#### manager
 - Manager needs to display a table of all virtual machines. 
 - Each virtual macine row must display the name, state, memory and cpu.
 - The current selected row must be highlighted. The row that's selected must be 'remembered' if I navigate to a different screen and come back.
 - The "k" and "up" keys must move the row selection in the table up. If the first row is already selected, then do nothing. 
 - The "j" "down" keys must move the row selection in the table down. If the last row is already selected, then do nothing.
 - When the state of a virtual machine changes, it must be reflected immediately - we already handle events to some extent and need to verify whether these already provide this functionality.
 - Actions that can be taken on a virtual machine must be bound to the following keys.
    1. Key: "t" | Action: "start" | Description: "Starts the virtual machine" | Type: "immediate" |
    2. Key: "p" | Action: "pause/resume" | Description: "Pauses or resumes the virtual machine" | Type: "immediate" |
    3. Key: "s" | Action: "shutdown" | Description: "Shutdown the virtual machine" | Type: "immediate" |
    4. Key: "r" | Action: "reboot" | Description: "Reboots the virtual machine" | Type: "immediate" |
    5. Key: "e" | Action: "reset" | Description: "Force resets the virtual machine" | Type: "immediate" |
    6. Key: "v" | Action: "save" | Description: "Saves the virtual machine" | Type: "prompt" |
    7. Key: "c" | Action: "clone" | Description: "Clones the virtual machine" | Type: "prompt" |
    8. Key: "x" | Action: "delete" | Description: "Deletes the virtual machine" | Type: "prompt" |
    9. Key: "o" | Action: "open" | Description: "Opens a screen with the details of the  virtual machine displayed" | Type: "prompt" |
 - For actions of type "immediate", when the key is pressed then the corresponding action must be taken on the domain via the libvirt service.
 - For actions of type "prompt", then we will eventually build out this functionality. For now, all we want to do is display a dialog message like "still work in progress". Pressing "esc" should close the dialog.
 - Long-term: filtering, sorting, pagination, detailed view. 

#### network

#### storage


---

### internal/app/model.go
 - [x] Register screens dynamically.
 - [x] Move `globalKeybindings` initialisation to a separate file.

### internal/app/update.go
 - [x] Extract the `availableScreenHeight` calculation into a private helper method to reduce redundancy.
 - [x] The explicit `m.currentScreen.SetDimensions()` calls when switching screens are redundant if the screen's `Update` method already handles `tea.WindowSizeMsg`, so those should be removed.
 - [x] Switch over actual keys instead of magic strings.


### internal/app/view.go
 - [x] `combinedKeyMap` could be moved to internal/app/model.go or internal/app/keymaps.go for better organization.
 - [x] While `lipgloss.NewStyle().Height(...).Render("")` for padding is acceptable, `log.Debug` lines should be removed. Actionable items are to fix the `lipgloss.Height(header)` formatting and remove `log.Debug` lines.

### internal/screens/{manager,network,storage}/model.go
 - [x] Remove the redundant `m.viewport.Width` and `m.viewport.Height` assignments in `SetDimensions()`.

### overall
 - [x] Tidy up all the keybindings stuff; refactoring out magic strings and shit
 - [ ] Unit test coverage

### internal/libvirt/*
 - [ ] This abstraction feels like a bit of a mess. Tidy it up.

### cmd/virtui/main.go
 - [ ] Configuration via viper

### cmd/virtui/main.go
 - [ ] Review graceful exit/shutdown, given interactions between bubbletea and libvirt

