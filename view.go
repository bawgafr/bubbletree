package bubbletree

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Graphics            TreeGraphics
	CursorStyle         CursorStyle
	SelectedHighlight   lipgloss.Style
	UnSelectedHighlight lipgloss.Style
	Lines               lipgloss.Style
}

func DefaultStyles() Styles {
	return Styles{
		Graphics:            DefaultTreeGraphics(),
		Lines:               lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
		CursorStyle:         Chevron,
		SelectedHighlight:   lipgloss.NewStyle().Foreground(lipgloss.Color("255")),
		UnSelectedHighlight: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
	}
}

type TreeGraphics struct {
	VerticalLine   string
	JoinLine       string
	HorizontalLine string
	Chevron        string
	Open           string
	Close          string
}

func DefaultTreeGraphics() TreeGraphics {
	return TreeGraphics{
		VerticalLine:   "│",
		JoinLine:       "└",
		HorizontalLine: "─",
		Chevron:        ">",
		Open:           "+",
		Close:          "-",
	}
}

func (m BubbleTreeModel) View() string {
	s := m.Tree.View(0, m.Current, m.Styles)

	return s
}

func (t Tree) View(initial int, current []int, style Styles) string {

	indent := strings.Repeat(" ", initial)
	var s string
	if len(t.locationId) == 0 {
		s += fmt.Sprintf("%s%s\n", indent, style.UnSelectedHighlight.Render(t.Title))
	}

	l := (len(t.Title) / 3) // should this use the average length of all siblings?
	indent += strings.Repeat(" ", l)

	joinLine := style.Lines.Render(style.Graphics.JoinLine)
	horizLine := style.Lines.Render(style.Graphics.HorizontalLine)
	vertLine := style.Lines.Render(style.Graphics.VerticalLine)

	// loop through the children of the passed in node
	for _, c := range t.Children {
		// cursor or highlighting
		cursor := " "
		displayString := c.Title
		if style.CursorStyle == Chevron {
			if reflect.DeepEqual(current, c.locationId) {
				cursor = style.Graphics.Chevron
			}
		} else if style.CursorStyle == Highlight {
			if reflect.DeepEqual(current, c.locationId) {
				displayString = style.SelectedHighlight.Render(c.Title)
			} else {
				displayString = style.UnSelectedHighlight.Render(c.Title)
			}
		}

		// open Close symbol to diplay
		openClose := " "
		if c.hasChildren() {
			if c.Open {
				openClose = style.Graphics.Close
			} else {
				openClose = style.Graphics.Open
			}
		}
		// print all the things we defined above
		s += fmt.Sprintf(" %s%s\n", indent, vertLine)
		s += fmt.Sprintf("%s%s%s%s%s %s\n", cursor, indent, joinLine, horizLine, displayString, openClose)
		if c.Open {
			s += c.View(l, current, style)
		}

	}

	return s
}

func LocationString(location []int) string {
	s := "(("
	for _, l := range location {
		s += fmt.Sprintf("%d ", l)
	}
	return s + "))"
}
