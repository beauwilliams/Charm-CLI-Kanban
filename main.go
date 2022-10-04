package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type status int

const (
	todo status = iota
	inProgress
	done
)

type Task struct {
	status      status
	title       string
	description string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) FilterKey() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

type Model struct {
	list     list.Model
	err      error
	quitting bool
	loaded   bool
}

func New() *Model {
	return &Model{}
}

func (m *Model) initList(width, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "Todo"
	m.list.SetItems([]list.Item{
		Task{status: todo, title: "Task 1", description: "This is a task"},
		Task{status: todo, title: "Task 2", description: "This is another task"},
		// Task{todo, "Task 2", "This is another task"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.initList(msg.Width, msg.Height)
		m.loaded = true
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return "Goodbye!"
	}
	if m.loaded {
		return m.list.View()
	}
	return "Loading..."
}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Printf("Program exited with error: %v", err)
	}
}
