package trees

import "math"

/*
Valid Binary Search Tree
Given the root of a binary tree, return true if it is a valid binary search tree, otherwise return false.

A valid binary search tree satisfies the following constraints:

The left subtree of every node contains only nodes with keys less than the node's key.
The right subtree of every node contains only nodes with keys greater than the node's key.
Both the left and right subtrees are also binary search trees.

Example 1:
Input: root = [2,1,3]
Output: true

Example 2:
Input: root = [1,2,3]
Output: false

Explanation: The root node's value is 1 but its left child's value is 2 which is greater than 1.
*/

// isValidBST inicia a validação usando limites infinitos.
func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

// validate retorna false assim que encontra um nó cujo valor não está
// no intervalo (lower, upper). Caso contrário, recursa para filhos.
func validate(node *TreeNode, lower, upper int64) bool {
	if node == nil {
		return true // árvore vazia ou fim de ramo sempre válido
	}
	v := int64(node.Val)
	// checa se valor do nó viola os limites
	if v <= lower || v >= upper {
		return false
	}
	// recursa: subárvore esquerda ganha novo upper = nó atual
	// e subárvore direita ganha novo lower = nó atual
	return validate(node.Left, lower, v) &&
		validate(node.Right, v, upper)
}

/*
Como foi feito:

Função pública isValidBST chama validate com limites iniciais infinitos (MinInt64, MaxInt64).

validate(node, lower, upper):

Se node == nil, retorna true.

Converte node.Val para int64 e checa lower < Val < upper.

Se falhar, retorna false imediatamente.

Senão, recursa para node.Left com (lower, Val) e para node.Right com (Val, upper).

Ao fim, se todas as chamadas retornarem true, a árvore atende às propriedades de BST.
*/
