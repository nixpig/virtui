package storage

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nixpig/virtui/internal/app"
	"github.com/nixpig/virtui/internal/common"
	"github.com/nixpig/virtui/internal/libvirtui"
	"github.com/nixpig/virtui/internal/messages"
)

var _ app.Screen = (*storageScreenModel)(nil)

type storageScreenModel struct {
	id       string
	title    string
	viewport viewport.Model
	width    int
	height   int
	keys     common.ScrollKeyMap
	storage  map[libvirtui.StoragePool][]libvirtui.StorageVolume
}

// NewStorageScreen returns an initialised storage screen model.
func NewStorageScreen() *storageScreenModel {
	return &storageScreenModel{
		id:       "storage",
		title:    "Storage",
		viewport: viewport.New(0, 0),
		keys:     common.DefaultScrollKeyMap(),
		storage:  make(map[libvirtui.StoragePool][]libvirtui.StorageVolume),
	}
}

func (m *storageScreenModel) Init() tea.Cmd {
	return nil
}

func (m *storageScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case messages.ScreenSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case messages.StoragePoolsMsg:
		m.storage = msg.Storage

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.viewport.ScrollUp(1)
		case key.Matches(msg, m.keys.Down):
			m.viewport.ScrollDown(1)
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *storageScreenModel) View() string {
	var sb strings.Builder
	if len(m.storage) == 0 {
		sb.WriteString("No storage pools found.")
	} else {
		for k, v := range m.storage {
			sb.WriteString("Name: " + k.Name() + "\n")
			sb.WriteString("UUID: " + k.UUID() + "\n")
			sb.WriteString("Type: " + k.Type() + "\n")
			capacityValue, capacityUnit := k.Capacity()
			availableValue, availableUnit := k.Available()
			sb.WriteString("Size: " + fmt.Sprintf("%d%s (%d%s available)", capacityValue, capacityUnit, availableValue, availableUnit) + "\n")
			sb.WriteString("Location: " + k.TargetPath() + "\n")
			sb.WriteString("Volumes:\n")
			if len(v) == 0 {
				sb.WriteString("  No volumes found.\n")
			} else {
				for i, x := range v {
					volumeCapacityValue, volumeCapacityUnit := x.Capacity()
					sb.WriteString(fmt.Sprintf(
						"  %d %s - %d %s - %s \n",
						i,
						x.Name(),
						volumeCapacityValue,
						volumeCapacityUnit,
						x.TargetFormatType(),
					))
				}
			}
			sb.WriteString("\n")
		}
	}

	renderedContent := sb.String()

	m.viewport.SetContent(renderedContent)

	style := lipgloss.NewStyle().Border(lipgloss.NormalBorder())

	m.viewport.Width = max(m.width-style.GetHorizontalFrameSize(), 0)
	m.viewport.Height = max(m.height-style.GetVerticalFrameSize(), 0)

	return style.Width(m.viewport.Width).
		Height(m.viewport.Height).
		Render(m.viewport.View())
}

// Title returns the title of the screen.
func (m *storageScreenModel) Title() string {
	return m.title
}

// Keybindings returns screen-specific keybindings.
func (m *storageScreenModel) Keybindings() []key.Binding {
	return []key.Binding{
		key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "action s")),
	}
}

// ScrollKeys returns the the screen-specific keybindings for scrolling.
func (m *storageScreenModel) ScrollKeys() common.ScrollKeyMap {
	return m.keys
}

// ID returns the screen ID used to reference the screen in other areas of app.
func (m *storageScreenModel) ID() string {
	return m.id
}
