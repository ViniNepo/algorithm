package trees

/*
Lowest Common Ancestor in Binary Search Tree
Given a binary search tree (BST) where all node values are unique, and two nodes from the tree p and q, return the lowest common ancestor (LCA) of the two nodes.

The lowest common ancestor between two nodes p and q is the lowest node in a tree T such that both p and q as descendants. The ancestor is allowed to be a descendant of itself.

Example 1:
Input: root = [5,3,8,1,4,7,9,null,2], p = 3, q = 8
Output: 5

Example 2:
Input: root = [5,3,8,1,4,7,9,null,2], p = 3, q = 4
Output: 3

Explanation: The LCA of nodes 3 and 4 is 3, since a node can be a descendant of itself.
*/

/*
Tipo de algoritmo:
Busca iterativa explorando a propriedade de BST.

Complexidades:

Tempo: O(h), onde h é a altura da árvore (no pior caso h≈n para árvore degenerada, mas em média h=O(log n) para BST balanceada).

Espaço: O(1) extra (somente variáveis auxiliares).

Por que usar esse algoritmo?
Em uma BST, para qualquer nó node:

Se ambos p.Val e q.Val são menores que node.Val, então o LCA está na subárvore esquerda.

Se ambos são maiores, o LCA está na subárvore direita.

Caso contrário (um em cada lado, ou um deles é node), node é o primeiro ponto em que p e q “se separam”—é o LCA.

Isso evita percorrer toda a árvore ou usar estruturas auxiliares: basta seguir um caminho único do raiz até o LCA.
*/

// lowestCommonAncestor encontra o LCA de p e q em uma BST.
// Se a árvore for vazia ou nenhum LCA existir, retorna nil.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// percorre do topo até encontrar o ponto de ramificação
	for root != nil {
		// ambos p e q à esquerda → vai para esquerda
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
			// ambos p e q à direita → vai para direita
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
			// caso contrário, root é o LCA
		} else {
			return root
		}
	}
	return nil
}

/*
Fluxo resumido:

Começamos em root.

Enquanto root != nil:

Se ambos p.Val e q.Val estiverem abaixo de root.Val, movemos para root.Left.

Se ambos estiverem acima, movemos para root.Right.

Caso contrário, encontramos o ponto de divisão: root é o LCA.

Se chegarmos a nil sem encontrar, devolvemos nil.

Essa solução é simples, roda em tempo O(h) e usa O(1) de espaço extra.
*/
