package trees

/*
Invert Binary Tree
You are given the root of a binary tree root. Invert the binary tree and return its root.

Example 1:

Input: root = [1,2,3,4,5,6,7]
Output: [1,3,2,7,6,5,4]

Example 2:

Input: root = [3,2,1]
Output: [3,1,2]

Example 3:

Input: root = []
Output: []
*/

// TreeNode define um nó de árvore binária
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree inverte uma árvore binária em-lugar.
// Se root for nil, retorna nil (caso base).
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil // nada a inverter
	}
	// 1) troca ponteiros das subárvores
	root.Left, root.Right = root.Right, root.Left

	// 2) inverte recursivamente as subárvores
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

/*
Tipo de algoritmo: Recursão (Depth-First Search pré-ordem)

Complexidades:

Tempo: O(n), pois visitamos cada nó exatamente uma vez.

Espaço: O(h) para a pilha de recursão, onde h é a altura da árvore (no pior caso O(n), no melhor caso O(log n) se a árvore for balanceada).

Por que usar recursão?

É direta e legível: basta inverter os ponteiros esquerdo/direito de cada nó e recursivamente processar suas subárvores.

Usa espaço extra somente para a pilha de chamadas, sem necessidade de estruturas auxiliares.

Segue naturalmente o paradigma “dividir para conquistar” aplicado a árvores.


// Explicação passo-a-passo (recursão):

Caso base: se o nó for nil, nada a fazer.

Troca de ponteiros: em tempo O(1) trocamos root.Left com root.Right.

Chamada recursiva: invertemos as subárvores resultantes.

Retorno: devolvemos o próprio root, já com todos os seus descendentes invertidos.

Esse processo garante que cada nó da árvore é visitado uma vez, atingindo a complexidade ótima de O(n).
*/
