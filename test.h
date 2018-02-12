//
// Created by hejunwei on 09/02/2018.
//

#ifndef TEST_TEST_H
#define TEST_TEST_H

#include<iostream>
#include<vector>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;

    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    // 将一个排好序的数组构建成平衡二叉查找树
    TreeNode *sortedArrayToBST(vector<int> &nums) {
        if (nums.size() == 0) return NULL;
        TreeNode *root = NULL;
        buildTree(root, nums);
        return root;
    }

    void buildTree(TreeNode *&sub_root, vector<int> &nums) {
        if (nums.size() == 0) return;
        int mid = nums.size() / 2;
        sub_root = new TreeNode(nums[mid]);
        nums.erase(nums.begin() + mid);
        vector<int> left_nums, right_nums;
        left_nums.assign(nums.begin(), nums.begin() + mid);
        right_nums.assign(nums.begin() + mid, nums.end());
        buildTree(sub_root->left, left_nums);
        buildTree(sub_root->right, right_nums);
    }
};

#endif //TEST_TEST_H
