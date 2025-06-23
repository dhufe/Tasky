package main

import (
	"log"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github/tasky"
)

type model struct {
	//
	ShowModalView bool
	// table to list tasks
	table table.Model
	// tasks data structure
	tasks *tasky.Todos
}

func NewModel(taskFile string) model {
	tasks := &tasky.Todos{}

	if err := tasks.Load(taskFile); err != nil {
		log.Fatalf("failed to load tasks file %s. %v", taskFile, err)
	}
	return model{
		ShowModalView: false,
		table:         table.New(),
		tasks:         tasks,
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
		case "a":
			if m.ShowModalView {
				m.ShowModalView = false
			} else {
				m.ShowModalView = true
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
