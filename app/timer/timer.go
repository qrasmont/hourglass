package timer

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type BackMsg struct{}

func BackCmd() tea.Msg {
	return BackMsg{}
}

type Model struct {
	project     *tea.Model
	span        time.Duration
	description string
}

func New(project *tea.Model, description string) tea.Model {
	m := Model{
		project:     project,
		span:        time.Duration(0),
		description: description,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "b", "esc":
			return m, BackCmd
		}
	}

	return m, cmd
}

func (m Model) View() string {
	return "Timer view"
}
