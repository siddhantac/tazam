package tui

import (
	"strconv"
	"tazam/store"
	"tazam/task"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type updateTable struct{}

type Model struct {
	db         store.Store
	tasks      []task.Task
	tasksTable table.Model
}

func New(tasks []task.Task, db store.Store) *Model {
	return &Model{
		db:         db,
		tasks:      tasks,
		tasksTable: newTable(tasks),
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.tasksTable.Focus()
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m, updateTableCmd
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "enter":
			idStr := m.tasksTable.SelectedRow()[0]
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return m, nil
			}
			t, err := m.db.Read(int(id))
			if err != nil {
				// log it
				return m, nil
			}
			t.Status = t.Status.Next()
			m.db.Update(t)
			return m, updateTableCmd
		default:
			var cmd tea.Cmd
			m.tasksTable, cmd = m.tasksTable.Update(msg)
			return m, cmd
		}
	case updateTable:
		m.setRows()
		return m, nil
	}
	return m, nil
}

func (m *Model) View() string {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	m.tasksTable.SetStyles(s)
	return m.tasksTable.View()
}

func updateTableCmd() tea.Msg {
	return updateTable{}
}

func (m *Model) setRows() {
	tasks, _ := m.db.List()
	rows := make([]table.Row, len(tasks))
	for i, t := range tasks {
		rows[i] = table.Row{strconv.Itoa(t.ID), t.Status.String(), t.Name, t.Project, "", strconv.Itoa(t.Priority)}
	}
	m.tasksTable.SetRows(rows)
}

func newTable(tasks []task.Task) table.Model {
	return table.New(
		table.WithColumns([]table.Column{
			{Title: "ID", Width: 4},
			{Title: "Status", Width: 6},
			{Title: "Task", Width: 20},
			{Title: "Project", Width: 20},
			{Title: "Tags", Width: 20},
			{Title: "Priority", Width: 8},
		}),
	)
}
