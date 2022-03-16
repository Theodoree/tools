package _00_299



/*
208. 实现 Trie (前缀树)

实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
*/

type Trie struct {
    Head *TrieNode
}

type TrieNode struct {
    arr [26]*TrieNode
    m   map[string]struct{}
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{Head: &TrieNode{arr: [26]*TrieNode{}}}

}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
    cur := this.Head
    for k, v := range word {
        if cur.arr[v-'a'] == nil {
            n := &TrieNode{arr: [26]*TrieNode{}, m: map[string]struct{}{}}
            cur.arr[v-'a'] = n
            cur = n
        } else {
            cur = cur.arr[v-'a']
        }

        if k == len(word)-1 {
            cur.m[word] = struct{}{}
        }
    }
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {

    var i int
    cur := this.Head
    for ; i < len(word); i++ {
        if cur.arr[word[i]-'a'] != nil {
            cur = cur.arr[word[i]-'a']
        } else {
            return false
        }
    }

    _, ok := cur.m[word]

    return ok

}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    var i int
    cur := this.Head
    for ; i < len(prefix); i++ {
        if cur.arr[prefix[i]-'a'] != nil {
            cur = cur.arr[prefix[i]-'a']
        } else {
            return false
        }
    }
    return i == len(prefix)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
