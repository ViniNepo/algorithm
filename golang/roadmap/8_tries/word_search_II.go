package tries

/*
Word Search II
Given a 2-D grid of characters board and a list of strings words, return all words that are present in the grid.

For a word to be present it must be possible to form the word with a path in the board with horizontally or vertically neighboring cells. The same cell may not be used more than once in a word.

Example 1:
Input:
board = [
  ["a","b","c","d"],
  ["s","a","a","t"],
  ["a","c","k","e"],
  ["a","c","d","n"]
],
words = ["bat","cat","back","backend","stack"]
Output: ["cat","back","backend"]

Example 2:
Input:
board = [
  ["x","o"],
  ["x","o"]
],
words = ["xoxo"]
Output: []
*/

/*
Tipo de algoritmo:
Construção de uma Trie com todas as palavras + DFS (backtracking) no tabuleiro, usando a Trie para podar caminhos que não formam nenhuma palavra válida.

Complexidades:

Tempo:

Construção da Trie: O(\∑mᵢ) onde mᵢ é o comprimento de cada palavra.

DFS no tabuleiro: em média O(M·N·L), onde M×N é o tamanho do tabuleiro e L é o comprimento máximo de palavra, graças à poda pela Trie. No pior caso teórico sem poda, seria O(M·N·4·3^(L−1)).

Espaço:

O(\∑mᵢ) para a Trie, mais O(L) para a pilha de recursão do DFS.

Por que usar Trie + DFS?

Trie agrupa as palavras por prefixos, permitindo abortar imediatamente todo o backtracking assim que o prefixo atual não existir em nenhuma palavra.

DFS explora todas as rotas possíveis no tabuleiro (células adjacentes), marcando visitados para não reutilizar a mesma célula na mesma palavra.

Essa combinação é muito mais eficiente que tentar buscar cada palavra individualmente no tabuleiro, pois compartilha a exploração de prefixos comuns.
*/

// TrieNode2 representa um nó da Trie, com até 26 filhos ('a'–'z') e, se
// word != "", significa que ali termina uma palavra válida.
type TrieNode2 struct {
	children [26]*TrieNode2
	word     string
}

// buildTrie constrói a Trie a partir da lista de words.
func buildTrie(words []string) *TrieNode2 {
	root := &TrieNode2{}
	for _, w := range words {
		node := root
		for _, ch := range w {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode2{}
			}
			node = node.children[idx]
		}
		node.word = w // marca fim de palavra
	}
	return root
}

// findWords retorna todas as palavras de words que aparecem em board.
func findWords(board [][]byte, words []string) []string {
	// 1) Constroi a Trie
	root := buildTrie(words)
	var res []string

	// 2) Para cada célula, dispara um DFS se o caractere for início de algum prefixo
	rows := len(board)
	cols := len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(board, i, j, root, &res)
		}
	}
	return res
}

// dfs explora recursivamente a partir de (i,j), seguindo a Trie.
// Quando encontra node.word != "", registra a palavra em res
// e zera node.word para evitar duplicatas.
func dfs(board [][]byte, i, j int, node *TrieNode2, res *[]string) {
	// checa limites e visitação
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}
	ch := board[i][j]
	if ch == '#' { // já visitado
		return
	}
	idx := ch - 'a'
	next := node.children[idx]
	if next == nil {
		return // nenhum prefixo começa com esse caminho
	}

	// se chegamos no fim de uma palavra, registramos
	if next.word != "" {
		*res = append(*res, next.word)
		next.word = "" // evita registrar de novo
	}

	// marca como visitado
	board[i][j] = '#'
	// explora as quatro direções
	dfs(board, i+1, j, next, res)
	dfs(board, i-1, j, next, res)
	dfs(board, i, j+1, next, res)
	dfs(board, i, j-1, next, res)
	// restaura o caractere
	board[i][j] = ch
}

/*
Explicação resumida:

Trie armazena todos os words, com cada TrieNode.word não vazio indicando o fim de uma palavra.

Para cada célula (i,j) do board, chamamos dfs:

Se o caractere não corresponde a nenhum filho de node, abortamos a busca naquele caminho.

Caso contrário, avançamos na Trie. Se next.word != "", registramos a palavra e limpamos next.word para não duplicar.

Marcamos a célula como visitada ('#'), exploramos quatro direções e, ao final, restauramos o caractere.

A poda natural da Trie evita ramificações irrelevantes, tornando o algoritmo eficiente.
*/
