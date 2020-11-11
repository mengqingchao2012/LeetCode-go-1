use crate::Solution;

impl Solution {
    pub fn longest_common_subsequence(text1: String, text2: String) -> i32 {
        let l1 = text1.len();
        let l2 = text2.len();
        if l1 == 0 || l2 == 0 {
            return 0;
        }
        let text1 = text1.as_bytes();
        let text2 = text2.as_bytes();
        let mut dp = vec![vec![0; l2 + 1]; l1 + 1];

        let mut max_len = 0;
        for i in 1..l1 + 1 {
            for j in 1..l2 + 1 {
                if text1[i - 1] == text2[j - 1] {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = std::cmp::max(dp[i - 1][j], dp[i][j - 1]);
                }
                if dp[i][j] > max_len {
                    max_len = dp[i][j];
                }
            }
        }
        max_len
    }
}