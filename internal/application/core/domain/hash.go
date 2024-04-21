/*
The MIT License (MIT)

Copyright (c) 2017-2020 Damian Gryski <damian@gryski.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package domain

type Rendezvous struct {
	// é um mapa que armazena a associação entre o identificador do nó (string) e seu índice (int) no slice.
	nodes map[string]int

	// é um slice que mantém os identificadores dos nós como strings.
	nstr []string

	// é um slice que armazena os valores hash dos nós.
	nhash []uint64

	// é uma função hasher que aceita uma string e retorna um hash uint64.
	hash Hasher
}

type Hasher func(s string) uint64

// NewRendezvous é um construtor para criar uma nova instância da estrutura Rendezvous.
// Ele inicializa os mapas e slices e calcula os hashes dos nós usando a função hasher fornecida.
//
// See https://medium.com/@dgryski/consistent-hashing-algorithmic-tradeoffs-ef6b8e2fcae8
// for consistent hashing algorithmic tradeoffs.
func NewRendezvous(nodes []string, hash Hasher) *Rendezvous {
	r := &Rendezvous{
		nodes: make(map[string]int, len(nodes)),
		nstr:  make([]string, len(nodes)),
		nhash: make([]uint64, len(nodes)),
		hash:  hash,
	}

	for i, n := range nodes {
		r.nodes[n] = i
		r.nstr[i] = n
		r.nhash[i] = hash(n)
	}

	return r
}

// Lookup é um método que, dado uma chave k, encontra e retorna o nó (como uma string) que deve armazenar a chave.
// Ele faz isso calculando o hash da chave e o combinando com cada hash de nó para encontrar o maior valor de hash
// resultante (usando a função xorshiftMult64 para calcular os pesos).
func (r *Rendezvous) Lookup(k string) string {
	// short-circuit if we're empty
	if len(r.nodes) == 0 {
		return ""
	}

	khash := r.hash(k)

	var midx int
	var mhash = xorshiftMult64(khash ^ r.nhash[0])

	for i, nhash := range r.nhash[1:] {
		if h := xorshiftMult64(khash ^ nhash); h > mhash {
			midx = i + 1
			mhash = h
		}
	}

	return r.nstr[midx]
}

// Add é um método que adiciona um novo nó à estrutura. Ele atualiza o mapa de nós,
// o slice de strings e o slice de hashes de nós.
func (r *Rendezvous) Add(node string) {
	r.nodes[node] = len(r.nstr)
	r.nstr = append(r.nstr, node)
	r.nhash = append(r.nhash, r.hash(node))
}

// Remove é um método que remove um nó da estrutura. Ele encontra o nó a ser removido,
// atualiza os slices e o mapa de nós para refletir a remoção e mantém a consistência dos índices.
func (r *Rendezvous) Remove(node string) {
	// find index of node to remove
	nidx := r.nodes[node]

	// remove from the slices
	l := len(r.nstr) - 1
	r.nstr[nidx] = r.nstr[l]
	r.nstr = r.nstr[:l]

	r.nhash[nidx] = r.nhash[l]
	r.nhash = r.nhash[:l]

	// update the map
	delete(r.nodes, node)
	moved := r.nstr[nidx]
	r.nodes[moved] = nidx
}

// xorshiftMult64 é uma função que realiza operações de bit-shifting e multiplicação em um valor uint64.
// Esta é uma implementação de uma função de hash de multiplicação xorshift,
// que é usada para calcular os pesos no método Lookup.
func xorshiftMult64(x uint64) uint64 {
	x ^= x >> 12 // a
	x ^= x << 25 // b
	x ^= x >> 27 // c
	return x * 2685821657736338717
}
