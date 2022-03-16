/*
590. N叉树的后序遍历

给定一个 N 叉树，返回其节点值的后序遍历。

例如，给定一个 3叉树 :

// Definition for a Node.
class Node {
public:
    int val;
    vector<Node*> children;

    Node() {}

    Node(int _val, vector<Node*> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Solution {
public:
vector<int> preorder(Node* root) {
        vector<int> ret ;
        if(!root) return ret;
        stack<Node*> st;
        st.push(root);
        while(!st.empty()) {

            auto node = st.top();
            st.pop();
            ret.push_back(node->val);
            vector<int>::reverse_iterator rit;
            for(auto rit = (node->children).rbegin(); rit != (node->children).rend(); rit++) {
                st.push(*rit);
            }
        }
        return ret;
    }
};