package ui

import (
	"errors"
	"fmt"
	"gojoke/internal/domain/app"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mitchellh/go-wordwrap"
)

const listHeight = 20
const listWidth = 80

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	fullTextStyle     = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).
				Width(80).Height(3).Align(lipgloss.Center).AlignVertical(lipgloss.Center)
	paginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type ListItem interface {
	TextShort() string
	TextFull() string
}

type listItem struct {
	textShort string
	textFull  string
	item      ListItem
}

func (i listItem) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, li list.Item) {
	i, ok := li.(listItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.textShort)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type selectListModel struct {
	list     list.Model
	choice   ListItem
	quitting bool
}

func (m selectListModel) Init() tea.Cmd {
	return nil
}

func (m selectListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(listItem)
			if ok {
				m.choice = i.item
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m selectListModel) View() string {
	selected, ok := m.list.SelectedItem().(listItem)
	if !ok {
		return lipgloss.JoinVertical(lipgloss.Center, "\n", m.list.View())
	}

	fullText := wordwrap.WrapString(selected.textFull, 80)
	return lipgloss.JoinVertical(lipgloss.Center, "\n", m.list.View(), fullTextStyle.Render(fullText))
}

func RunSelectList(title string, items []ListItem) (ListItem, error) {
	listItems := make([]list.Item, len(items))
	for i, t := range items {
		listItems[i] = listItem{textShort: t.TextShort(), textFull: t.TextFull(), item: t}
	}

	l := list.New(listItems, itemDelegate{}, listWidth, listHeight)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetShowFilter(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	tModel, err := tea.NewProgram(selectListModel{list: l}, tea.WithAltScreen()).Run()
	if err != nil {
		return nil, app.NewInternalError(err)
	}

	model, ok := tModel.(selectListModel)
	if !ok {
		return nil, app.NewInternalError(errors.New("failed to cast to SelectList model"))
	}

	if model.quitting {
		return nil, app.NewAbortedError(errors.New("user aborted"))
	}

	return model.choice, nil
}
