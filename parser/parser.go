package parser

import (
	"errors"
	"fmt"
	"strings"
)

// ParseBencoded takes a slice of tokenizer.token and returns the root node of the parse tree.
func ParseBencoded(filePath string) (*Node, error) {
	var tokens = tokenize(filePath)
	node, _, err := parseTokens(tokens)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Print pretty-prints a Node.
func Print(node *Node) {
	printNode(node, 0)
}

// parseTokens takes a slice of tokenizer.token and returns a node and the number of tokens processed.
func parseTokens(tokens []token) (*Node, int, error) {
	var i = 0
	var nextToken = tokens[i]

	switch nextToken.tokenType {
	case stringToken:
		var newNode = Node{
			NodeType: StringNode,
			Value:    nextToken.value,
		}
		return &newNode, 1, nil

	case intToken:
		var newNode = Node{
			NodeType: IntNode,
			Value:    nextToken.value,
		}
		return &newNode, 1, nil

	case listStartToken:
		var children []*Node
		i++
		for {
			var listEntry = tokens[i]
			switch listEntry.tokenType {
			case endToken:
				var newNode = Node{
					NodeType: ListNode,
					Children: children,
				}
				return &newNode, i + 1, nil

			default:
				var childNode, tokensRead, err = parseTokens(tokens[i:])
				if err != nil {
					return nil, 0, err
				}
				children = append(children, childNode)
				i += tokensRead
			}
		}

	case dictStartToken:
		var children []*Node
		i++
		for {
			var dictEntry = tokens[i]
			switch dictEntry.tokenType {
			case endToken:
				var newNode = Node{
					NodeType: DictNode,
					Children: children,
				}
				return &newNode, i + 1, nil

			default:
				var keyNode, keyTokensRead, keyErr = parseTokens(tokens[i:])
				if keyErr != nil {
					return nil, 0, keyErr
				}
				if keyNode.NodeType != StringNode {
					// TODO - Improve error message. Identify location of bad node in the tokens.
					return nil, 0, errors.New("dict used a type other than string as a key")
				}
				i += keyTokensRead

				var valNode, valueTokensRead, valErr = parseTokens(tokens[i:])
				if valErr != nil {
					return nil, 0, valErr
				}
				i += valueTokensRead

				var dictEntryNode = Node{
					NodeType: DictEntryNode,
					Children: []*Node{keyNode, valNode},
				}
				children = append(children, &dictEntryNode)
			}
		}

	default:
		// TODO - Improve error message. Identify location of unrecognised node in the tokens.
		return nil, 0, errors.New("unrecognised node type in parse tree")
	}
}

// printNode pretty-prints a Node, using offset to provide indentation.
func printNode(node *Node, offset int) {
	switch node.NodeType {
	case DictNode:
		fmt.Print(strings.Repeat(" ", offset*4))
		fmt.Println("Found dict node")
		for _, child := range node.Children {
			printNode(child, offset+1)
		}
	case DictEntryNode:
		fmt.Print(strings.Repeat(" ", offset*4))
		fmt.Println("Found dict entry node")
		for _, child := range node.Children {
			printNode(child, offset+1)
		}
	case ListNode:
		fmt.Print(strings.Repeat(" ", offset*4))
		fmt.Println("Found list node")
		for _, child := range node.Children {
			printNode(child, offset+1)
		}
	case StringNode:
		fmt.Print(strings.Repeat(" ", offset*4))
		fmt.Print("Found string node: ")
		fmt.Println(node.Value)
	case IntNode:
		fmt.Print(strings.Repeat(" ", offset*4))
		fmt.Print("Found integer node: ")
		fmt.Println(node.Value)
	}
}
