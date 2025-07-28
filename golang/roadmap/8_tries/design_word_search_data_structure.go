package tries

/*
Design Add and Search Word Data Structure
Design a data structure that supports adding new words and searching for existing words.

Implement the WordDictionary class:

void addWord(word) Adds word to the data structure.
bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.
Example 1:

Input:
["WordDictionary", "addWord", "day", "addWord", "bay", "addWord", "may", "search", "say", "search", "day", "search", ".ay", "search", "b.."]

Output:
[null, null, null, null, false, true, true, true]

Explanation:
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("day");
wordDictionary.addWord("bay");
wordDictionary.addWord("may");
wordDictionary.search("say"); // return false
wordDictionary.search("day"); // return true
wordDictionary.search(".ay"); // return true
wordDictionary.search("b.."); // return true
Constraints:

1 <= word.length <= 20
word in addWord consists of lowercase English letters.
word in search consist of '.' or lowercase English letters.
There will be at most 2 dots in word for search queries.
At most 10,000 calls will be made to addWord and search.
*/

/*
Tipo de estrutura: Trie com busca DFS para suportar o caractere curinga '.'.

Complexidades:

addWord(word):

Tempo: O(m), onde m = comprimento de word.

Espaço: O(m) no pior caso, para criar novos nós.

search(word):

Tempo: no pior caso O(26ᵏ · n), onde n = tamanho de word e k = número de '.'. Em cenários típicos, O(n · 26).

Espaço: O(h) para a pilha de recursão, h ≤ n.

Por que usar Trie + DFS?

A Trie permite inserções e buscas de prefixos em tempo proporcional ao tamanho da string.

Para tratar '.' como “qualquer letra”, executamos uma pequena DFS: ao encontrar '.', tentamos todos os ramos filhos daquele nó.

Mantemos marcação de fim de palavra (isEnd) para saber se a sequência completa existe.
*/

// TrieNode representa um nó da Trie, com filhos para 'a'–'z' e flag de fim de palavra.
//type TrieNode struct {
//	children [26]*TrieNode
//	isEnd    bool
//}

// WordDictionary contém a raiz da Trie.
type WordDictionary struct {
	root *TrieNode
}

// Constructor inicializa a estrutura.
func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}

// AddWord insere a palavra na Trie.
func (wd *WordDictionary) AddWord(word string) {
	node := wd.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

// Search retorna true se existe alguma palavra na Trie que case com word.
// O caractere '.' casa com qualquer letra.
func (wd *WordDictionary) Search(word string) bool {
	return searchInNode(word, wd.root)
}

// searchInNode faz DFS recursiva começando em node.
func searchInNode(word string, node *TrieNode) bool {
	for i, ch := range word {
		if ch == '.' {
			// tenta todos os 26 ramos
			for _, child := range node.children {
				if child != nil && searchInNode(word[i+1:], child) {
					return true
				}
			}
			return false
		}
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	// só é match se terminar exatamente numa palavra completa
	return node.isEnd
}

/*
Como funciona:

Inserção (AddWord): percorre cada caractere, criando nós conforme necessário, e marca isEnd=true no último.

Busca (Search):

Para letras normais, desce diretamente pelo ramo correspondente.

Para '.', itera todos os filhos não-nulos, recursivamente verificando o restante da palavra.

No final da string, verifica isEnd para garantir correspondência completa.

Essa abordagem mantém operações de inserção e busca eficientes, suportando o wildcard '.' com uma pequena sobrecarga de DFS quando necessário.
*/
