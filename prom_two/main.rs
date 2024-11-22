#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

struct Solution;

impl Solution {
  pub fn add_two_numbers(l1: Option<Box<ListNode>>, l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
      match (l1, l2)  {
          (None, None) => None,
          (Some(v), None) | (None, Some(v)) => Some(v),
          (Some(v1), Some(v2)) => {
              let sum = v1.val + v2.val;
              let mut result = ListNode::new(sum);

              if sum < 10 {
                  result.next = Solution::add_two_numbers(v1.next, v2.next);
              } else {
                  let carry = Some(Box::new(ListNode::new(1)));
                  result.val = sum - 10;
                  result.next = Solution::add_two_numbers(Solution::add_two_numbers(carry, v1.next), v2.next)
              }

              Some(Box::new(result))
          }
          (_, _) => todo!()

      }
  }
}

fn main()   {
  let res = Solution::add_two_numbers(Some(Box::new(ListNode::new(1))), Some(Box::new(ListNode::new(2))));
  println!("{:#?}", res)
}
