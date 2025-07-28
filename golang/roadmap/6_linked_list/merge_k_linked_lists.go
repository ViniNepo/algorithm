package linked_list

/*
Merge K Sorted Linked Lists
You are given an array of k linked lists lists, where each list is sorted in ascending order.

Return the sorted linked list that is the result of merging all of the individual linked lists.

Example 1:

Input: lists = [[1,2,4],[1,3,5],[3,6]]

Output: [1,1,2,3,3,4,5,6]
Example 2:

Input: lists = []

Output: []
Example 3:

Input: lists = [[]]

Output: []
*/

// Definição do nó de lista encadeada
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists mescla duas listas já ordenadas e retorna o head da nova lista.
// Complexidade: O(m + n), espaço O(1) extra.
func mergeTwoLists(a, b *ListNode) *ListNode {
	// nó sentinela para facilitar a montagem
	dummy := &ListNode{}
	tail := dummy

	// enquanto houver elementos em ambas as listas
	for a != nil && b != nil {
		if a.Val < b.Val {
			tail.Next = a // anexa nó de 'a'
			a = a.Next    // avança em 'a'
		} else {
			tail.Next = b // anexa nó de 'b'
			b = b.Next    // avança em 'b'
		}
		tail = tail.Next // avança o tail
	}

	// anexa o restante de uma lista que não acabou
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}

	return dummy.Next
}

// mergeKLists aplica Divide and Conquer para mesclar k listas.
// Se lists for vazio, retorna nil.
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil // caso de erro: nenhuma lista para mesclar
	}

	// mescla em pares enquanto resta mais de 1 lista
	for interval := 1; interval < n; interval *= 2 {
		for i := 0; i+interval < n; i += interval * 2 {
			// mescla lists[i] com lists[i+interval]
			lists[i] = mergeTwoLists(lists[i], lists[i+interval])
		}
	}
	return lists[0]
}

/*
Tipo de algoritmo: Divide and Conquer (mescla em pares)

Complexidades:

Tempo: O(N log k), onde N é o número total de nós em todas as listas e k é o número de listas.

Espaço: O(1) extra ideal (além da lista de recursão de profundidade O(log k) se usar recursão, ou podemos usar uma abordagem iterativa bottom-up para ter espaço estritamente constante).

Por que usar Divide and Conquer?

Mesclar duas listas ordenadas leva O(m + n).

Se juntarmos as k listas em pares, cada nó participa de O(log k) mesclas ao longo do processo, resultando em O(N log k) de tempo total.

Não precisamos de estruturas auxiliares grandes (como um heap de tamanho k), basta reaproveitar ponteiros.

Espacialmente mais eficiente que um heap (O(k)), consumindo apenas O(1) extra.
*/
