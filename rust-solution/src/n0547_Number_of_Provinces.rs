use crate::Solution;

impl Solution {
    fn find(x: usize, p: &mut [usize]) -> usize {
        if p[x] != x {
            p[x] = Solution::find(p[x], p);
        }
        p[x]
    }

    pub fn find_circle_num(is_connected: Vec<Vec<i32>>) -> i32 {
        let n = is_connected.len();
        let mut p = (0..n).into_iter().collect::<Vec<usize>>();

        let mut cnt = n;
        for i in 0..n {
            for j in 0..n {
                let parent_i = Solution::find(i, p.as_mut_slice());
                let parent_j = Solution::find(j, p.as_mut_slice());
                if is_connected[i][j] == 1 && parent_i != parent_j {
                    p[parent_i] = parent_j;
                    cnt -= 1;
                }
            }
        }
        cnt as i32
    }
}