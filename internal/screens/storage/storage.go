package storage

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/exp/charmtone"
	"github.com/nixpig/virtui/internal/libvirtui"
	"github.com/nixpig/virtui/internal/messages"
	"github.com/nixpig/virtui/internal/screen"
)

var _ screen.Screen = (*storageScreenModel)(nil)

const (
	panePools = iota
	paneVolumes
	paneVolumeDetails
)

var (
	inactiveTitleColor = lipgloss.Color(
		lipgloss.Color(charmtone.Charcoal.Hex()),
	)
	activeTitleColor = lipgloss.Color("170")
)

// item is a wrapper around a storage pool or volume that implements the
// list.Item interface.
type item struct {
	libvirtui.StoragePool
	libvirtui.StorageVolume
}

func (i item) FilterValue() string {
	if i.StoragePool.StoragePool != nil {
		return i.StoragePool.Name()
	}

	if i.StorageVolume.StorageVol != nil {
		return i.StorageVolume.Name()
	}

	return ""
}

type itemDelegate struct{}

func (d itemDelegate) Height() int {
	return 1
}

func (d itemDelegate) Spacing() int {
	return 0
}

func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d itemDelegate) Render(
	w io.Writer,
	m list.Model,
	index int,
	listItem list.Item,
) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	var str string
	if i.StoragePool.StoragePool != nil {
		str = i.StoragePool.Name()
	} else if i.StorageVolume.StorageVol != nil {
		str = i.StorageVolume.Name()
	}

	fn := lipgloss.NewStyle().Padding(0, 1).Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return lipgloss.NewStyle().
				Padding(0, 1).
				Background(lipgloss.Color(charmtone.Dolly.Hex())).
				Foreground(lipgloss.Color(charmtone.Zest.Hex())).
				Render(s[0])
		}
	}

	fmt.Fprint(w, fn(str))
}

type storageScreenModel struct {
	id                     string
	title                  string
	service                libvirtui.Service
	pools                  list.Model
	volumes                list.Model
	volumeDetails          viewport.Model
	activePane             int
	width, height          int
	poolsPaneWidth         int
	volumesPaneWidth       int
	volumeDetailsPaneWidth int
	containerHeight        int
	err                    error
}

func NewStorageScreen(service libvirtui.Service) *storageScreenModel {
	poolsList := list.New([]list.Item{}, itemDelegate{}, 0, 0)
	poolsList.Title = "Pools"
	poolsList.SetShowHelp(false)
	poolsList.SetShowStatusBar(false)

	poolsList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color(charmtone.Ash.Hex())).
		Align(lipgloss.Left)
	poolsList.Styles.TitleBar = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, false).
		BorderForeground(lipgloss.Color(charmtone.Charcoal.Hex())).
		Padding(0, 1)

	volumesList := list.New([]list.Item{}, itemDelegate{}, 0, 0)
	volumesList.Title = "Volumes"
	volumesList.SetShowHelp(false)
	volumesList.SetShowStatusBar(false)

	volumesList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color(charmtone.Ash.Hex())).
		Align(lipgloss.Left)
	volumesList.Styles.TitleBar = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, false, false).
		BorderForeground(lipgloss.Color(charmtone.Charcoal.Hex())).
		Padding(0, 1)

	return &storageScreenModel{
		id:            "storage",
		title:         "Storage",
		service:       service,
		pools:         poolsList,
		volumes:       volumesList,
		volumeDetails: viewport.New(0, 0),
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

		containerStyle := lipgloss.NewStyle().Border(lipgloss.HiddenBorder())

		containerHeight := m.height - containerStyle.GetVerticalFrameSize()
		containerWidth := m.width - containerStyle.GetHorizontalFrameSize()

		log.Debug(
			"ScreenSizeMsg",
			"width", m.width,
			"height", m.height,
			"containerVerticalFrameSize", containerStyle.GetVerticalFrameSize(),
			"containerHorizontalFrameSize", containerStyle.GetHorizontalFrameSize(),
			"containerHeight", containerHeight,
			"containerWidth", containerWidth,
		)

		// calculate the total width available for the content of the panes, accounting for separators
		totalContentWidth := containerWidth - (2 * 2) // 2 separators, each 2 chars (1 width + 1 border)
		totalContainerHeight := containerHeight - 2   // for the borders

		log.Debug("ScreenSizeMsg", "totalContentWidth", totalContentWidth)

		m.poolsPaneWidth = totalContentWidth / 4
		m.volumesPaneWidth = totalContentWidth / 4
		m.volumeDetailsPaneWidth = totalContentWidth - m.poolsPaneWidth - m.volumesPaneWidth

		m.containerHeight = totalContainerHeight

		log.Debug(
			"ScreenSizeMsg",
			"poolsPaneWidth", m.poolsPaneWidth,
			"volumesPaneWidth", m.volumesPaneWidth,
			"volumeDetailsPaneWidth", m.volumeDetailsPaneWidth,
		)

		m.pools.SetSize(m.poolsPaneWidth, m.containerHeight)
		m.volumes.SetSize(m.volumesPaneWidth, m.containerHeight)
		m.volumeDetails.Width = m.volumeDetailsPaneWidth
		m.volumeDetails.Height = m.containerHeight

		log.Debug(
			"ScreenSizeMsg",
			"pools.Width", m.pools.Width(),
			"pools.Height", m.pools.Height(),
			"volumes.Width", m.volumes.Width(),
			"volumes.Height", m.volumes.Height(),
			"volumeDetails.Width", m.volumeDetails.Width,
			"volumeDetails.Height", m.volumeDetails.Height,
		)

	case messages.StoragePoolsMsg:
		items := make([]list.Item, len(msg.Pools))

		for i, pool := range msg.Pools {
			items[i] = item{StoragePool: pool}
		}

		m.pools.SetItems(items)
		if len(items) > 0 {
			cmds = append(cmds, m.loadVolumesCmd())
		}

	case messages.StorageVolumesMsg:
		log.Info("received StorageVolumesMsg", "count", len(msg.Volumes))
		items := make([]list.Item, len(msg.Volumes))

		for i, vol := range msg.Volumes {
			items[i] = item{StorageVolume: vol}
		}

		m.volumes.SetItems(items)
		log.Info("set volumes in list model", "count", len(m.volumes.Items()))

	case messages.ErrorMsg:
		m.err = msg.Err

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("tab"))):
			m.activePane = (m.activePane + 1) % 3

		case key.Matches(msg, key.NewBinding(key.WithKeys("shift+tab"))):
			m.activePane = (m.activePane - 1 + 3) % 3

		default:
			switch m.activePane {
			case panePools:
				switch {
				case key.Matches(msg, key.NewBinding(key.WithKeys("enter"))):
					m.activePane = paneVolumes

				case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
					m.pools, cmd = m.pools.Update(msg)

				case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
					m.pools, cmd = m.pools.Update(msg)

				default:
					originalIndex := m.pools.Index()
					m.pools, cmd = m.pools.Update(msg)

					if m.pools.Index() != originalIndex {
						cmds = append(cmds, m.loadVolumesCmd())
					}

					cmds = append(cmds, cmd)
					m.pools.SetSize(m.poolsPaneWidth, m.containerHeight)
				}

			case paneVolumes:
				switch {
				case key.Matches(msg, key.NewBinding(key.WithKeys("enter"))):
					m.activePane = paneVolumeDetails
					cmds = append(cmds, m.loadVolumeDetailsCmd())

				case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
					m.volumes, cmd = m.volumes.Update(msg)

				case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
					m.volumes, cmd = m.volumes.Update(msg)

				default:
					originalIndex := m.volumes.Index()
					m.volumes, cmd = m.volumes.Update(msg)

					if m.volumes.Index() != originalIndex {
						cmds = append(cmds, m.loadVolumeDetailsCmd())
					}

					cmds = append(cmds, cmd)
					m.volumes.SetSize(m.volumesPaneWidth, m.containerHeight)
				}

			case paneVolumeDetails:
				switch {
				case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
					m.volumeDetails, cmd = m.volumeDetails.Update(msg)

				case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
					m.volumeDetails, cmd = m.volumeDetails.Update(msg)

				default:
					m.volumeDetails, cmd = m.volumeDetails.Update(msg)
					m.volumeDetails.Width = m.volumeDetailsPaneWidth
					m.volumeDetails.Height = m.containerHeight

					cmds = append(cmds, cmd)
				}
			}
		}

	case messages.StorageVolumeDetailsMsg:
		m.volumeDetails.SetContent(strings.Join(msg.Details, "\n"))
	}

	return m, tea.Batch(cmds...)
}

func (m *storageScreenModel) View() string {
	// clean titles and set colours based on focus
	m.pools.Title = strings.Trim(m.pools.Title, "> ")
	m.volumes.Title = strings.Trim(m.volumes.Title, "> ")

	poolsTitleStyle := m.pools.Styles.Title
	volumesTitleStyle := m.volumes.Styles.Title

	volumeDetailsTitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(charmtone.Ash.Hex())).
		Align(lipgloss.Left)

	switch m.activePane {
	case panePools:
		poolsTitleStyle = poolsTitleStyle.Foreground(activeTitleColor)
		volumesTitleStyle = volumesTitleStyle.Foreground(inactiveTitleColor)

		volumeDetailsTitleStyle = volumeDetailsTitleStyle.Foreground(
			inactiveTitleColor,
		)

		m.pools.Title = "> " + m.pools.Title

	case paneVolumes:
		poolsTitleStyle = poolsTitleStyle.Foreground(inactiveTitleColor)
		volumesTitleStyle = volumesTitleStyle.Foreground(activeTitleColor)

		volumeDetailsTitleStyle = volumeDetailsTitleStyle.Foreground(
			inactiveTitleColor,
		)

		m.volumes.Title = "> " + m.volumes.Title

	case paneVolumeDetails:
		poolsTitleStyle = poolsTitleStyle.Foreground(inactiveTitleColor)
		volumesTitleStyle = volumesTitleStyle.Foreground(inactiveTitleColor)

		volumeDetailsTitleStyle = volumeDetailsTitleStyle.Foreground(
			activeTitleColor,
		)
	}

	m.pools.Styles.Title = poolsTitleStyle
	m.volumes.Styles.Title = volumesTitleStyle

	// render titles manually for viewport
	volumeDetailsTitle := volumeDetailsTitleStyle.Render("Volume Details")
	if m.activePane == paneVolumeDetails {
		volumeDetailsTitle = volumeDetailsTitleStyle.Render("> Volume Details")
	}

	// render the viewport with its title
	rightPaneContent := lipgloss.JoinVertical(
		lipgloss.Left,
		volumeDetailsTitle,
		lipgloss.NewStyle().
			Width(m.volumeDetailsPaneWidth).
			Height(m.containerHeight-lipgloss.Height(volumeDetailsTitle)).
			Render(m.volumeDetails.View()),
	)

	separator := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(lipgloss.Color(charmtone.Charcoal.Hex())).
		Height(m.containerHeight).Width(1)

	log.Debug(
		"View",
		"separator.Width",
		separator.GetWidth(),
		"separator.Height",
		separator.GetHeight(),
	)

	poolsView := lipgloss.NewStyle().
		Width(m.poolsPaneWidth).
		Render(m.pools.View())

	emptyMiddlePaneStyle := lipgloss.NewStyle().
		Width(m.volumesPaneWidth).
		Height(m.containerHeight).
		Align(lipgloss.Center, lipgloss.Center)

	var middlePane string
	if m.err != nil {
		middlePane = emptyMiddlePaneStyle.Render(
			fmt.Sprintf("Error: %v", m.err),
		)
	} else if len(m.volumes.Items()) == 0 {
		middlePane = emptyMiddlePaneStyle.Render("No volumes found in this pool.")
	} else {
		middlePane = lipgloss.NewStyle().Width(m.volumesPaneWidth).Render(m.volumes.View())
	}

	content := lipgloss.JoinHorizontal(
		lipgloss.Top,
		poolsView,
		separator.String(),
		middlePane,
		separator.String(),
		rightPaneContent,
	)

	log.Debug(
		"View",
		"poolsView.Width", lipgloss.Width(poolsView),
		"middlePane.Width", lipgloss.Width(middlePane),
		"rightPaneContent.Width", lipgloss.Width(rightPaneContent),
		"content.Width", lipgloss.Width(content),
	)

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(charmtone.Charcoal.Hex())).
		Render(content)
}

func (m *storageScreenModel) loadVolumesCmd() tea.Cmd {
	m.err = nil

	selectedItem := m.pools.SelectedItem()
	if selectedItem == nil {
		return nil
	}

	pool, ok := selectedItem.(item)
	if !ok {
		return nil
	}

	return func() tea.Msg {
		volumes, err := m.service.ListStorageVolumes(pool.StoragePool)
		if err != nil {
			return messages.ErrorMsg{Err: err}
		}

		return messages.StorageVolumesMsg{Volumes: volumes}
	}
}

func (m *storageScreenModel) loadVolumeDetailsCmd() tea.Cmd {
	m.err = nil

	selectedItem := m.volumes.SelectedItem()
	if selectedItem == nil {
		return nil
	}

	volume, ok := selectedItem.(item)
	if !ok {
		return nil
	}

	return func() tea.Msg {
		details, err := m.service.GetStorageVolumeInfo(volume.StorageVolume)
		if err != nil {
			return messages.ErrorMsg{Err: err}
		}

		return messages.StorageVolumeDetailsMsg{
			Details: strings.Split(details, "\n"),
		}
	}
}

func (m *storageScreenModel) Title() string {
	return m.title
}

func (m *storageScreenModel) HelpKeys() [][]key.Binding {
	return [][]key.Binding{
		{
			key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "move up")),
			key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "move down")),
		},
		{
			key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next pane")),
			key.NewBinding(
				key.WithKeys("shift+tab"),
				key.WithHelp("shift+tab", "prev pane"),
			),
			key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
		},
	}
}

func (m *storageScreenModel) ID() string {
	return m.id
}
