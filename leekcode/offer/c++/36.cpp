class Solution {
public:
    Node* treeToDoublyList(Node* root) {
        if(!root)
            return nullptr;
        stack<Node*> stk;
        Node* cur = root;
        Node* pre = nullptr;
        Node* faker = new Node(-1); //加一个虚假的节点
        Node* head = faker;
        while(!stk.empty() || cur)
        {
            while(cur)
            {
                stk.push(cur);
                cur = cur->left;
            }
            cur = stk.top();
            stk.pop();
            faker->right = cur;
            cur->left = faker;
            faker = cur;
            cur = cur->right;
        }
        faker->right = head->right;
        head->right->left = faker;

        return head->right;
    }
};