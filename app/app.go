package app

import (
	"fmt"
	"os"

	"github.com/qrasmont/hourglass/app/projects"

	tea "github.com/charmbracelet/bubbletea"
)

var p *tea.Program

type State int

const (
	projectState State = iota
)

type MainModel struct {
	state    State
	projects tea.Model
}

func New() MainModel {
	return MainModel{
		state:    projectState,
		projects: projects.New(),
	}
}

func Start() {
	m := New()
	p = tea.NewProgram(m)
	p.EnterAltScreen()
	if err := p.Start(); err != nil {
		fmt.Println("App error: ", err)
		os.Exit(1)
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.projects, cmd = m.projects.Update(msg)

	return m, cmd
}

func (m MainModel) View() string {
	return m.projects.View()
}
