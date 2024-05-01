package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/shtayeb/organizer/cmd/schedulers"
)

type model struct {
	height int
	width  int

	ready bool

	viewport viewport.Model
	form     *huh.Form
}

var (
	category string
	path     string
	schedule string
)

func NewModel() model {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("category").
				Title("What to organize?").
				Description("Choose a category to organize.").
				Value(&category).
				Options(
					huh.NewOptions("Current Directory", "Path")...,
				),
		),

		huh.NewGroup(
			huh.NewInput().Key("path").Title("Enter the full path").Value(&path),
		).WithHideFunc(func() bool { return category == "Current Directory" }),

		huh.NewGroup(
			huh.NewSelect[string]().
				Key("schedule").
				Title("Schedule the organize command ?").
				Value(&schedule).
				Options(huh.NewOptions("No", "Weekly", "Monthly")...),

			huh.NewConfirm().
				Key("done").
				Title("All done?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("finish up then")
					}
					return nil
				}).
				Affirmative("Yes").
				Negative("Wait, no"),
		),
	).
		WithWidth(45).
		WithShowHelp(true).
		WithShowErrors(false)

	m := model{
		form: form,
	}

	return m
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// widthToUse := m.width / 3 * 2

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width/2, msg.Height-5)
			m.viewport.Style = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).MaxHeight(msg.Height).MarginTop(1)
			m.viewport.KeyMap.Up = key.Binding{}
			m.viewport.KeyMap.Down = key.Binding{}

			oldScheduledCommands := schedulers.GetScheduledTasks()

			m.viewport.SetContent(oldScheduledCommands)

			m.form.WithWidth(msg.Width / 2)
			m.form.WithHeight(msg.Height - 5)

			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

func (m model) helpView() string {
	return helpStyle("\n  ↑ page up/↓ page down: Navigate • q: Quit\n")
}

func (m model) View() string {
	form := lipgloss.NewStyle().Margin(1, 0).Render(m.form.View())

	main := lipgloss.JoinHorizontal(lipgloss.Top, form, m.viewport.View()+m.helpView())

	return main
}

// func main() {

// 	f, err := tea.LogToFile("debug.log", "debug")
// 	if err != nil {
// 		fmt.Println("fatal:", err)
// 		os.Exit(1)
// 	}
// 	defer f.Close()

// 	_, err = tea.NewProgram(
// 		NewModel(),
// 		tea.WithAltScreen(),
// 		tea.WithMouseCellMotion(),
// 	).Run()

// 	if err != nil {
// 		fmt.Println("Oh no:", err)
// 		os.Exit(1)
// 	}
// }

const content = `
# Today’s Menu

## Appetizers

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Tsukemono   | $2    | Just an appetizer               |
| Tomato Soup | $4    | Made with San Marzano tomatoes  |
| Okonomiyaki | $4    | Takes a few minutes to make     |
| Curry       | $3    | We can add squash if you’d like |

## Seasonal Dishes

| Name                 | Price | Notes              |
| ---                  | ---   | ---                |
| Steamed bitter melon | $2    | Not so bitter      |
| Takoyaki             | $3    | Fun to eat         |
| Winter squash        | $3    | Today it's pumpkin |

## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Dorayaki     | $4    | Looks good on rabbits |
| Banana Split | $5    | A classic             |
| Cream Puff   | $3    | Pretty creamy!        |

All our dishes are made in-house by Karen, our chef. Most of our ingredients
are from our garden or the fish market down the street.

Some famous people that have eaten here lately:

* [x] René Redzepi
* [x] David Chang
* [ ] Jiro Ono (maybe some day)

Bon appétit!
`
