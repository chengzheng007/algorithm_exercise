/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    // 1.复制一份，新节点穿插在原始结点后面
    if head == nil {
        return head
    }
    
    // var (
    //     curr = head
    //     next *Node
    // )
    curr := head
    for curr != nil {
        // next = curr.Next
        // random := &Node{Val:curr.Val}
        // random.Next = next
        // curr.Next = random
        // curr = next
        random := &Node{Val:curr.Val}
        random.Next = curr.Next
        curr.Next = random
        curr = curr.Next.Next
    }   
    
    // 2.改写随机结点中的random结点，使之指向新生成链表中的结点
    curr = head 
    for curr != nil {
        if curr.Random != nil {
            // 精髓，全部结点都附加在老结点后面，新复制的随机结点
            // 肯定也在老链表对应随机结点后面
            curr.Next.Random = curr.Random.Next 
        }
        curr = curr.Next.Next
    }
    
    // 3.处理新旧交替的链表，将新复制的摘出来，老的恢复
    copyhead := &Node{}
    copyIter := copyhead
    curr = head
    for curr != nil {
        // 摘出复制的结点
        copyIter.Next = curr.Next
        copyIter = copyIter.Next
        // 断开老链表结点
        curr.Next = curr.Next.Next
        curr = curr.Next
    }
    
    return copyhead.Next
}
