package main

type Trie struct {
	root *Node
}
type Node struct{
	son [26]*Node
	end bool
}


func Constructor() Trie {
	return Trie{&Node{}}
}

func index(c rune) rune {
	return c - 'a'
}

func (this *Trie) Insert(word string)  {
	cur := (*this).root
	for _, c := range(word) {
			i := index(c)
			if cur.son[i] == nil{
					cur.son[i] = &Node{}
			} 
			cur = cur.son[i]
	}
	cur.end = true
}

func (this Trie) find(word string) int {
	cur := this.root
	for _, c := range(word){
			i := index(c)
			if cur.son[i] == nil {
					return -1
			}
			cur = cur.son[i]
	}
	if cur.end {
			return 1
	}
	return 2
}


func (this *Trie) Search(word string) bool {
	return this.find(word) == 1
	
}


func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix) > 0
}


/**
* Your Trie object will be instantiated and called as such:
* obj := Constructor();
* obj.Insert(word);
* param_2 := obj.Search(word);
* param_3 := obj.StartsWith(prefix);
*/