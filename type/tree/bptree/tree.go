package bptree

import (
    "container/list"
    "fmt"
    "unsafe"
)

const (
    DEFAULT_ORDER = 1 << 10 // 1024 2048 4096
    //DEFAULT_ORDER = 4 // 1024 2048 4096
    // 每次节点都有order个节点
    order = DEFAULT_ORDER

    nextLeafIndex     = order - 1
    leafMaxIndex      = order - 1
    indexNodeMaxIndex = order - 1
)

/*tip: tree
                                                  root
                                      child[node_a,node_b,node_c]
                          /                          |                        \
                      nodea                       node_b                      node_c
  child[value_a,value_b....  node_b]    child[value_c value_d ...... node_c]   child[value_e value_f ....... 0]
*/

/*
tip:                                        那么索引最大值就是order-2
    规则A:在叶子节点中,key和child节点的关系对应的,value节点最多数量为order-1,最后一个child指针指向右边的叶子节点(如果为空,那么该节点已经处于最右边)
    规则B:索引从0开始,指针在i+1处指向带有键的子树,如果key值在index[0]和index[1]之间,那么选择的child就是child[0]
    规则C:在叶子节点中有效指针数为num_keys,
    规则D:索引节点中有效指针为num_keys+1
    规则E:num_keys字段用来跟踪当前节点的有效指针数(数组都为预分配)
*/

type BpTree struct {
    root *node
    pool *resourcePool // 8byte
}

/*
tip: <<<<<<<<<<<<<<-打印相关函数->>>>>>>>>>>>>
*/

/*
desc: 返回child到root的长度
*/
func (tree *BpTree) pathToRoot(child *node) uint64 {
    c := child
    var length uint64
    for c != tree.root {
        c = c.Parent
        length++
    }
    return length

}

/*
desc: 返回树的高度
*/
func (tree *BpTree) height() uint64 {
    var cnt uint64
    root := tree.root
    for root != nil {
        root = (*node)(root.child[0])
        cnt++
    }
    return cnt
}

/*
desc: 打印所有叶子节点
*/
func (tree *BpTree) PrintLeaves() {
    if tree.root == nil {
        return
    }

    cur := tree.root

    // 递归找到叶子节点
    for !cur.IsLeaf {
        cur = (*node)(cur.child[0])
    }

    for {
        for i := 0; i < cur.numKeys; i++ {
            fmt.Print(cur.childKeys[i], (*value)(cur.child[i]).val, "\t")
        }

        /*
           case:如果child[order-1]等于nil的话,那么说明已经遍历完毕,否则继续获取下一个叶子节点
                nextLeafIndex = order - 1
        */
        if cur.child[nextLeafIndex] != nil {
            cur = (*node)(cur.child[nextLeafIndex])
        } else {
            break
        }
    }
    fmt.Println()
}

/*
desc: 打印整个tree
*/
func (tree *BpTree) PrintTree() {

    root := tree.root
    if root == nil {
        return
    }

    var (
        depth uint64
        l     = list.New()
    )

    // 先进先出
    l.PushBack(root)
    for l.Len() > 0 {
        l.Front()

        el := l.Front()
        n := el.Value.(*node)
        l.Remove(el)
        if n.Parent != nil && unsafe.Pointer(n) == n.Parent.child[0] {
            curDepth := tree.pathToRoot(n)
            if curDepth != depth {
                depth = curDepth // 更新深度
                println()        // 换行
            }
        }

        if n.IsLeaf {
            print("leaf: ")
        } else {
            print("索引: ")
        }

        for i := 0; i < n.numKeys; i++ {
            fmt.Printf("%d ", n.childKeys[i])
        }

        if !n.IsLeaf {
            for i := 0; i <= n.numKeys; i++ {
                l.PushBack((*node)(n.child[i])) // 入队
            }
        }

        print("|  ")
    }
    println()
}

/*
desc: 打印单个value节点
*/
func (tree *BpTree) findAndPrint(key uint64) {
    val := tree.find(key, nil)
    if val != nil {
        fmt.Printf("%p %d %v \n", val, key, val.val)
    } else {
        println("not found")
    }
}

/*
desc: 范围遍历
*/
func (tree *BpTree) findAndPrintRange(start, end uint64) {
    key, arr := tree.findRange(start, end)
    if len(arr) > 0 {
        for i := 0; i < len(arr); i++ {
            fmt.Printf("Key: %d   Location: %p  Value: %d \n", key[i], arr[i], arr[i].val)
        }
    }
}

/*
tip: <<<<<<<<<<<<<<-查找相关函数->>>>>>>>>>>>>
*/

/*
desc: 返回value
*/
func (tree *BpTree) find(key uint64, leafOut **node) *value {
    if tree.root == nil {
        return nil
    }

    var (
        i    int
        leaf = tree.findLeaf(key)
    )

    for i = 0; i < leaf.numKeys; i++ {
        if leaf.childKeys[i] == key {
            break
        }
    }

    if leafOut != nil {
        *leafOut = leaf
    }

    /*
       case: 如果i< left.numKeys,那么说明找到了
    */
    if i < leaf.numKeys {
        return (*value)(leaf.child[i])
    }

    return nil
}

/*
desc: 返回Start-end之间的节点
*/
func (tree *BpTree) findRange(start, end uint64) ([]uint64, []*value) {
    // 找到叶子节点
    leaf := tree.findLeaf(start)
    if leaf == nil {
        return nil, nil
    }
    var i int
    /*
       case: 在子keys寻找指定节点
    */
    for i < leaf.numKeys && leaf.childKeys[i] < start {
        i++
    }

    // 没有找到
    if i == leaf.numKeys {
        return nil, nil
    }

    var arr []*value
    var key []uint64

    /*
       case: 叶子节点
    */
    for leaf != nil {
        /*
           case: 将所有符合条件value节点都加入数组
        */
        for ; i < leaf.numKeys && leaf.childKeys[i] <= end; i++ {
            key = append(key, leaf.childKeys[i])
            arr = append(arr, (*value)(leaf.child[i]))
        }
        // 那么说明没有遍历到尾部,也不需要扫描接下来的叶子节点了
        if i != leaf.numKeys {
            break
        }
        leaf = (*node)(leaf.child[order-1]) // 获取下一个叶子节点
        i = 0
    }

    return key, arr
}

/*
desc: 返回叶子节点
*/
func (tree *BpTree) findLeaf(key uint64) *node {
    if tree.root == nil {
        return nil
    }

    n := tree.root
    for !n.IsLeaf {
        i := 0

        for i < n.numKeys { //寻找子节点位置
            if key >= n.childKeys[i] {
                i++
            } else {
                break
            }
        }
        n = (*node)(n.child[i])
    }

    return n
}

/*
desc: 返回node在parent subChild中的下标
*/
func (tree *BpTree) GetLeftIndex(parent, left *node) int {
    var leftIndex = 0
    for leftIndex <= parent.numKeys && parent.child[leftIndex] != unsafe.Pointer(left) {
        leftIndex++
    }
    return leftIndex
}

/*
desc: 返回邻居节点下标 如果没有邻居那么是返回-1
*/
func (tree *BpTree) getNeighborIndex(n *node) (int, error) {

    for i := 0; i < n.Parent.numKeys; i++ {
        if n.Parent.child[i] == unsafe.Pointer(n) {
            return i - 1, nil
        }
    }

    return 0, notFoundNeighborError
}

/*
tip: <<<<<<<<<<<<<<-插入相关函数->>>>>>>>>>>>>
*/

/*
desc: 将k、v 插入叶子节点
*/
func (tree *BpTree) insertIntoLeaf(leaf *node, key uint64, val *value) *node {
    var (
        insertionPoint int
    )

    // 寻找插入位置
    for insertionPoint < leaf.numKeys && leaf.childKeys[insertionPoint] < key {
        insertionPoint++
    }

    for i := leaf.numKeys; i > insertionPoint; i-- {
        // 插入位置之后的所有元素向前移动一格
        leaf.childKeys[i] = leaf.childKeys[i-1]
        leaf.child[i] = leaf.child[i-1]
    }

    leaf.childKeys[insertionPoint] = key
    leaf.child[insertionPoint] = unsafe.Pointer(val)
    leaf.numKeys++
    return leaf
}

/*
desc: 插入后,叶子节点分裂
*/
func (tree *BpTree) insertIntoLeafAfterSplitting(leaf *node, key uint64, record *value) {

    if !leaf.IsLeaf {
        return
    }

    newLeaf := tree.pool.GetLeaf()
    tmpKey := tree.pool.GetTempSlice()
    tmpChild := tree.pool.GetTempSlice()

    insertionIndex := 0
    // 找到插入位置,既然是页分裂,那么该叶子节点必然已经满了
    for insertionIndex < leafMaxIndex && leaf.childKeys[insertionIndex] < key {
        insertionIndex++
    }

    var i, j int
    for i < leaf.numKeys {
        // skip 插入位置
        if j == insertionIndex {
            j++
        }
        tmpChild[j] = leaf.child[i]
        tmpKey[j] = leaf.childKeys[i]
        i++
        j++
    }

    tmpKey[insertionIndex] = key
    tmpChild[insertionIndex] = unsafe.Pointer(record)

    leaf.numKeys = 0
    split := cut(leafMaxIndex)

    // 页分裂,重置节点
    for i := 0; i < split; i++ {
        leaf.childKeys[i] = tmpKey[i].(uint64)
        leaf.child[i] = tmpChild[i].(unsafe.Pointer)
        leaf.numKeys++
    }

    i = split
    j = 0
    // j 是newNode的索引
    for i < order {
        newLeaf.childKeys[j] = tmpKey[i].(uint64)
        newLeaf.child[j] = tmpChild[i].(unsafe.Pointer)
        newLeaf.numKeys++
        i++
        j++
    }

    tree.pool.PutTempSlice(tmpKey)
    tree.pool.PutTempSlice(tmpChild)

    // before: oldLeaf      -      oldRightLeaf
    // after: oldLeaf - newLeaf - oldRightLeaf
    newLeaf.child[nextLeafIndex] = leaf.child[nextLeafIndex]
    leaf.child[nextLeafIndex] = unsafe.Pointer(newLeaf)

    // 将分裂出去的child节点设置为空
    for i := leaf.numKeys; i < leafMaxIndex; i++ {
        leaf.child[i] = nil
    }

    for i := newLeaf.numKeys; i < leafMaxIndex; i++ {
        newLeaf.child[i] = nil
    }

    // 设置共同父节点
    newLeaf.Parent = leaf.Parent
    // 当前节点的子节点的最小key
    newKey := newLeaf.childKeys[0]

    tree.insertIntoParent(leaf, newKey, newLeaf)
}

/*
desc: 插入后,node节点分裂
*/
func (tree *BpTree) insertIntoNodeAfterSplitting(oldNode, right *node, leftIndex int, key uint64) {

    tmpChild := tree.pool.GetTempSlice()
    tmpKey := tree.pool.GetTempSlice()

    var i, j int

    for i < oldNode.numKeys+1 {
        // 插入位置,跳过
        if j == leftIndex+1 {
            j++
        }
        tmpChild[j] = oldNode.child[i]
        j++
        i++
    }

    i = 0
    j = 0
    // childkey比child少一个
    for i < oldNode.numKeys {
        if j == leftIndex {
            j++
        }
        tmpKey[j] = oldNode.childKeys[i]
        j++
        i++
    }

    tmpChild[leftIndex+1] = unsafe.Pointer(right)
    tmpKey[leftIndex] = key

    //开始分裂
    split := cut(order)
    newNode := tree.pool.GetNode()
    oldNode.numKeys = 0

    for i = 0; i < split-1; i++ {
        oldNode.childKeys[i] = tmpKey[i].(uint64)
        oldNode.child[i] = tmpChild[i].(unsafe.Pointer)
        oldNode.numKeys++
    }

    oldNode.child[i] = tmpChild[i].(unsafe.Pointer)
    kPrime := tmpKey[split-1].(uint64)

    // j 是newnode的下标
    j = 0
    i++
    for i < order {
        newNode.child[j] = tmpChild[i].(unsafe.Pointer)
        newNode.childKeys[j] = tmpKey[i].(uint64)
        newNode.numKeys++
        i++
        j++
    }

    // 最后一个节点是链接的节点 这里的i等于order
    newNode.child[j] = tmpChild[i].(unsafe.Pointer)

    tree.pool.PutTempSlice(tmpChild)
    tree.pool.PutTempSlice(tmpKey)

    newNode.Parent = oldNode.Parent
    // 遍历子节点
    for i := 0; i <= newNode.numKeys; i++ {
        (*node)(newNode.child[i]).Parent = newNode
    }
    // 将新parent节点与祖先节点相关联
    tree.insertIntoParent(oldNode, kPrime, newNode)
}

/*
desc: 将right节点和parent节点关联
*/
func (tree *BpTree) insertIntoParent(left *node, key uint64, right *node) {
    parent := left.Parent
    if parent == nil { // 如果为空,俺么插入新root节点
        tree.insertIntoNewRoot(left, right, key)
        return
    }

    /*
       case:查找left在父节点child中的下标
    */
    leftIndex := tree.GetLeftIndex(parent, left)
    /*
       case:如果有位置,那么直接插入
    */
    if parent.numKeys < indexNodeMaxIndex {
        tree.insertIntoNode(parent, right, leftIndex, key)
        return
    }

    /*
       case:那么就需要分裂
    */
    tree.insertIntoNodeAfterSplitting(parent, right, leftIndex, key)
}

/*
desc: 将right节点插入n(parent)节点
*/
func (tree *BpTree) insertIntoNode(n, right *node, leftIndex int, key uint64, ) {
    for i := n.numKeys; i > leftIndex; i-- {
        // 整体向前移动一步
        n.child[i+1] = n.child[i]
        n.childKeys[i] = n.childKeys[i-1]
    }

    n.child[leftIndex+1] = unsafe.Pointer(right)
    n.childKeys[leftIndex] = key
    n.numKeys++

}

/*
desc: 插入新root
*/
func (tree *BpTree) insertIntoNewRoot(left, right *node, key uint64) {
    tree.root = tree.pool.GetNode()
    tree.root.childKeys[0] = key
    tree.root.child[0] = unsafe.Pointer(left)
    tree.root.child[1] = unsafe.Pointer(right)
    tree.root.numKeys++
    tree.root.Parent = nil
    left.Parent = tree.root
    right.Parent = tree.root
}

/*
desc: 创建新root
*/
func (tree *BpTree) startNewTree(key uint64, value *value) {
    tree.root = tree.pool.GetLeaf()
    tree.root.childKeys[0] = key
    tree.root.child[0] = unsafe.Pointer(value)
    tree.root.numKeys++
}

/*
tip: <<<<<<<<<<<<<<-删除相关函数->>>>>>>>>>>>>
*/

/*
desc: 从n的subChild中删除entry
*/
func (tree *BpTree) removeEntryFromNode(n *node, entry unsafe.Pointer, key uint64) {

    i := 0
    /*
       case:先找到节点
    */
    for n.childKeys[i] != key {
        i++
    }

    i++
    /*
       case:将要删除的位置往后的元素都向前移动一格(覆盖要删除的节点)
    */
    for ; i < n.numKeys; i++ {
        n.childKeys[i-1] = n.childKeys[i]
    }

    // 确定需要扫描的数量
    numPointer := n.numKeys
    if !n.IsLeaf {
        numPointer++
    }

    i = 0
    // key是可以重复的
    for n.child[i] != entry {
        i++
    }

    i++
    /*
       case:将要删除的位置往后的元素都向前移动一格(覆盖要删除的节点)
    */
    for ; i < numPointer; i++ {
        n.child[i-1] = n.child[i]
    }

    n.numKeys--

    if n.IsLeaf {
        for i := n.numKeys; i < order-1; i++ {
            n.child[i] = nil
        }
    } else {
        for i := n.numKeys + 1; i < order; i++ {
            n.child[i] = nil
        }
    }

}

/*
desc: 删除一个entry,并且调整树
*/
func (tree *BpTree) deleteEntry(n *node, entry unsafe.Pointer, key uint64) error {

    tree.removeEntryFromNode(n, entry, key)

    /*
       case: 如果n是root节点
    */
    if n == tree.root {
        tree.adjustRoot()
        return nil
    }

    var minKeys int
    // node节点多一个
    minKeys = cut(order) - 1
    if n.IsLeaf { // 叶子节点最后一个是链接的node节点
        minKeys = cut(order - 1)
    }

    /*
       case: 如果n子节点数量大于minKeys,那么不需要调整
    */
    if n.numKeys >= minKeys {
        return nil
    }

    /*
       case: 小于那么就需要调整了,获取n的邻居节点
    */
    neighborIndex, err := tree.getNeighborIndex(n)
    if err == nil {
        return err
    }

    kPrimeIndex := neighborIndex
    /*
       case:如果n处于child[0],不能让kPrimeIndex为负值
    */
    if neighborIndex == -1 {
        kPrimeIndex = 0
    }

    kPrime := n.Parent.childKeys[kPrimeIndex]

    var neighbor *node
    /*
       case:如果neighborIndex是负一,那么邻居节点必然为1
    */
    if neighborIndex == -1 {
        neighbor = (*node)(n.Parent.child[1])
    } else {
        neighbor = (*node)(n.Parent.child[neighborIndex])
    }

    /*
       case: 索引节点,cap = order - 1
       case: 叶子节点,那么cap = order
    */
    var capacity int
    capacity = order - 1
    if n.IsLeaf {
        capacity = order
    }

    /*
       case: 如果小于单个节点的容量,那么就合并
    */
    if neighbor.numKeys+n.numKeys < capacity {
        tree.coalesceNode(n, neighbor, neighborIndex, kPrime)
        return nil
    }
    /*
       case: 否则,就调整节点
    */
    tree.redistributeNodes(n, neighbor, neighborIndex, kPrimeIndex, kPrime)
    return nil
}

/*
tip: <<<<<<<<<<<<<<-调整相关函数->>>>>>>>>>>>>
*/

/*
desc: 调整root
*/
func (tree *BpTree) adjustRoot() {

    if tree.root == nil {
        return
    }

    /*
       case:child不为空,直接返回
    */
    if tree.root.numKeys > 0 {
        return
    }

    /*
       case:如果不是叶子节点(子节点是value),且子节点不为空,那么提拔第一个为新root
    */
    if !tree.root.IsLeaf {
        n := tree.root
        tree.root = (*node)(tree.root.child[0])
        tree.root.Parent = nil

        // 内存池优化
        tree.pool.PutNode(n)
    } else { // 如果是叶子节点,那么说明整个树都是空的
        tree.root = nil
    }

    return

}

/*
desc: 合并节点
*/
func (tree *BpTree) coalesceNode(n, neighbor *node, neighborIndex int, kPrime uint64) {

    /*
       case:如果n节点位于数组的[0]的位置,那么neighborIndex等于-1。
            因为bpt的规则是,左边必然比右边小。然后n处于最0的位置,那么这里就需要交换一下指针
            最后的结果就是将neighbor合并到n上。
    */
    if neighborIndex == -1 {
        n, neighbor = neighbor, n
    }

    neighborInsertionIndex := neighbor.numKeys

    /*
       case:非叶子节点,直接添加到后面
    */
    if !n.IsLeaf {
        neighbor.childKeys[neighborInsertionIndex] = kPrime
        neighbor.numKeys++

        var (
            nEnd = n.numKeys
            j    = 0
            i    = neighborInsertionIndex + 1
        )

        // 将n合并到neighbor
        for j < nEnd {
            neighbor.childKeys[i] = n.childKeys[j]
            neighbor.child[i] = n.child[j]
            neighbor.numKeys++
            n.numKeys--

            i++
            j++
        }

        // 指针的数量永远比key多一个
        neighbor.child[i] = n.child[j]

        /*
           case:设置child的parent
        */
        for i := 0; i < neighbor.numKeys; i++ {
            (*node)(neighbor.child[i]).Parent = neighbor
        }

    } else {
        /*
           case:叶子节点
        */
        var (
            j = 0
            i = neighborInsertionIndex
        )

        for j < n.numKeys {

            neighbor.childKeys[i] = n.childKeys[j]
            neighbor.child[i] = n.child[j]
            neighbor.numKeys++
            n.numKeys--

            // 下标移动
            i++
            j++
        }

        // 连接底层value的node
        neighbor.child[nextLeafIndex] = n.child[nextLeafIndex]
    }

    tree.deleteEntry(n.Parent, unsafe.Pointer(n), kPrime)

    // 内存池优化
    tree.pool.PutNode(n)
}

/*
desc:  重新分配child节点
*/
func (tree *BpTree) redistributeNodes(n, neighbor *node, neighborIndex int, kPrimeIndex int, kPrime uint64) {

    /*
       case:邻居在右边( 0 - 1 = -1),n处于数组[0]
    */

    if neighborIndex != -1 {
        /*
           case: 不是叶子节点,那么需要保存最后一个key
        */
        if !n.IsLeaf {
            n.child[n.numKeys+1] = n.child[n.numKeys]
        }

        /*
           case: 整体向后移动一格,那么0和1的值是一致的
        */
        for i := n.numKeys; i > 0; i-- {
            n.child[i] = n.child[i-1]
            n.childKeys[i] = n.childKeys[i-1]
        }

        /*
           case:
        */
        if !n.IsLeaf {
            // 如果是node节点,那么数量比叶子节点多一个键值对
            n.child[0] = neighbor.child[neighbor.numKeys]
            (*node)(n.child[0]).Parent = n
            n.childKeys[0] = kPrime

            neighbor.child[neighbor.numKeys] = nil

            // 重新设置邻居节点在父节点的最大值(因为最后一个节点移动到b节点了)
            n.Parent.childKeys[kPrimeIndex] = neighbor.childKeys[neighbor.numKeys-1]
        } else {
            // 叶子节点比node少一个键值对
            n.child[0] = neighbor.child[neighbor.numKeys-1] //获取最后一个叶子节点
            neighbor.child[neighbor.numKeys-1] = nil
            n.childKeys[0] = neighbor.childKeys[neighbor.numKeys-1] //获取最后一个叶子节点
            n.Parent.childKeys[kPrimeIndex] = n.childKeys[0]

            //(*node)(n.child[0]).Parent = n                          // 重新设置parent节点

        }
    } else {
        /*
           邻居在右边,从右边的邻居取一对键指针,增加到n上
        */

        /*
           case: 索引节点:
           case: 否则:
        */

        if !n.IsLeaf {
            n.child[n.numKeys] = neighbor.child[0]
            n.childKeys[n.numKeys] = neighbor.childKeys[0]
            n.Parent.childKeys[kPrimeIndex] = neighbor.childKeys[1]
        } else {
            n.childKeys[n.numKeys] = kPrime
            n.child[n.numKeys+1] = neighbor.child[0]
            (*node)(n.child[n.numKeys+1]).Parent = n
            n.Parent.childKeys[kPrimeIndex] = neighbor.childKeys[0]
        }

        var i int

        /*
           case: 向前移动一格
        */
        for i = 0; i < neighbor.numKeys-1; i++ {
            neighbor.childKeys[i] = neighbor.childKeys[i+1]
            neighbor.child[i] = neighbor.child[i+1]
        }

        /*
           case: 如果找到,那么就直接删除
        */
        if !n.IsLeaf {
            neighbor.child[i] = neighbor.child[i+1]
        }

    }

    n.numKeys++
    neighbor.numKeys--

}

/*
tip: <<<<<<<<<<<<<<-对外暴露函数->>>>>>>>>>>>>
*/

func (tree *BpTree) Delete(key uint64) {

    var leaf *node

    value := tree.find(key, &leaf)
    /*
       case: 如果找到,那么就直接删除
    */
    if value != nil && leaf != nil {
        tree.deleteEntry(leaf, unsafe.Pointer(value), key)

        // 内存池优化
        tree.pool.PutValue(value)
    }

    return
}

func (tree *BpTree) Insert(key uint64, value uint64) {

    valueNode := tree.find(key, nil)

    /*
       case: value节点存在,直接更新
    */
    if valueNode != nil { // 如果找到直接更新
        valueNode.val = value
        return
    }

    valueNode = tree.pool.GetValue()
    valueNode.val = value

    /*
       case: root不存在,创建
    */

    if tree.root == nil {
        tree.startNewTree(key, valueNode)
        return
    }

    /*
       case: 根据key寻找对应的叶子节点
    */
    leaf := tree.findLeaf(key)

    /*
       case: 如果叶子节点有足够位置,直接插入
    */
    if leaf.numKeys < leafMaxIndex {
        tree.insertIntoLeaf(leaf, key, valueNode)
        return
    }

    /*
       case: 空间不足,就得页分裂了.
    */
    tree.insertIntoLeafAfterSplitting(leaf, key, valueNode)
}

/*
desc:实际上将指针置为空,那么gc就可以正常运行了
*/
func (tree *BpTree) DestroyTree() {
    tree.root = nil
    tree.pool = nil
}

func NewBpTree() *BpTree {
    bpTree := &BpTree{
        root: nil,
        pool: defaultResourcePool,
    }
    return bpTree
}
