use crate::Solution;

impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        if nums.is_empty() {
            return vec![-1, -1];
        }

        let start = Solution::binary_search(&nums, target - 1) + 1;
        let end = Solution::binary_search(&nums, target);
        if start >= 0 && start <= end && end < nums.len() as i32 {
            return vec![start, end];
        }
        vec![-1, -1]
    }

    fn binary_search(nums: &[i32], target: i32) -> i32 {
        // 注意这里 low 和 high 都设为 i32 类型，因为后续的 mid - 1 可能为负数，负数
        // 设置给 usize 会得到一个非常大的数
        let mut low: i32 = 0;
        let mut high = (nums.len() - 1) as i32;
        while low <= high {
            let mid = low + (high - low) / 2;
            if nums[mid as usize] > target {
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }
        high
    }
}
