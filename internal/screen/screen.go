package screen

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
)

// Screen is the interface that all application screens must implement.
type Screen interface {
	tea.Model
	ID() string
	Title() string
	HelpKeys() [][]key.Binding
}
