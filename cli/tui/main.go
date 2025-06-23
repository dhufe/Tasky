package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	cTaskFile = ".tasky.json"
)

func main() {
	m := NewModel(cTaskFile)

	fmt.Printf("items: %d", len(*m.tasks))
	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "Task", Width: 20},
		{Title: "State", Width: 10},
		{Title: "Created At", Width: 20},
		{Title: "Completed At", Width: 20},
	}

	var rows []table.Row

	for index, item := range *m.tasks {
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

	m.table = table.New(
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
	m.table.SetStyles(s)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
		os.Exit(1)
	}
}
