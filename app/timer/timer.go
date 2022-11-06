package timer

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/qrasmont/hourglass/app/projects"
	"github.com/qrasmont/hourglass/data/record"
)

var (
	pageSytle = lipgloss.NewStyle().
			Padding(1, 2)

	titleSyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("62")).
			Padding(0, 1).
			MarginBottom(1)

	counterStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 0).
			Width(15).
			Align(lipgloss.Center)

	runningCounterStyle = counterStyle.Copy().
				BorderForeground(lipgloss.Color("48")).
				Foreground(lipgloss.Color("48"))

	recordsDb *record.GormRepository
)

type BackMsg struct{}

func BackCmd() tea.Msg {
	return BackMsg{}
}

type TickMsg struct{}

func TickCmd() tea.Msg {
	time.Sleep(1 * time.Second)
	return TickMsg{}
}

type WrapperUpMsg struct{}

func WrapUpCmd(pId uint, duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		logDuration(pId, duration)
		return WrapperUpMsg{}
	}
}

type Model struct {
	name         string
	projectId    uint
	current_span time.Duration
	description  string
	running      bool
	last_time    time.Time
}

func New(project *projects.Project, db *record.GormRepository) tea.Model {
	recordsDb = db
	m := Model{
		name:         project.Name,
		projectId:    project.Id,
		current_span: time.Duration(0),
		description:  "",
		running:      false,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "b", "esc":
			cmd = WrapUpCmd(m.projectId, m.current_span)
		case " ":
			m.running = !m.running

			if m.running {
				m.last_time = time.Now()
				cmd = TickCmd
			}
		}

	case TickMsg:
		m.current_span += time.Now().Sub(m.last_time)
		if m.running {
			m.last_time = time.Now()
			cmd = TickCmd
		}

	case WrapperUpMsg:
		cmd = BackCmd
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	title := titleSyle.Render(m.name)
	counter := fmt.Sprintf("%s", m.current_span.Round(time.Second))

	var counterStyled string
	if m.running {
		counterStyled = runningCounterStyle.Render(counter)
	} else {
		counterStyled = counterStyle.Render(counter)
	}

	return pageSytle.Render(title + "\n" + counterStyled)
}

func logDuration(pId uint, duration time.Duration) {
	// TODO handle day switch here ?
	record, err := recordsDb.GetRecordForProjectForDay(pId, time.Now())
	if err != nil {
		// No record for today
		recordsDb.CreateRecord(duration, pId, time.Now())
		return
	}

	record.Duration += duration
	err = recordsDb.UpdateRecord(record)
	if err != nil {
		panic("Could not update record")
	}
}
