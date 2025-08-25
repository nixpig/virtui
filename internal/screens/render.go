package screens

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// RenderViewport renders the given viewport model with content and applies styling.
// It takes the available width and height for the viewport area.
func RenderViewport(vp viewport.Model, content string, availableWidth, availableHeight int) string {
	vp.SetContent(content)

	style := lipgloss.NewStyle().Border(lipgloss.NormalBorder())

	// Calculate viewport dimensions considering the border
	vp.Width = max(availableWidth-style.GetHorizontalFrameSize(), 0)
	vp.Height = max(availableHeight-style.GetVerticalFrameSize(), 0)

	return style.Width(vp.Width).
		Height(vp.Height).
		Render(vp.View())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
