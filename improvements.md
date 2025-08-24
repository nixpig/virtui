### overall
 - [ ] Tidy up all the keybindings stuff; refactoring out magic strings and shit
 - [ ] Unit test coverage

### cmd/virtui/main.go
 - [ ] Graceful exit/shutdown
 - [ ] Configuration via viper

### internal/app/model.go
 - [x] Register screens dynamically.
 - [x] Move `globalKeybindings` initialisation to a separate file.

### internal/app/update.go
 - [x] Extract the `availableScreenHeight` calculation into a private helper method to reduce redundancy.
 - [ ] The explicit `m.currentScreen.SetDimensions()` calls when switching screens are redundant if the screen's `Update` method already handles `tea.WindowSizeMsg`, so those should be removed.
 - [ ] Switch over actual keys instead of magic strings.


### internal/app/view.go
 - [ ] `combinedKeyMap` could be moved to internal/app/model.go or internal/app/keymaps.go for better organization.
 - [x] While `lipgloss.NewStyle().Height(...).Render("")` for padding is acceptable, `log.Debug` lines should be removed. Actionable items are to fix the `lipgloss.Height(header)` formatting and remove `log.Debug` lines.

### internal/screens/{manager,network,storage}/model.go
 - [ ] Remove the duplicated content in `View()`.
 - [x] Remove the redundant `m.viewport.Width` and `m.viewport.Height` assignments in `SetDimensions()`.

