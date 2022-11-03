package app

import (
	"fmt"
	"os"

	"github.com/qrasmont/hourglass/app/projects"
	"github.com/qrasmont/hourglass/app/timer"
	"github.com/qrasmont/hourglass/data/project"
	"github.com/qrasmont/hourglass/data/record"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	p         *tea.Program
	projectDb *project.GormRepository
	recordDb  *record.GormRepository
)

type State int

const (
	projectState State = iota
	timerState
)

type MainModel struct {
	state    State
	projects tea.Model
	timer    tea.Model
}

func New() MainModel {
	return MainModel{
		state:    projectState,
		projects: projects.New(),
	}
}

func Start(project project.GormRepository, record record.GormRepository) {
	projectDb = &project
	recordDb = &record
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

	case projects.GoToTimerMsg:
		m.state = timerState
		m.timer = timer.New(msg.ProjectName, "")

	case timer.BackMsg:
		m.state = projectState
	}

	switch m.state {
	case projectState:
		m.projects, cmd = m.projects.Update(msg)
	case timerState:
		m.timer, cmd = m.timer.Update(msg)
	}

	return m, cmd
}

func (m MainModel) View() string {
	switch m.state {
	case projectState:
		return m.projects.View()
	case timerState:
		return m.timer.View()
	}

	return ""
}
