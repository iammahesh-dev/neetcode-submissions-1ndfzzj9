class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        map = {}
        return len(set(nums)) < len(nums)