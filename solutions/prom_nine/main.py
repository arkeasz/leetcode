class Solution:
    def isPalindrome(self, s: int) -> bool:
        s = str(s)
        if s == s[::-1]:
            return True
        return False

res = Solution.isPalindrome(Solution, 13345432654)

print(res)
