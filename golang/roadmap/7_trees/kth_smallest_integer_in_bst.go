package trees

/*
Kth Smallest Integer in BST
Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) in the tree.

A binary search tree satisfies the following constraints:

The left subtree of every node contains only nodes with keys less than the node's key.
The right subtree of every node contains only nodes with keys greater than the node's key.
Both the left and right subtrees are also binary search trees.
Example 1:



Input: root = [2,1,3], k = 1

Output: 1
Example 2:



Input: root = [4,3,5,2,null], k = 4

Output: 5
*/

/*
Tipo de algoritmo:
Travessia em-ordem (DFS in-order) recursiva, que explora primeiro a subárvore esquerda, depois o nó atual e por fim a subárvore direita. Em um BST, essa travessia produz os valores em ordem crescente, então basta contar até o k-ésimo nó visitado.

Complexidades:

Tempo: O(n) no pior caso (se k≈n), mas O(h+k) em média, onde h é a altura da árvore, pois paramos logo que alcançamos o k-ésimo.

Espaço: O(h) na pilha de recursão (h = altura da árvore).

Por que usar in-order?
Em uma BST, a propriedade “left < node < right” faz com que a travessia in-order liste os valores em ordem crescente. Assim, o k-ésimo elemento visitado durante essa travessia é exatamente o k-ésimo menor valor.
*/

// kthSmallest retorna o k-ésimo menor valor em uma BST.
// Se a árvore for menor que k nós, retorna -1 como sinal de erro.
func kthSmallest(root *TreeNode, k int) int {
	res := -1  // guardará o resultado
	count := 0 // contador de nós visitados

	// inorder faz a travessia em-ordem recursiva.
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return // para se esgotou a árvore ou já encontrou o k-ésimo
		}
		inorder(node.Left) // explora esquerda
		count++            // visita o nó atual
		if count == k {    // se for o k-ésimo
			res = node.Val // salva resultado
			return
		}
		inorder(node.Right) // explora direita
	}

	inorder(root)
	return res
}

/*
Como foi feito:

Variáveis auxiliares: count conta quantos nós já foram visitados; res guarda o valor quando count == k.

Travessia in-order: função inorder(node) recursa em node.Left, então processa o node (incrementa count e, se for o k-ésimo, armazena res), e finalmente recursa em node.Right.

Poda: quando count >= k, interrompemos novas recursões para economizar tempo.

Retorno: após chamar inorder(root), res contém o k-ésimo menor valor, ou -1 se não houver elementos suficientes.
*/
