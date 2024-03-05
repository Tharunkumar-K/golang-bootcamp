package main

import (
	"fmt"
)

type Node struct {
	Value     string
	LeftNode  *Node
	RightNode *Node
}

func PreorderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Value, " ")
	PreorderTraversal(root.LeftNode)
	PreorderTraversal(root.RightNode)
}

func PostorderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostorderTraversal(root.LeftNode)
	PostorderTraversal(root.RightNode)
	fmt.Print(root.Value, " ")
}

func constructTree(infix string) *Node {
	postfix := infixToPostfix(infix)
	stack := []*Node{}

	for _, token := range postfix {
		node := &Node{Value: token}

		if isOperator(token) {
			node.RightNode = stack[len(stack)-1]
			node.LeftNode = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
		}

		stack = append(stack, node)
	}

	return stack[0]
}

func infixToPostfix(infix string) []string {
	var postfix []string
	var stack []string

	precedence := map[string]int{
		"+": 1,
		"-": 1,
	}

	for _, token := range infix {
		t := string(token)
		if isOperand(t) {
			postfix = append(postfix, t)
		} else if t == "(" {
			stack = append(stack, t)
		} else if t == ")" {
			for {
				op := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if op == "(" {
					break
				}
				postfix = append(postfix, op)
			}
		} else if isOperator(t) {
			for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[t] {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, t)
		}
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return postfix
}

func isOperator(token string) bool {
	return token == "+" || token == "-"
}

func isOperand(token string) bool {
	return token >= "a" && token <= "z"
}

func main() {
	var infix string

	fmt.Println("Enter the expression:")
	_, err := fmt.Scan(&infix)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	root := constructTree(infix)

	fmt.Println("Infix expression:", infix)

	fmt.Println("Preorder Traversal:")
	PreorderTraversal(root)
	fmt.Println()

	fmt.Println("Postorder Traversal:")
	PostorderTraversal(root)
	fmt.Println()
}
