class RandomizedSet:

    def __init__(self):
        self._loc = dict()
        self._nums = list()


    def insert(self, val: int) -> bool:
        if val in self._loc:
            return False
        self._loc[val] = len(self._nums)
        self._nums.append(val)

        return True


    def remove(self, val: int) -> bool:
        if val not in self._loc:
            return False
        l = self._loc[val]
        self._loc[self._nums[-1]] = l
        self._nums[l] = self._nums[-1]
        del self._loc[val]
        self._nums.pop()

        return True


    def getRandom(self) -> int:
        return self._nums[random.randint(0, len(self._nums) - 1)]



