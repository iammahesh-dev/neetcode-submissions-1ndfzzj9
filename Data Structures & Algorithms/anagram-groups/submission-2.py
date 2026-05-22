class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        map = {}
        for s in strs:
            count = [0] * 26
            for ch in s:
                index = ord(ch) - ord('a')
                count[index] = count[index] + 1
            keys = tuple(count)
            if keys in map:
                map[keys].append(s)
            else:
                map[keys] = [s]

        return list(map.values())