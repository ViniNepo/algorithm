package trees

/*
Binary Tree Level Order Traversal
Given a binary tree root, return the level order traversal of it as a nested list, where each sublist contains the values of nodes at a particular level in the tree, from left to right.

Example 1:
Input: root = [1,2,3,4,5,6,7]
Output: [[1],[2,3],[4,5,6,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []
*/

/*
Tipo de algoritmo:
Busca em largura (Breadth-First Search — BFS) usando fila para percorrer a árvore nível a nível.

Complexidades:

Tempo: O(n), pois cada nó é enfileirado e desenfileirado exatamente uma vez.

Espaço: O(n) no pior caso (último nível pode conter até n/2 nós), mais o espaço para o resultado que também é O(n).

Por que usar BFS?

Precisamos coletar os nós nível por nível, exatamente a ordem natural de uma BFS.

A fila garante que, ao processar um nó, seus filhos ficam ao final da fila, servindo aos próximos elementos do mesmo nível ou do próximo.

É simples de implementar e faz exatamente o que o problema pede, sem estruturas auxiliares complexas.
*/

// levelOrder retorna os valores da árvore em ordem de níveis,
// onde cada slice interno contém os valores de um nível.
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result // árvore vazia → nenhum nível
	}

	// fila para armazenar nós a visitar
	queue := []*TreeNode{root}

	// enquanto houver nós na fila
	for len(queue) > 0 {
		levelSize := len(queue) // número de nós no nível atual
		levelVals := make([]int, 0, levelSize)

		// processa todos os nós deste nível
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // desenfileira

			// registra o valor
			levelVals = append(levelVals, node.Val)

			// enfileira filhos para o próximo nível
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// adiciona os valores deste nível ao resultado
		result = append(result, levelVals)
	}

	return result
}

/*
Como foi feito:

Inicialização: se a árvore estiver vazia, retorna [][]int{} imediatamente.

Fila: começamos com o root na fila.

Loop de níveis: enquanto a fila não estiver vazia, determinamos levelSize (quantos nós estão neste nível).

Processamento de nível: para cada nó deste nível, desenfileiramos, gravamos seu valor em levelVals e enfileiramos seus filhos (se existirem).

Acumulação: após consumir todos os nós do nível, adicionamos levelVals ao result.

Continuação: repetimos até não restarem nós na fila.

Retorno: result contém todos os níveis, na ordem correta.
*/
