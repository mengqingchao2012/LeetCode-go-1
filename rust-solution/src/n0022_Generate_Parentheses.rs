use crate::Solution;

impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        if n < 0 {
            return vec![];
        }
        let mut res = Vec::new();
        Solution::generate(&mut res, "".parse().unwrap(), n, n);
        res
    }

    fn generate(res: &mut Vec<String>, cur: String, left: i32, right: i32) {
        if left == 0 && right == 0 {
            res.push(cur)
        } else {
            if left > 0 {
                Solution::generate(res, cur.clone() + "(", left - 1, right)
            }
            if right > left {
                Solution::generate(res, cur + ")", left, right - 1)
            }
        }
    }
}
