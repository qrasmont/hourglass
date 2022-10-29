package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var p *tea.Program

type State int

const (
	projectList State = iota
)

type MainModel struct {
	state State
}

func New() MainModel {
	return MainModel{
		state: projectList,
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

	return m, cmd
}

func (m MainModel) View() string {
	return "app"
}
