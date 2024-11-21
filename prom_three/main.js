/**
 * @param {string} s
 * @return {number}
 */
let lengthOfLongestSubstring = (s) => {
    let maxLength = 0;
    let str = "";
    for (let i = 0; i < s.length; i++)  {
        if (str.includes(s[i]))    {
            let index = str.indexOf(s[i])
            str = str.slice(index + 1)
        }

        str += s[i];

        maxLength = Math.max(maxLength, str.length);

    }

    return maxLength
};


let sa = "abcabcbb"
let result = lengthOfLongestSubstring(sa)

console.log(result)
