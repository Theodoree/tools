package _600_1699


#include "unordered_set"

struct TreeNode {
int val;
TreeNode *left;
TreeNode *right;
TreeNode() : val(0), left(nullptr), right(nullptr) {}
TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};


class Solution {
std::unordered_set<int> set;

public:
TreeNode* correctBinaryTree(TreeNode* root) {
if (!root) {
return nullptr;
}

DRL(root,root);
return root;
}

void DRL(TreeNode* root,TreeNode* parent) {
if(!root) {
return;
}



if (root->left){
auto iter = set.find(root->left->val);
if (iter != set.end()){
if (root == parent->left){
parent->left = nullptr;
}else{
parent->right = nullptr;
}

return ;
}
}

if (root->right){
auto iter = set.find(root->right->val);
if (iter != set.end()){
if (root == parent->left){
parent->left = nullptr;
}else{
parent->right = nullptr;
}
return;
}
}


set.insert(root->val);
DRL(root->right,root);
DRL(root->left,root);


}
};

class Solution { // 大概思路就是用队列去level遍历,从左边到右边,那么就不需要遍历到叶子节点
public:
TreeNode* correctBinaryTree(TreeNode* root) {
queue<pair<TreeNode*,TreeNode*>> q;
q.push({root, nullptr});

unordered_set<int> nodes;
nodes.insert(root->val);

while (!q.empty()) {
int n = q.size();
for (int i = 0; i < n; ++i) {
auto [node, parent] = q.front();
q.pop();

if (node->left) {
q.push({node->left, node});
nodes.insert(node->left->val);
}
if (node->right) {
if (nodes.find(node->right->val) != nodes.end()) {
if (parent->left == node) {
parent->left = nullptr;
} else {
parent->right = nullptr;
}
return root;
}
q.push({node->right, node});
nodes.insert(node->right->val);
}
}
}
return root;
}
};

