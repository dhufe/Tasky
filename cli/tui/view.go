package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
	baseStyle     = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
)

func modalView() string {
	return fmt.Sprintf(
		` 
 %s %s
`,
		inputStyle.Width(30).Render("Taskname:"),
		// m.inputs[exp].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

func (m model) View() string {
	if m.ShowModalView {
		return modalView()
	} else {
		return baseStyle.Render(m.table.View()) + "\n"
	}
}
