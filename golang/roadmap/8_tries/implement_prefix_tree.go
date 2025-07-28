package tries

/*
Implement Trie (Prefix Tree)
A prefix tree (also known as a trie) is a tree data structure used to efficiently store and retrieve keys in a set of strings. Some applications of this data structure include auto-complete and spell checker systems.

Implement the PrefixTree class:

PrefixTree() Initializes the prefix tree object.
void insert(String word) Inserts the string word into the prefix tree.
boolean search(String word) Returns true if the string word is in the prefix tree (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
Example 1:

Input:
["Trie", "insert", "dog", "search", "dog", "search", "do", "startsWith", "do", "insert", "do", "search", "do"]

Output:
[null, null, true, false, true, null, true]

Explanation:
PrefixTree prefixTree = new PrefixTree();
prefixTree.insert("dog");
prefixTree.search("dog");    // return true
prefixTree.search("do");     // return false
prefixTree.startsWith("do"); // return true
prefixTree.insert("do");
prefixTree.search("do");     // return true
*/

/*
Tipo de estrutura: Árvore de Prefixos (Trie)

Complexidades:

insert(word): O(m) de tempo, onde m = comprimento de word; O(m) de espaço no pior caso para criar novos nós.

search(word): O(m) de tempo, O(1) de espaço extra.

startsWith(prefix): O(m) de tempo, O(1) de espaço extra.

Por que usar Trie?
Uma Trie permite inserir e buscar strings (ou verificar prefixes) em tempo proporcional ao seu comprimento,
independente de quantas palavras já existam. É ideal para auto-complete, dicionários e aplicações de busca de prefixos.
*/

// TrieNode representa um nó da Trie, com até 26 filhos (a–z)
// e um flag que indica se ali termina uma palavra.
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

// PrefixTree é o wrapper que mantém a raiz da Trie.
type PrefixTree struct {
	root *TrieNode
}

// NewPrefixTree inicializa e retorna uma Trie vazia.
func NewPrefixTree() *PrefixTree {
	return &PrefixTree{root: &TrieNode{}}
}

// Insert adiciona a palavra word na Trie.
func (pt *PrefixTree) Insert(word string) {
	node := pt.root
	for _, ch := range word {
		idx := ch - 'a' // índice de 0 a 25
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{} // cria nó se não existir
		}
		node = node.children[idx] // avança para o filho
	}
	node.isEnd = true // marca fim de palavra
}

// Search retorna true se word foi inserida anteriormente.
func (pt *PrefixTree) Search(word string) bool {
	node := pt.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false // falhou ao encontrar caminho
		}
		node = node.children[idx]
	}
	return node.isEnd // só é palavra se isEnd estiver marcado
}

// StartsWith retorna true se existir alguma palavra inserida que comece com prefix.
func (pt *PrefixTree) StartsWith(prefix string) bool {
	node := pt.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return false // não há tal prefixo
		}
		node = node.children[idx]
	}
	return true // percorreu todo prefixo com sucesso
}

/*
Como foi feito:

Nó base (TrieNode) com um array fixo de 26 filhos (children[0] = 'a', …, children[25] = 'z') e um booleano isEnd para marcar o término de uma palavra.

Insert percorre cada caractere da palavra, criando nós conforme necessário, e ao final marca o nó como isEnd = true.

Search faz o mesmo percurso e, chegando ao fim da string, verifica se isEnd está marcado.

StartsWith percorre apenas o prefixo, retornando true se não faltar nenhum nó ao longo do caminho.

Essa implementação garante buscas e inserções eficientes, em tempo O(m) onde m é o tamanho da palavra ou prefixo.
*/
