package timer

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type Model struct {
	name         string
	current_span time.Duration
	description  string
	running      bool
	last_time    time.Time
}

func New(name string, description string) tea.Model {
	m := Model{
		name:         name,
		current_span: time.Duration(0),
		description:  description,
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
			return m, BackCmd
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
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	title := titleSyle.Render(m.name)
	counter := counterStyle.Render(fmt.Sprintf("%s", m.current_span.Round(time.Second)))
	return pageSytle.Render(title + "\n" + counter)
}
