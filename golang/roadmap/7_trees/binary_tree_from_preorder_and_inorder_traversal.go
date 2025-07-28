package trees

/*
Tipo de algoritmo:
Divide and Conquer (reconstrução recursiva usando DFS e dois ponteiros em preorder e índices em inorder).

Complexidades:

Tempo: O(n), onde n = número de nós (cada nó é criado uma vez, buscas em inorder são O(1) com mapa).

Espaço: O(n) para o mapa de índices e até O(n) de profundidade de pilha de recursão no pior caso (árvore degenerada).

Por que usar este algoritmo?

O primeiro elemento de preorder é sempre a raiz da (sub)árvore.

Em inorder, tudo que está à esquerda desse valor pertence à subárvore esquerda, e tudo à direita, à subárvore direita.

Mantendo um ponteiro (preIndex) em preorder e um mapa de valor→posição em inorder, podemos em O(1) localizar a separação de subárvores e recursivamente construir cada lado.
*/

// buildTree reconstrói a árvore a partir de preorder e inorder.
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil // caso de árvore vazia
	}

	// 1) Cria mapa para achar índice de cada valor em inorder em O(1)
	inIndex := make(map[int]int, n)
	for i, v := range inorder {
		inIndex[v] = i
	}

	// 2) Ponteiro para percorrer preorder
	preIndex := 0

	// 3) Função recursiva que reconstrói subárvore de inorder[l..r]
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil // nenhum nó nessa faixa
		}
		// 3.1) O próximo valor em preorder é a raiz desta subárvore
		rootVal := preorder[preIndex]
		preIndex++
		// 3.2) Cria o nó raiz
		root := &TreeNode{Val: rootVal}

		// 3.3) Encontra posição de rootVal em inorder
		mid := inIndex[rootVal]

		// 3.4) Constroi recursivamente subárvore esquerda (inorder[l..mid-1])
		root.Left = build(l, mid-1)
		// 3.5) Constroi recursivamente subárvore direita (inorder[mid+1..r])
		root.Right = build(mid+1, r)

		return root
	}

	// 4) Dispara a construção para toda a faixa de inorder
	return build(0, n-1)
}

/*
Como funciona:

Mapa inIndex: armazena para cada valor de inorder a sua posição, permitindo separar fronteiras de subárvore em O(1).

preIndex: avança pelo slice preorder, garantindo que cada chamada recursiva use o próximo elemento como raiz.

build(l, r):

Se l > r, não há nós — retorna nil.

Pega rootVal = preorder[preIndex++], cria o nó.

Usa mid = inIndex[rootVal] para saber quantos e quais elementos ficam à esquerda e à direita.

Chama recursivamente build(l, mid-1) e build(mid+1, r) para montar Left e Right.

Resultado: a árvore é reconstruída em pré-ordem lógica, correspondendo exatamente às traversais originais.
*/
