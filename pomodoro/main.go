package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
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
	stopwatch          stopwatch.Model
	isBreak            bool
	isInited           bool
	isRunning          bool
	keymaps            Keymap
	help               help.Model
}

type Keymap struct {
	start  key.Binding
	stop   key.Binding
	reset  key.Binding
	quit   key.Binding
	work   key.Binding
	breakk key.Binding
	pause  key.Binding
}

func initialize() Model {
	m := Model{
		keymaps: Keymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "Start"),
			),

			pause: key.NewBinding(
				key.WithKeys("p"),
				key.WithHelp("p", "Pause/play"),
			),

			stop: key.NewBinding(
				key.WithKeys("x"),
				key.WithHelp("x", "Stop the stopwatch"),
			),

			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "Restart the stopwatch"),
			),

			quit: key.NewBinding(
				key.WithKeys("q", "ctr+c"),
				key.WithHelp("q", "Quit the application"),
			),

			work: key.NewBinding(
				key.WithKeys("w"),
				key.WithHelp("w", "Start work"),
			),

			breakk: key.NewBinding(
				key.WithKeys("b"),
				key.WithHelp("b", "Start break"),
			),
		},
		help:               help.New(),
		workTimeInMinutes:  25,
		breakTimeInMinutes: 5,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymaps.quit):
			return m, tea.Quit

		case key.Matches(msg, m.keymaps.reset):
			if m.isRunning {

				if m.isBreak {
					m.stopwatch = stopwatch.NewWithInterval(time.Second)
				} else {
					m.stopwatch = stopwatch.NewWithInterval(time.Second)
				}
				return m, m.stopwatch.Init()
			}

			return m, nil
		case key.Matches(msg, m.keymaps.stop):
			m.isRunning = false
			m.isInited = false
			return m, m.stopwatch.Stop()

		case key.Matches(msg, m.keymaps.pause):
			m.isRunning = !m.isRunning
			if m.isRunning {
				return m, m.stopwatch.Start()
			} else {
				return m, m.stopwatch.Stop()
			}

		case key.Matches(msg, m.keymaps.breakk):
			if m.isBreak {
				return m, nil
			}
			m.isBreak = true
			m.stopwatch = stopwatch.NewWithInterval(time.Second)
			return m, m.stopwatch.Init()

		case key.Matches(msg, m.keymaps.start, m.keymaps.work):
			if m.isRunning {
				return m, nil
			}
			m.isInited = true
			m.stopwatch = stopwatch.NewWithInterval(time.Second)
			m.isRunning = true
			return m, m.stopwatch.Init()
		}
	}

	if m.isRunning {
		var cmd tea.Cmd
		m.stopwatch, cmd = m.stopwatch.Update(msg)
		return m, cmd
	}

	if !m.isInited {
		return m, nil
	}

	timeSoFar := m.stopwatch.Elapsed()

	if !m.isBreak {
		if timeSoFar >= time.Duration(m.workTimeInMinutes)*time.Minute {
			m.isBreak = true
			m.stopwatch = stopwatch.NewWithInterval(time.Second)
			return m, m.stopwatch.Init()
		}
	} else {
		if timeSoFar >= time.Duration(m.breakTimeInMinutes)*time.Minute {
			m.isBreak = false
			m.stopwatch = stopwatch.NewWithInterval(time.Second)
			return m, m.stopwatch.Init()
		}
	}

	return m, nil
}

func (m Model) helpView() string {
	bindings := []key.Binding{
		m.keymaps.start,
		m.keymaps.quit,
	}

	// determine if the stopwatch has been initialized or not

	if m.isInited {
		bindings = append(bindings, m.keymaps.reset)
		bindings = append(bindings, m.keymaps.pause)
		bindings = append(bindings, m.keymaps.stop)

		if m.isBreak {
			bindings = append(bindings, m.keymaps.work)
		} else {
			bindings = append(bindings, m.keymaps.breakk)
		}

	}

	return "\n" + m.help.ShortHelpView(bindings)
}

func (m Model) View() string {
	s := "\n"

	if m.isInited {
		if m.isBreak {
			s += "Break... "
		} else {
			s += "Work... "
		}

		s += m.stopwatch.View()
	} else {
		s += "You currently have no session running, start a new session!"
	}

	s += "\n" + m.helpView()

	return s
}

func main() {
	m := initialize()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
