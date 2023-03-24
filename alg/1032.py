class Trie:
  def __init__(self):
    self.children = [None for _ in range(26)]
    self.word = False
    
class StreamChecker:
  def __init__(self, words: List[str]):
    def add(p, word):
      for x in word[::-1]:
        if p.children[ord(x) - ord('a')] is None:
          p.children[ord(x) - ord('a')] = Trie()
        p = p.children[ord(x) - ord('a')]
      p.word = True
    self.root = Trie()
    self.stream = ""
    for word in words:
      add(self.root, word)
  
  def query(self, letter: str) -> bool:
    def find(p, stream):
      for x in stream[::-1]:
        cur = p.children[ord(x) - ord('a')]
        if not cur:
          return False
        if cur and cur.word:
          return True
        p = cur
      return False
    self.stream += letter

    return find(self.root, self.stream)



# Your StreamChecker object will be instantiated and called as such:
# obj = StreamChecker(words)
# param_1 = obj.query(letter
