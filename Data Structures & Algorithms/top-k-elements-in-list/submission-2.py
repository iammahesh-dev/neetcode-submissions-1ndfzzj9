class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq = {}
        for n in nums:
            freq[n] = freq.get(n, 0) + 1

        bucket = [[] for _ in range(len(nums) + 1)]

        for key, value in freq.items():
            bucket[value].append(key)
        result = []

        for i in range(len(bucket) - 1, -1, -1):
            for v in bucket[i]:
                result.append(v)
                if len(result) == k:
                    return result        