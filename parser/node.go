package parser

// A Node in a Bencoding file.
//
// Value is the string contained by a StringNode, the integer contained by an IntNode, or the empty string otherwise.
// Children are the child nodes for a DictNode, DictEntryNode or ListNode
type Node struct {
	NodeType nodeType
	Value    []byte
	Children []*Node
}

// The type of the node.
type nodeType uint8

// The types of node.
// DictNode - has Children of type DictEntryNode
// DictEntryNode - has two Children, a StringNode as the key and another nodeType as the Value
// ListNode - has Children of type nodeType
// StringNode - nodeType that stores a string
// IntNode - nodeType that stores an integer
const (
	DictNode nodeType = iota
	DictEntryNode
	ListNode
	StringNode
	IntNode
)
