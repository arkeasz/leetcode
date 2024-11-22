
let findMedianSortedArrays = (nums1, nums2) => {
    let nums = nums1.concat(nums2).sort((a,b) => a - b)
    if (nums.length % 2 == 0)  {
        let re = Math.round((nums.length)/2) - 1
        let median = (nums[re] + nums[re+1])/2
        return median
    } else  {
        let re = Math.round((nums.length)/2) - 1
        return nums[re]
    }
};

let nums1 = [1,2,3,4,5]
let nums2 = [6,7,8,9,10,11,12,13,14,15,16,17]

let result = findMedianSortedArrays(nums1, nums2);

console.log(result);
