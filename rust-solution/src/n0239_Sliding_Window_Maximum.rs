use crate::Solution;
use std::collections::VecDeque;

impl Solution {
    pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut deque: VecDeque<usize> = VecDeque::with_capacity(k as usize);
        let mut res = Vec::<i32>::new();
        for (idx, num) in nums.iter().enumerate() {
            if !deque.is_empty() && (idx as i32 - k + 1).gt(&(*deque.front().unwrap() as i32)) {
                deque.pop_front();
            }
            while let Some(&i) = deque.back() {
                if nums[i] <= *num {
                    deque.pop_back();
                } else {
                    break;
                }
            }
            deque.push_back(idx);
            if idx as i32 >= k - 1 {
                res.push(nums[*deque.front().unwrap()]);
            }
        }
        res
    }
}
