package trees

/*
Same Binary Tree
Given the roots of two binary trees p and q, return true if the trees are equivalent, otherwise return false.

Two binary trees are considered equivalent if they share the exact same structure and the nodes have the same values.

Example 1:
Input: p = [1,2,3], q = [1,2,3]
Output: true

Example 2:
Input: p = [4,7], q = [4,null,7]
Output: false

Example 3:
Input: p = [1,2,3], q = [1,3,2]
Output: false
*/

/*
Busca em profundidade (DFS) dupla — uma para percorrer cada nó de root e outra para comparar subárvores.

Complexidades:

Tempo: O(N·M) no pior caso, onde N é o número de nós em root e M em subRoot.

Espaço: O(h₁ + h₂) para a pilha de recursão (h₁ = altura de root, h₂ = altura de subRoot).

Por que essa abordagem?

Direto ao ponto: para cada nó de root, verificamos se a subárvore começa ali e é igual a subRoot.

Sem estruturas extras: só ponteiros e chamadas recursivas.

Clareza: o código reflete exatamente o conceito de “para cada nó de root, compare as duas árvores”.
*/

// isSameTree compara duas árvores n1 e n2.
// Retorna true se tiverem mesma estrutura e valores.
func isSameTree(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		// ambas vazias => iguais
		return true
	}
	if n1 == nil || n2 == nil {
		// uma vazia e outra não => diferente
		return false
	}
	if n1.Val != n2.Val {
		// valores diferentes => diferente
		return false
	}
	// recursa em ambas as subárvores
	return isSameTree(n1.Left, n2.Left) && isSameTree(n1.Right, n2.Right)
}

// isSubtree percorre cada nó de root e usa isSameTree
// para checar se subRoot aparece ali.
func isSubtree(root, subRoot *TreeNode) bool {
	if subRoot == nil {
		// árvore vazia sempre é subárvore
		return true
	}
	if root == nil {
		// root acabou e subRoot não => não achou
		return false
	}
	// se as árvores a partir de root batem, ok
	if isSameTree(root, subRoot) {
		return true
	}
	// senão, continua buscando à esquerda ou à direita
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

/*
Como funciona passo a passo
Percorrer root

isSubtree chama a si mesma recursivamente em cada nó de root (pré-ordem).

Comparar subárvores

Em cada nó, chama isSameTree(root, subRoot).

isSameTree retorna imediatamente false se um nó for nil quando o outro não, ou se valores diferirem.

Se o nó atual bate, recursa em Left e Right.

Decisão final

Se qualquer chamada de isSameTree retorna true, isSubtree interrompe e devolve true.

Se percorrer tudo e não encontrar, devolve false.

Esse método usa apenas recursão e ponteiros, mantendo o código simples e fiel ao conceito de subárvore.
*/
