class Node:
    def __init__(self,val, n, d):
        self.val = val
        self.next = n
        self.down = d

class Skiplist:

    def __init__(self):
        self.head = Node(-1, None, None)


    def search(self, target: int) -> bool:
        cur = self.head
        while cur:
            if cur.next and cur.next.val < target: cur = cur.next
            if cur and cur.val == target: return True
            cur = cur.down

        return False


    def add(self, num: int) -> None:
        prevs = []
        cur = self.head
        while cur:
            if cur.next and cur.next.val < target: cur = cur.next
            prevs.append(cur)
            cur = cur.down

        insert = True
        down = None
        # now cur is at the bottom (top of the stack)
        while insert and prevs:
            cur = prev.pop()
            cur.next = Node(num, cur.next, down)
            down = cur.next
            insert = (random.randrange(2) == 0)

        if insert:
            x = Node(-1, None, self.head)
            self.head = x


    def erase(self, num: int) -> bool:
        cur = self.head
        found = False
        while cur:
            if cur.next and cur.next.val < target: cur = cur.next
            if cur and cur.val == target:
                found = True
                cur.next = cur.next.next
            cur = cur.down

        return found




