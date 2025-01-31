package bubbletree

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m BubbleTreeModel) Init() tea.Cmd {
	// walk the tree and assign the locationid to each node in the form of a slice of ints
	// top level node is an empty slice

	if !m.walkled {
		m.Tree.WalkTree([]int{})
	}
	return nil
}

func (m BubbleTreeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	currentNode := m.getSelectedNode()
	parentNode := m.getSelectedParentNode()
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch {
		case key.Matches(msg, m.KeyMap.Up):

			// if my last current number is not 0 everything is horribly complicated
			// go decrement last index to get older sibling, but need to check if it has children and is open, if it is need to set the current
			// to the last child of that older sibling.
			// If it doesn't then just set the current to the older sibling's locationId
			if m.Current[len(m.Current)-1] > 0 {

				// so, my index is not 0, so I have an older sibling.
				myIndex := m.Current[len(m.Current)-1] // which I know is >0
				olderSiblingNode := parentNode.Children[myIndex-1]
				if olderSiblingNode.hasChildren() && olderSiblingNode.Open {
					// if the older sibling has children and is open, then I need to go down to the last child of the older sibling
					lastChild := olderSiblingNode.Children[len(olderSiblingNode.Children)-1]
					m.Current = append(m.Current, 0)
					copy(m.Current, lastChild.locationId)

					return m, nil
				}
				copy(m.Current, olderSiblingNode.locationId)
				return m, nil
			}

			// if my last current number is 0, then I might need to go up a level, as long as thats not to the root
			if m.Current[len(m.Current)-1] == 0 && len(parentNode.locationId) > 0 {
				// remove current last digit and decrement new last one
				m.Current = m.Current[:len(m.Current)-1]
				return m, nil
			}

		case key.Matches(msg, m.KeyMap.Down):

			//if currently selected node has children and is open
			if currentNode.hasChildren() && currentNode.Open {
				m.Current = append(m.Current, 0)
				return m, nil
			}

			// if current is not the last child of its parent
			noSiblings := len(parentNode.Children)
			myPositionInSiblings := m.Current[len(m.Current)-1]

			// if there are more siblings than me, add one
			if myPositionInSiblings < noSiblings-1 {
				copy(m.Current, parentNode.Children[myPositionInSiblings+1].locationId)
				return m, nil
			}

			if myPositionInSiblings == noSiblings-1 && len(parentNode.locationId) != 0 && parentNode.next { // I am the last in the list and not first level
				m.Current = m.Current[:len(m.Current)-1]
				m.Current[len(m.Current)-1]++

				return m, nil
			}

		case key.Matches(msg, m.KeyMap.Open):
			if currentNode.hasChildren() && !currentNode.Open {
				currentNode.Open = true
			}
		case key.Matches(msg, m.KeyMap.Close):
			if currentNode.hasChildren() && currentNode.Open {
				currentNode.Open = false
			}
		case key.Matches(msg, m.KeyMap.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.KeyMap.Top):
			m.Current = []int{0}
		case key.Matches(msg, m.KeyMap.Bottom):
			// need to check if the bottom one is open (down the chain!), and if it is assign current to last open child

			lastOpenChild := m.Tree.LastOpenChild()
			l := len(lastOpenChild.locationId)
			m.Current = make([]int, l)
			copy(m.Current, lastOpenChild.locationId)

		case key.Matches(msg, m.KeyMap.Action):
			return m, func() tea.Msg {
				return SelectedNode(currentNode)
			}

		}

	}
	return m, nil
}

func (m BubbleTreeModel) getSelectedParentNode() *Tree {
	l := len(m.Current)

	if l == 0 {
		panic("No Data")
	}

	t := &m.Tree

	for i := 0; i < l-1; i++ {
		t = &t.Children[m.Current[i]]
	}

	return t
}

func (m BubbleTreeModel) getSelectedNode() *Tree {
	l := len(m.Current)

	if l == 0 {
		panic("No Data")
	}

	t := &m.Tree

	for i := 0; i < l; i++ {
		t = &t.Children[m.Current[i]]
	}

	return t
}
