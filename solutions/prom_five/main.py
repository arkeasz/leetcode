class Solution:
    def expand_around_center(self, left: int, right: int, s: str) -> str:
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        return s[left + 1:right]

    def longestPalindrome(self, s: str) -> str:
        if not s or len(s) == 1:
            return s


        longest = ""
        for i in range(len(s)):
            odd_palindrome = self.expand_around_center(i, i, s)
            even_palindrome = self.expand_around_center(i, i + 1, s)
            if len(odd_palindrome) > len(longest):
                longest = odd_palindrome
            if len(even_palindrome) > len(longest):
                longest = even_palindrome
        return longest

st = "babad"


res = Solution.longestPalindrome(Solution, st)

print(res)
