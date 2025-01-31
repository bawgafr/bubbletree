package bubbletree

import (
	"fmt"
)

type CursorStyle int

const (
	Chevron CursorStyle = iota
	Highlight
)

type BubbleTreeModel struct {
	KeyMap  KeyMap
	Current []int
	Styles  Styles
	Tree    Tree
	walkled bool
}

type Tree struct {
	Title       string // this is what will be displayed in the tree
	Description string // this is what will be displayed when I suss how to get modal messages popping up and they click ? on a node
	Open        bool   // is the node open or closed in the ui
	Children    []Tree // Other nodes in the tree
	Level       int    // probably not needed
	locationId  []int  // location in the tree in the form [0, 1, 1] if it was the second child of the second child of the tree parent
	next        bool   // is this the last child of the parent

	Debug string
}

type SelectedNode *Tree

func (t Tree) String() string {
	return fmt.Sprintf("%s %v \n\t%v", t.Title, t.locationId, t.Children)
}

func (t Tree) hasChildren() bool {
	return len(t.Children) > 0
}

func (t *Tree) WalkTree(parentLocationId []int) {
	t.locationId = parentLocationId
	childCount := len(t.Children)
	for i, c := range t.Children {
		c.locationId = append(parentLocationId, i)
		c.next = i != childCount-1
		t.Children[i] = c
		if c.hasChildren() {
			c.WalkTree(c.locationId)
		}
	}
}

func NewBubbleTreeModel(t Tree) BubbleTreeModel {
	bubbleTreeModel := BubbleTreeModel{
		Tree:    t,
		Current: []int{0},
		Styles:  DefaultStyles(),
		KeyMap:  DefaultKeyMap(),
	}

	return bubbleTreeModel
}

func (t *Tree) LastOpenChild() *Tree {
	if t.hasChildren() && t.Open {

		lastChild := t.Children[len(t.Children)-1]

		if lastChild.hasChildren() && lastChild.Open {
			for i := len(t.Children) - 1; i >= 0; i -= 1 {
				if t.Children[i].Open {

					t = t.Children[i].LastOpenChild()
					return t
				}
			}
		} else {
			return &lastChild
		}
	}
	return t
}
