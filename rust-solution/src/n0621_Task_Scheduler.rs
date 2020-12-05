use crate::Solution;

impl Solution {
    pub fn least_interval(tasks: Vec<char>, n: i32) -> i32 {
        let mut res = 0;
        let mut idle = 0;
        let mut counter = std::collections::HashMap::new();
        for t in tasks {
            *counter.entry(t).or_insert(0) += 1;
        }
        let mut hp: std::collections::BinaryHeap<i32> = counter.values().cloned().collect();

        while !hp.is_empty() {
            let mut remain = Vec::with_capacity(hp.len());
            res += n + 1;
            idle = n + 1 - (hp.len() as i32);
            for _ in 0..n + 1 {
                if let Some(cnt) = hp.pop() {
                    if cnt > 1 {
                        remain.push(cnt - 1);
                    }
                }
            }
            if !remain.is_empty() {
                hp.extend(remain.into_iter())
            }
        }
        res - idle
    }

    pub fn least_interval_math(tasks: Vec<char>, n: i32) -> i32 {
        let mut counter = std::collections::HashMap::new();
        for &t in &tasks {
            *counter.entry(t).or_insert(0) += 1;
        }
        let max_frequency = counter.values().max().unwrap();
        let max_cnt = counter.values().filter(|&x| x == max_frequency).count() as i32;
        (tasks.len() as i32).max((n + 1) * (max_frequency - 1) + max_cnt)
    }
}
