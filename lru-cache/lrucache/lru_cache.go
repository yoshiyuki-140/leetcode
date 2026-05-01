package lrucache

// このコードの肝は「リストのどこにいても、ポインタをつなぎ変えて先頭に戻ってくる」
// というパズルのような操作

// 1. 構造の定義
type Node struct {
	Key, Val   int   // キーと値を保持
	Next, Prev *Node // 前後のノードへのリンク(双方向連結リスト)
	//
}

type List struct {
	Head, Tail *Node // リストの先頭と末尾を管理
}

// Listの操作メソッド
// NOTE: 今いる場所から抜き取って、先頭に差し込む
// LRUを実現するためにはFIFO(First in first out)
func (l *List) MoveToFront(node *Node) {
	if node == l.Head {
		return
	} // すでに戦闘なら何もしない
	if node == l.Tail { // もし後ろにノードがいたら
		l.Tail = node.Prev // 後ろのノードの前を自分のノードの前に繋ぐ
	}

	// --- 先頭へ差し込み ---
	// 自分を先頭にするので「前」
	node.Prev = nil
	// 今の先頭を自分の「次」にする
	node.Next = l.Head
	// 今の先頭の「前」を自分にする
	l.Head.Prev = node
	// リスト自体の先頭を自分に更新する
	l.Head = node
}

// PushToFront(新しいノードを先頭に追加)
func (l *List) PushToFront(node *Node) {
	// リストが空なら
	if l.Head == nil {
		// 先頭も末尾も自分
		l.Head, l.Tail = node, node
		return
	}
	node.Next = l.Head
	l.Head.Prev = node
	l.Head = node
}

// 一番古い末尾を削除
func (l *List) DelLastNode() *Node {
	tail := l.Tail
	// 要素が一つしかない場合は両方空にする
	if l.Head == l.Tail {
		l.Head, l.Tail = nil, nil
		return tail
	}

	l.Tail.Prev.Next = nil
	l.Tail = l.Tail.Prev
	tail.Prev = nil // 切り離したノードのリンクを消す
	return tail
}

// LRUCache本体のロジック
type LRUCache struct {
	keyAddr  map[int]*Node // キーからノードを即時に見つける地図
	capacity int           // キャッシュの最大容量
	list     List          // 使用頻度順に並んだリスト
}

// Get(データの取得)
func (l *LRUCache) Get(key int) int {
	node, ok := l.keyAddr[key]
	if !ok {
		return -1
	}
	l.list.MoveToFront(node)
	return node.Val
}

// データの保存・更新
func (l *LRUCache) Put(key int, value int) {
	node, ok := l.keyAddr[key]
	// すでにキーがある場合
	if ok {
		// 値を書き換えて
		node.Val = value
		// 最新へ移動
		l.list.MoveToFront(node)
		return
	}

	// 容量オーバーなら
	if len(l.keyAddr) >= l.capacity {
		// 末尾を消す
		node := l.list.DelLastNode()
		// マップから消す
		delete(l.keyAddr, node.Key)
	}

	// 新規作成
	node = &Node{Key: key, Val: value}
	// マップに登録
	l.keyAddr[node.Key] = node
	// 先頭に追加
	l.list.PushToFront(node)
}
