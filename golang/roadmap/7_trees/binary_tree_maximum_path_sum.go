package trees

import "math"

/*
Binary Tree Maximum Path Sum
Given the root of a non-empty binary tree, return the maximum path sum of any non-empty path.

A path in a binary tree is a sequence of nodes where each pair of adjacent nodes has an edge connecting them. A node can not appear in the sequence more than once. The path does not necessarily need to include the root.

The path sum of a path is the sum of the node's values in the path.

Example 1:

Input: root = [1,2,3]

Output: 6
Explanation: The path is 2 -> 1 -> 3 with a sum of 2 + 1 + 3 = 6.

Example 2:

Input: root = [-15,10,20,null,null,15,5,-5]

Output: 40
Explanation: The path is 15 -> 20 -> 5 with a sum of 15 + 20 + 5 = 40.
*/

/*
Tipo de algoritmo:
Divide-and-Conquer via DFS (“post-order”) com acumulação de soma máxima e poda de ganhos negativos.

Complexidades:

Tempo: O(n), pois visitamos cada nó exatamente uma vez.

Espaço: O(h) para a pilha de recursão, onde h é a altura da árvore (no pior caso O(n), no melhor O(log n) se balanceada).

Por que usar este algoritmo?

Visita única a cada nó, computando em cada retorno o “maior ganho” que podemos levar para cima (incluindo-se ou não o ramo, se positivo).

Mantém uma variável global com a soma máxima de qualquer caminho que passe pelo nó atual somando esquerda + nó + direita.

Poda sub árvores que dariam soma negativa, tratando-as como ganho zero.

Usa apenas O(1) extra (além da recursão), e roda em tempo linear.
*/

// maxPathSum retorna a soma máxima de qualquer caminho em root.
// Se root for nil, retorna 0.
func maxPathSum(root *TreeNode) int {
	// variável global que guarda a melhor soma de um caminho qualquer
	maxSum := math.MinInt64

	// dfs retorna o maior ganho que posso levar deste nó para o pai.
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// obtém ganho máximo das subárvores; se for negativo, zera (poda)
		leftGain := max(dfs(node.Left), 0)
		rightGain := max(dfs(node.Right), 0)

		// soma de um caminho que passa por node unindo esquerda + node + direita
		priceNewPath := node.Val + leftGain + rightGain

		// atualiza o máximo global, se for melhor
		if priceNewPath > maxSum {
			maxSum = priceNewPath
		}

		// devolve o ganho máximo que podemos usar no pai (somente um lado + node)
		return node.Val + max(leftGain, rightGain)
	}

	// dispara a recursão
	dfs(root)
	return maxSum
}

// max é utilitário para inteiros
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Explicação passo a passo
Chamada inicial: maxPathSum(root) define maxSum = −∞ e dispara dfs(root).

DFS recursivo (post-order):

Se node == nil, retorna ganho 0.

Calcula leftGain = max(dfs(node.Left), 0) e rightGain = max(dfs(node.Right), 0).

priceNewPath = node.Val + leftGain + rightGain é a soma de um caminho que atravessa node.

Atualiza maxSum se priceNewPath for maior.

Retorna ao pai node.Val + max(leftGain, rightGain) (escolhendo apenas um lado).

Resultado final: após visitar todos os nós, maxSum contém a máxima soma de qualquer caminho na árvore.
*/
