class Solution:

    def encode(self, strs: List[str]) -> str:
        res = ""

        for s in strs:
            res = res + str(len(s)) + "#" + s

        return res

    def decode(self, s: str) -> List[str]:
        strs = []
        i = 0

        while i < len(s):
            j = i
            while s[j] != "#":
                j = j + 1
            l = int(s[i:j])
            i = j + 1
            j = i + l
            word = s[i:j]
            strs.append(word)
            i=j
        return strs