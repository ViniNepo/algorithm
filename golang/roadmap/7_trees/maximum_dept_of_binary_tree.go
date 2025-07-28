package trees

/*
Maximum Depth of Binary Tree
Given the root of a binary tree, return its depth.

The depth of a binary tree is defined as the number of nodes along the longest path from the root node down to the farthest leaf node.

Example 1:



Input: root = [1,2,3,null,null,4]

Output: 3
Example 2:

Input: root = []

Output: 0
*/

/*
Tipo de algoritmo: Recursão (Depth-First Search pós-ordem)

Complexidades:

Tempo: O(n), pois visitamos cada nó exatamente uma vez.

Espaço: O(h) para a pilha de recursão, onde h é a altura da árvore (no pior caso O(n), no melhor caso O(log n) se a árvore for balanceada).

Por que usar recursão?

É a maneira mais direta de “dividir para conquistar” em árvores: a profundidade máxima de um nó é 1 + max(profundidade à esquerda, profundidade à direita).

Usa apenas espaço extra para a pilha de chamadas, sem estruturas auxiliares.

Facilita a implementação e a leitura do código.
*/

// maxDepth retorna a profundidade máxima da árvore cujo root é o nó raiz.
// Se root for nil, a profundidade é 0.
func maxDepth(root *TreeNode) int {
	if root == nil {
		// Caso base: árvore vazia tem profundidade zero
		return 0
	}
	// Calcula profundidade das subárvores
	leftDepth := maxDepth(root.Left)   // profundidade da esquerda
	rightDepth := maxDepth(root.Right) // profundidade da direita

	// Profundidade do nó atual é 1 + maior profundidade entre as duas subárvores
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

/*
Explicação passo a passo:

Caso base: se o nó for nil, retornamos 0.

Recursão: obtemos leftDepth chamando maxDepth(root.Left) e rightDepth chamando maxDepth(root.Right).

Combinação: a profundidade máxima a partir de root é 1 + max(leftDepth, rightDepth).

Resultado: percorremos toda a árvore em O(n) e usamos apenas O(h) de espaço para a pilha de chamadas.
*/
