package projects

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var WindowSize tea.WindowSizeMsg

type GoToTimerMsg struct {
	ProjectName string
}

func GoToTimerCmd(name string) tea.Cmd {
	return func() tea.Msg {
		return GoToTimerMsg{ProjectName: name}
	}
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
	items []Project
	list  list.Model
	input textinput.Model
}

func New() tea.Model {

	input := textinput.New()
	input.Prompt = "> "
	input.Placeholder = "project name"
	input.CharLimit = 100
	input.Width = 50

	m := Model{
		list:  list.NewModel([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		input: input,
	}

	m.list.Title = "Projects"
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
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-1)
		WindowSize = msg

	case tea.KeyMsg:
		if m.input.Focused() {
			switch msg.String() {

			case "enter":
				m.items = append(m.items, Project{name: m.input.Value(), description: ""})
				cmd = m.list.SetItems(toListItems(m.items))
				cmds = append(cmds, cmd)

				m.input.SetValue("")
				m.input.Blur()
				m.list.SetSize(WindowSize.Width, WindowSize.Height-1)

			case "esc":
				m.input.SetValue("")
				m.input.Blur()
				m.list.SetSize(WindowSize.Width, WindowSize.Height-1)
			}

			m.input, cmd = m.input.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			switch msg.String() {

			case "enter":
				projectName := m.items[m.list.Index()].name
				cmd = GoToTimerCmd(projectName)

			case "a":
				m.input.Focus()
				m.list.SetSize(WindowSize.Width, WindowSize.Height-3)
				cmd = textinput.Blink
			default:
				m.list, cmd = m.list.Update(msg)
			}

			cmds = append(cmds, cmd)
		}

	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.input.Focused() {
		return "\n" + m.list.View() + "\n" + m.input.View()
	}

	return "\n" + m.list.View()
}

func (m Model) GetSelectedName() string {
	return m.items[m.list.Index()].name
}

func toListItems(projects []Project) []list.Item {
	items := make([]list.Item, len(projects))
	for i, project := range projects {
		items[i] = list.Item(project)
	}
	return items
}
