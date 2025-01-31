package bubbletree

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Open   key.Binding
	Close  key.Binding
	Quit   key.Binding
	Action key.Binding
	Top    key.Binding
	Bottom key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↓", "down"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↑", "up"),
		),
		Open: key.NewBinding(
			key.WithKeys("+", "right"),
			key.WithHelp("+", "Open"),
		),
		Close: key.NewBinding(
			key.WithKeys("-", "left"),
			key.WithHelp("-", "Close"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Action: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "action"),
		),
		Top: key.NewBinding(
			key.WithKeys("ctrl+up", "pageUp"),
			key.WithHelp("pageup", "top"),
		),
		Bottom: key.NewBinding(
			key.WithKeys("ctrl+down", "pageDown"),
			key.WithHelp("pagedown", "bottom"),
		),
	}

}
