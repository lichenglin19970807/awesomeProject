# 链表

# 单链表

思考需要哪些变量: res、left、right、cur

反转单链表: https://leetcode.cn/problems/reverse-linked-list/description/

合并两个有序链表: https://leetcode.cn/problems/merge-two-sorted-lists/

链表加法: https://leetcode.cn/problems/add-two-numbers/description/

*大数相加溢出, 可以采用链表方式相加*

分隔链表: https://leetcode.cn/problems/partition-list/description/

*链表循环时, 循环内尽量将指向下一个节点的动作前置, 同时用cur指向并操作当前节点: for head != nil { cur := head; head = head.Next }*