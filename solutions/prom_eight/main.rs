enum Sign {
    Positive,
    Negative
}

struct Solution;

impl Solution {
    pub fn my_atoi(s: String) -> i32 {
        let s = s.trim();
        let mut chars = s.chars().peekable();
        let mut sign = Sign::Positive;
        let mut result: i32 = 0;

        if let Some(&c) = chars.peek() {
             if c == '-' {
                sign = Sign::Negative;
                chars.next();
            } else if c == '+' {
                sign = Sign::Positive;
                chars.next();
            }
        }

        while let Some(c) = chars.next() {
            println!("{:#?}", c);
            if c == '0' && result == 0 {
                continue;
            }
            if let Some(digit) = c.to_digit(10) {
                if let Some(new_result) = result.checked_mul(10).and_then(|r| r.checked_add(digit as i32)) {
                    result = new_result;
                } else {
                    return match sign {
                        Sign::Positive => i32::MAX,
                        Sign::Negative => i32::MIN,
                    };
                }
            } else {
                break;
            }
        }

        match sign {
            Sign::Positive => result,
            Sign::Negative => -result,
        }
    }
}

fn main()   {
    let res = Solution::my_atoi("          -043".to_string());
    println!("{:#?}", res)
  }
