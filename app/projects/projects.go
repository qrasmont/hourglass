package projects

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type TimerMsg struct{}

func TimerCmd() tea.Msg {
	return TimerMsg{}
}

type Project struct {
	name        string
	description string
}

func (p Project) FilterValue() string {
	return p.name
}

func (p Project) Title() string {
	return p.name
}

func (p Project) Description() string {
	return p.description
}

type Model struct {
	list list.Model
}

func New() tea.Model {

	m := Model{list: list.NewModel([]list.Item{}, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Projects"
	m.list.SetItems([]list.Item{
		Project{name: "P1", description: ""},
		Project{name: "P2", description: ""},
		Project{name: "P3", description: ""},
	})
	m.list.SetShowStatusBar(false)
	m.list.SetFilteringEnabled(false)
	m.list.SetShowHelp(false)

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-1)

	case tea.KeyMsg:
		switch msg.String() {

		case "enter":
			return m, TimerCmd
		}
	}
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return "\n" + m.list.View()
}
