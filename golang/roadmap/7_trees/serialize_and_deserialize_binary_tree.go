package trees

import (
	"fmt"
	"strings"
)

/*
Serialize and Deserialize Binary Tree
Implement an algorithm to serialize and deserialize a binary tree.

Serialization is the process of converting an in-memory structure into a sequence of bits so that it can be stored or sent across a network to be reconstructed later in another computer environment.

You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure. There is no additional restriction on how your serialization/deserialization algorithm should work.

Note: The input/output format in the examples is the same as how NeetCode serializes a binary tree. You do not necessarily need to follow this format.

Example 1:
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]

Example 2:
Input: root = []
Output: []
*/

/*
Tipo de algoritmo:

Serialização: DFS pré-ordem com marcadores de nós nulos.

Desserialização: DFS recursivo consumindo tokens da string serializada.

Complexidades:

Tempo: O(n) tanto para serializar quanto para desserializar, pois cada nó é visitado exatamente uma vez.

Espaço: O(n) para a string/resultados (serialização) e para a pilha de recursão (desserialização).

Por que usar esta abordagem?

A pré-ordem permite codificar estrutura e valores em sequência linear, incluindo # para indicar subárvores vazias.

Na desserialização, basta “ler” token a token e recriar exatamente a mesma árvore, consumindo cada token uma única vez.

É simples, robusto e não requer estruturas auxiliares além de slices e chamadas recursivas.
*/

// serialize converte a árvore em uma string usando pré-ordem.
// Nós nulos são marcados com "#". Tokens separados por vírgula.
func serialize(root *TreeNode) string {
	var sb []string
	// dfs percorre a árvore em pré-ordem
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb = append(sb, "#") // marcador de nil
			return
		}
		// adiciona valor do nó
		sb = append(sb, fmt.Sprint(node.Val))
		dfs(node.Left)  // serializa subárvore esquerda
		dfs(node.Right) // serializa subárvore direita
	}
	dfs(root)
	return strings.Join(sb, ",")
}

// deserialize reconstrói a árvore a partir da string serializada.
func deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	tokens := strings.Split(data, ",")
	idx := 0

	// build consome tokens em pré-ordem, recriando nós
	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(tokens) {
			return nil
		}
		tok := tokens[idx]
		idx++
		if tok == "#" {
			// token "#" corresponde a nil
			return nil
		}
		// cria nó com valor convertido
		val := 0
		fmt.Sscan(tok, &val)
		node := &TreeNode{Val: val}
		node.Left = build()  // reconstrói esquerda
		node.Right = build() // reconstrói direita
		return node
	}

	return build()
}

/*
Como funciona:

serialize: DFS pré-ordem anotando node.Val ou # para nil.

deserialize: lê tokens em ordem, cri-ando nós quando o token não é #, e recursa para filhos.

printLevelOrder: verifica que a árvore reconstruída corresponde à original.
*/
