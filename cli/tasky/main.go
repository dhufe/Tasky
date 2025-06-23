package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github/tasky"
)

const (
	taskFile = ".tasky.json"
)

// style
var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

// model
type model struct {
	table table.Model
	tasks *tasky.Todos
}

func main() {
	tasks := &tasky.Todos{}

	// Load tasks from the file.
	if err := tasks.Load(taskFile); err != nil {
		fmt.Printf("failed to load tasks: %w", err)
	} else {
		PrintTable(*tasks)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "t":
			indexTask, _ := strconv.Atoi(m.table.SelectedRow()[0])
			m.tasks.ToggleState(indexTask)

			if (*m.tasks)[indexTask-1].Done {
				m.table.SelectedRow()[2] = "✅"
				m.table.SelectedRow()[4] = (*m.tasks)[indexTask-1].CompletedAt.Format(time.RFC822)

			} else {
				m.table.SelectedRow()[2] = "❌"
				m.table.SelectedRow()[4] = "-"
			}
			m.table, cmd = m.table.Update(msg)
			m.table.UpdateViewport()
			return m, cmd

		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func PrintTable(tasks tasky.Todos) {
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Task", Width: 20},
		{Title: "State", Width: 10},
		{Title: "Created At", Width: 20},
		{Title: "Completed At", Width: 20},
	}

	var rows []table.Row

	for index, item := range tasks {
		task := item.Task
		done := "❌"
		completedAt := "-"

		if item.Done {
			task = item.Task
			done = "✅"
			completedAt = item.CompletedAt.Format(time.RFC822)
		}

		rows = append(rows, table.Row{
			fmt.Sprintf("%d", index+1),
			task,
			done,
			item.CreatedAt.Format(time.RFC822),
			completedAt,
		})
	}

	tb := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	tb.SetStyles(s)

	m := model{tb, &tasks}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
