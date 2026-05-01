package lrucache

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type List struct {
	Head, Tail *Node
}

// どういうパターンが必要か？
// - リストの先頭にデータを追加する処理
// - 末尾からデータを削除する処理
// - ノードを先頭に動かす
func (l *List) MoveToFront(node *Node) {
	// 素手に先頭なら何もしない
	if node == l.Head {
		return
	}

	// もし末尾なら
	if node == l.Tail {
		// ひとつ前のノードを新しい末尾にする
		l.Tail = node.Prev
	}

	// 抜き取り作業
	// 前のノードの次を、自分の次につなぐ
	node.Prev.Next = node.Next
	// もし後ろにノードがいたら
	if node.Next != nil {
		// 後ろのノードの前を自分の前んつなぐ
		node.Next.Prev = node.Prev
	}

	// 先頭へ差し込み
	node.Prev = nil
	node.Next = l.Head
	l.Head.Prev = node
	l.Head = node
}

// 新しいノードを先頭に追加
func (l *List) PushToFront(node *Node) {
	// リストが空なら
	if l.Head == nil {
		// 先頭も末尾も自分
		l.Head, l.Tail = node, node
		return
	}
	// 自分の次を今の先頭に
	node.Next = l.Head
	l.Head.Prev = node
	// リストの先頭を自分に更新
	l.Head = node
}

// 一番古い末尾を削除
func (l *List) DelLastNode() *Node {
	// 削除するノードを記録(戻り値にするため)
	tail := l.Tail
	// 要素が一つしかなければ
	if l.Head == l.Tail {
		// 両方空にする
		l.Head, l.Tail = nil, nil
		return tail
	}

	// 末尾の一個前の「次」を空にする
	l.Tail.Prev.Next = nil
	// リストの末尾を一つ前にずらす
	l.Tail = l.Tail.Prev
	// 切り離したノードのリンクを消す
	tail.Prev = nil
	return tail
}

type LRUCache struct {
	// キーからノードを即座に見つける地図
	keyAddr map[int]*Node
	// キャッシュの最大容量
	capacity int
	// 使用頻度順に並んだリスト
	list List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		keyAddr:  map[int]*Node{},
		list:     List{},
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.keyAddr[key]
	// マップになければ-1を返却する
	if !ok {
		return -1
	}

	// 使ったので「最新」としてリストの先頭へ移動
	this.list.MoveToFront(node)
	return node.Val
}

func (this *LRUCache) Put(key, value int) {
	node, ok := this.keyAddr[key]
	// すでにキーがーある場合
	if ok {
		// 値を書き換えて
		node.Val = value
		// 最新へ移動
		this.list.MoveToFront(node)
		return
	}

	// 容量オーバーなら
	if len(this.keyAddr) >= this.capacity {
		// 一番古い(末尾)を消す
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
