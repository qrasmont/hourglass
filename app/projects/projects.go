package projects

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/qrasmont/hourglass/data/project"
)

var (
	WindowSize tea.WindowSizeMsg
	projectDb  *project.GormRepository
)

type GoToTimerMsg struct {
	P *Project
}

func GoToTimerCmd(p *Project) tea.Cmd {
	return func() tea.Msg {
		return GoToTimerMsg{P: p}
	}
}

type RedrawProjectsMsg struct {
	projects []Project
}

func addProjectCmd(name string) tea.Cmd {
	return func() tea.Msg {
		_, err := projectDb.CreateProject(name)
		if err != nil {
			s := fmt.Sprintf("%v\n", err)
			panic(s)
		}

		return RedrawProjectsMsg{projects: getProjects()}
	}
}

func deleteProjectCmd(id uint) tea.Cmd {
	return func() tea.Msg {
		deleteProject(id)
		return RedrawProjectsMsg{projects: getProjects()}
	}
}

type Project struct {
	Id          uint
	Name        string
	description string
}

func (p Project) FilterValue() string {
	return p.Name
}

func (p Project) Title() string {
	return p.Name
}

func (p Project) Description() string {
	return p.description
}

type Model struct {
	items []Project
	list  list.Model
	input textinput.Model
}

func New(db *project.GormRepository) tea.Model {

	projectDb = db
	input := textinput.New()
	input.Prompt = "> "
	input.Placeholder = "project name"
	input.CharLimit = 100
	input.Width = 50

	m := Model{
		items: getProjects(),
		list:  list.NewModel([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		input: input,
	}

	m.list.SetItems(toListItems(m.items))

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
				cmd = addProjectCmd(m.input.Value())
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
				p := m.items[m.list.Index()]
				cmd = GoToTimerCmd(&p)

			case "a":
				m.input.Focus()
				m.list.SetSize(WindowSize.Width, WindowSize.Height-3)
				cmd = textinput.Blink

			case "d":
				id := m.items[m.list.Index()].Id
				cmd = deleteProjectCmd(id)

			default:
				m.list, cmd = m.list.Update(msg)
			}

			cmds = append(cmds, cmd)
		}

	case RedrawProjectsMsg:
		m.items = msg.projects
		cmd = m.list.SetItems(toListItems(m.items))
		cmds = append(cmds, cmd)
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
	return m.items[m.list.Index()].Name
}

func toListItems(projects []Project) []list.Item {
	items := make([]list.Item, len(projects))
	for i, project := range projects {
		items[i] = list.Item(project)
	}
	return items
}

func getProjects() []Project {
	dbPrj, err := projectDb.GetProjects()
	if err != nil {
		// No projects in db
		return []Project{}
	}

	prjs := make([]Project, len(dbPrj))

	for i, p := range dbPrj {
		prjs[i] = Project{Name: p.Name, Id: p.ID}
	}

	return prjs
}

func deleteProject(id uint) {
	err := projectDb.DeleteProject(id)
	if err != nil {
		panic("Could not delete the project")
	}
}
