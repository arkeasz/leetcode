class Solution:
    def reverse(self, x: int) -> int:
        s = str(x)
        re = ""

        if x < 0:
            re = "-"
            s = s[1:]
            s = s[::-1]
            print(re)
        else:
            s = s[::-1]

        re += s
        re = int(re)
        if re < -2**31 or re > 2**31 - 1:
          return 0

        return re

x = 4

res = Solution.longestPalindrome(Solution, x)

print(res)
