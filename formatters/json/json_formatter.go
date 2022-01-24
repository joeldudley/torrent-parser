package json

import (
	"strings"
	"torrent-parser/parser"
)

const piecesKey string = "\"pieces\""
const piecesPlaceholder string = "\"<SHA1 hashes>\""

// FormatAsJson writes the parse tree rooted at node as a JSON string.
func FormatAsJson(node *parser.Node) string {
	return format(node, 0)
}

// format writes the parse tree rooted at node as a JSON string, using offset for correct indentation.
func format(node *parser.Node, offset int) string {
	var stringBuilder = strings.Builder{}

	switch node.NodeType {
	case parser.DictNode:
		stringBuilder.WriteString("{\n")
		for idx, child := range node.Children {
			stringBuilder.WriteString(format(child, offset+1))
			if idx != len(node.Children)-1 {
				stringBuilder.WriteString(", ")
			}
			stringBuilder.WriteString("\n")
		}
		stringBuilder.WriteString(getIndent(offset))
		stringBuilder.WriteString("}")

	case parser.DictEntryNode:
		stringBuilder.WriteString(getIndent(offset))
		var key = format(node.Children[0], offset+1)
		stringBuilder.WriteString(key)
		stringBuilder.WriteString(": ")

		// We use a placeholder value for the SHA1 hashes of the pieces, which are not readable.
		if key == piecesKey {
			stringBuilder.WriteString(piecesPlaceholder)
		} else {
			stringBuilder.WriteString(format(node.Children[1], offset))
		}

	case parser.ListNode:
		stringBuilder.WriteString("[\n")
		for idx, child := range node.Children {
			stringBuilder.WriteString(getIndent(offset + 1))
			stringBuilder.WriteString(format(child, offset+1))
			if idx != len(node.Children)-1 {
				stringBuilder.WriteString(", ")
			}
			stringBuilder.WriteString("\n")
		}
		stringBuilder.WriteString(getIndent(offset))
		stringBuilder.WriteString("]")

	case parser.StringNode:
		stringBuilder.WriteString("\"")
		stringBuilder.WriteString(string(node.Value))
		stringBuilder.WriteString("\"")

	case parser.IntNode:
		stringBuilder.WriteString(string(node.Value))

	}

	return stringBuilder.String()
}

// getIndent returns a blank string of offset * 4 blank characters.
func getIndent(offset int) string {
	return strings.Repeat(" ", offset*4)
}
