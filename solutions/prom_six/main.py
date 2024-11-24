class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        arr = [""] * numRows
        i, j = 0, 1

        for v in s:
            arr[i] += v
            if i == 0:
                j = 1
            elif i == numRows - 1:
                j = -1
            i += j


        return "".join(arr)

st = "PAYPALISHIRING"
n = 4

res = Solution.longestPalindrome(Solution, st, 4)

print(res)
