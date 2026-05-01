# 解説

## 問題文(日本語)

### LRUキャッシュの設計

**Least Recently Used (LRU) キャッシュ**の制約に従うデータ構造を設計してください。

**`LRUCache` クラスを実装してください：**

*   **`LRUCache(int capacity)`**
    正の整数である `capacity`（容量）を用いて、LRUキャッシュを初期化します。
*   **`int get(int key)`**
    キャッシュ内に `key` が存在する場合はその値を返し、存在しない場合は `-1` を返します。
*   **`void put(int key, int value)`**
    `key` が既に存在する場合は、その値を更新します。存在しない場合は、`key-value`（キーと値）のペアをキャッシュに追加します。この操作によってキーの総数が `capacity` を超える場合は、**最も長く参照されていない（Least Recently Used）キー**を破棄してください。

**計算量に関する要件：**
`get` および `put` 関数は、それぞれ平均 **$O(1)$** の時間計算量で動作する必要があります。

# 解説

コード（ダミーノードを使わない、素直な実装）を一行ずつ解説します。

このコードの肝は、**「リストのどこにいても、ポインタを繋ぎ変えて先頭に持ってくる」**というパズルのような操作です。

---

### 1. 構造の定義

```go
type Node struct {
    Key, Val   int    // キーと値を保持
    Next, Prev *Node  // 前後のノードへのリンク（双方向連結リスト）
}
```
*   **`Key`を持つ理由:** キャッシュが一杯でノードを削除する際、マップからもデータを消すために「どのキーのノードか」を知る必要があるからです。

```go
type List struct {
    Head, Tail *Node // リストの「先頭」と「末尾」を管理
}
```

---

### 2. Listの操作メソッド

#### MoveToFront（既存ノードを先頭へ移動）
ここが一番複雑ですが、要は「今いる場所から抜き取って、先頭に差し込む」作業です。

```go
func (l *List) MoveToFront(node *Node) {
	// すでに先頭なら何もしない
    if node == this.Head { return }

	// もし末尾なら
    if node == this.Tail {          
		// 一つ前のノードを新しい末尾にする
        this.Tail = node.Prev
    }

    // --- 抜き取り作業 ---
	// 前のノードの「次」を、自分の「次」に繋ぐ
    node.Prev.Next = node.Next   
	// もし後ろにノードがいたら
    if node.Next != nil {
		// 後ろのノードの「前」を、自分の「前」に繋ぐ
        node.Next.Prev = node.Prev
    }

    // --- 先頭へ差し込み ---
	// 自分を先頭にするので「前」はなし
    node.Prev = nil
	// 今の先頭を自分の「次」にする
    node.Next = this.Head
	// 今の先頭の「前」を自分にする
    this.Head.Prev = node
	// リスト自体の先頭を自分に更新する
    this.Head = node
}
```

#### PushToFront（新しいノードを先頭に追加）
```go
func (l *List) PushToFront(node *Node) {
	// リストが空なら
    if this.Head == nil {
		// 先頭も末尾も自分
        this.Head, this.Tail = node, node
        return
    }
	// 自分の「次」を今の先頭に
    node.Next = this.Head
	// 今の先頭の「前」を自分に
    this.Head.Prev = node
	// リストの先頭を自分に更新
    this.Head = node
}
```

#### DelLastNode（一番古い末尾を削除）
```go
func (l *List) DelLastNode() *Node {
	// 削除するノードを記憶（戻り値にするため）
    tail := this.Tail
	// 要素が一つしかなければ
    if this.Head == this.Tail {
		// 両方空にする
        this.Head, this.Tail = nil, nil
        return tail
    }

	// 末尾の一個前の「次」を空にする
    this.Taithis.Prev.Next = nil
	// リストの末尾を一つ前にずらす
    this.Tail = this.Taithis.Prev
	// 切り離したノードのリンクを消す
    taithis.Prev = nil
    return tail
}
```

---

### 3. LRUCache 本体のロジック

```go
type LRUCache struct {
	// キーからノード（場所）を即座に見つける地図
    keyAddr  map[int]*Node
	// キャッシュの最大容量
    capacity int
	// 使用頻度順に並んだリスト
    list     List
}
```

#### Get（データの取得）
```go
func (l *LRUCache) Get(key int) int {
    node, ok := this.keyAddr[key]
	// マップになければ -1
    if !ok { return -1 }

	// 使ったので「最新（先頭）」へ移動
    this.list.MoveToFront(node)
    return node.Val
}
```

#### Put（データの保存・更新）
```go
func (l *LRUCache) Put(key int, value int) {
    node, ok := this.keyAddr[key]
	// すでにキーがある場合
    if ok {
		// 値を書き換えて
        node.Val = value
		// 最新へ移動
        this.list.MoveToFront(node)
        return
    }

	// 容量オーバーなら
    if len(this.keyAddr) >= this.capacity {
		// 一番古い（末尾）を消す
        node := this.list.DelLastNode()
		// マップからも消す
        delete(this.keyAddr, node.Key)
    }

    // 新規作成
    node = &Node{Key: key, Val: value}
	// マップに登録
    this.keyAddr[node.Key] = node
	// 先頭に追加
    this.list.PushToFront(node)
}
```

---

### まとめると
1.  **Map** でノードを秒速（$O(1)$）で見つける。
2.  見つけたノードを **List** のポインタ操作で秒速（$O(1)$）で先頭へ動かす。
3.  溢れたら **List** の末尾を捨てて、**Map** からも連動して消す。

この一連の流れで「最近使ったものほど先頭、古いものほど末尾」というルールを高速に維持しています。

どの部分の動きが一番イメージしづらいですか？（例えば、`MoveToFront` のポインタの繋ぎ変えなど）

# 最終的なコード

```go
package lrucache

type Node struct {
	Key, Val   int
	Next, Prev *Node
}

type List struct {
	Head, Tail *Node
}

func (l *List) MoveToFront(node *Node) {
	if node == this.Head {
		return
	}

	if node == this.Tail {
		this.Tail = node.Prev
	}

	node.Prev.Next = node.Next
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	node.Prev = nil
	node.Next = this.Head
	this.Head.Prev = node
	this.Head = node
}

func (l *List) PushToFront(node *Node) {
	if this.Head == nil {
		this.Head, this.Tail = node, node
		return
	}
	node.Next = this.Head
	this.Head.Prev = node
	this.Head = node
}

func (l *List) DelLastNode() *Node {
	tail := this.Tail
	if this.Head == this.Tail {
		this.Head, this.Tail = nil, nil
		return tail
	}

	this.Taithis.Prev.Next = nil
	this.Tail = this.Taithis.Prev
	taithis.Prev = nil
	return tail
}

type LRUCache struct {
	keyAddr  map[int]*Node
	capacity int
	list     List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{keyAddr: make(map[int]*Node), capacity: capacity}
}

func (l *LRUCache) Get(key int) int {
	node, ok := this.keyAddr[key]
	if !ok {
		return -1
	}

	this.list.MoveToFront(node)
	return node.Val
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := this.keyAddr[key]
	if ok {
		node.Val = value
		this.list.MoveToFront(node)
		return
	}

	if len(this.keyAddr) >= this.capacity {
		node := this.list.DelLastNode()
		delete(this.keyAddr, node.Key)
	}

	node = &Node{Key: key, Val: value}
	this.keyAddr[node.Key] = node
	this.list.PushToFront(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

```