#include "debug.hpp"
using namespace std;
/*
894. 所有可能的真二叉树
中等
相关标签
相关企业

给你一个整数 n ，请你找出所有可能含 n 个节点的 真二叉树
，并以列表形式返回。答案中每棵树的每个节点都必须符合 Node.val == 0 。

答案的每个元素都是一棵真二叉树的根节点。你可以按 任意顺序
返回最终的真二叉树列表。

真二叉树 是一类二叉树，树中每个节点恰好有 0 或 2 个子节点。



示例 1：

输入：n = 7
输出：[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]

示例 2：

输入：n = 3
输出：[[0,0,0]]



提示：

    1 <= n <= 20
*/

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left),
 * right(right) {}
 * };
 */
class Solution {
public:
  vector<TreeNode *> allPossibleFBT(int n) {
    vector<TreeNode *> res;
    if (n % 2 == 0) {
      return res;
    }
    if (n == 1) {
      return vector<TreeNode *>{new TreeNode(0, nullptr, nullptr)};
    }
    for (int i = 1; i < n; i += 2) {
      auto &&left = allPossibleFBT(i);
      auto &&right = allPossibleFBT(n - 1 - i);
      for (auto l : left) {
        for (auto r : right) {
          TreeNode *root = new TreeNode(0);
          root->left = l;
          root->right = r;
          res.push_back(root);
        }
      }
    }
    printf("n = %d\n", n);
    return res;
  }
};

int main() {
  Solution sl;
  auto res = sl.allPossibleFBT(7);
  printf("res: %zu\n", res.size());
  for (auto x : res) {
    printf("[");
    print(x);
    printf("]");
    printf("\n");
  }
}
