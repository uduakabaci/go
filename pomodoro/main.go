package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type Pomodoro struct {
	WorkTimeInMinutes  int
	BreakTimeInMinutes int
}

type Model struct {
	breakTimeInMinutes int
	workTimeInMinutes  int
	slices             []Pomodoro
	timer              timer.Model
	quitting           bool
	keymaps            Keymap
	help               help.Model
}

type Keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
}

func initialize() Model {
	m := Model{
		keymaps: Keymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "Start"),
			),

			stop: key.NewBinding(
				key.WithKeys("p"),
				key.WithHelp("p", "Stop the timer"),
			),

			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "Restart the timer"),
			),

			quit: key.NewBinding(
				key.WithKeys("q", "ctr+c"),
				key.WithHelp("q", "Quit the application"),
			),
		},
		help:               help.New(),
		workTimeInMinutes:  25,
		breakTimeInMinutes: 5,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.TimeoutMsg:
		m.quitting = true
		return m, nil

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymaps.quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keymaps.reset):
			if m.timer.Running() {
				m.timer.Stop()
			}
			// m.timer.Timeout = time.Second * 5

			return m, nil
		case key.Matches(msg, m.keymaps.stop):
			return m, m.timer.Stop()

		case key.Matches(msg, m.keymaps.start):
			timeout := time.Minute * time.Duration(m.workTimeInMinutes)
			m.timer = timer.NewWithInterval(timeout, time.Millisecond)
			m.timer.Init()

			return m, nil
		}
	}

	return m, nil
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymaps.start,
		m.keymaps.stop,
		m.keymaps.quit,
	})
}

func (m Model) View() string {
	s := m.timer.View()

	if m.timer.Timedout() {
		s = "\nAll is done!"
	}

	s += "\n"

	if !m.quitting {
		s = "\nExiting in " + s
		s += m.helpView()
	}

	return s
}

func main() {
	m := initialize()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
