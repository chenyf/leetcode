class Node(object):
    def __init__(self, val, prev, next, child):
        self.val = val
        self.prev = prev
        self.next = next
        self.child = child

class Solution(object):

    def handle(self, head):
        pass

    def flatten(self, head):
        return head

def addNode(l, node):
    l.next = node
    node.prev = l
    return l

if __name__ == "__main__":
    s = Solution()

    node1 = Node(1, None, None, None)
    node2 = Node(2, None, None, None)

    head = addNode(node1, node2)
    result = s.flatten(head)

    node = result
    while True:
        print node.val
        node = node.next
        if node == None:
            break

