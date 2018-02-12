//
// Created by hejunwei on 01/02/2018.
//

#ifndef TEST_BINARY_TREE_H
#define TEST_BINARY_TREE_H

enum Color {
    red, black
};

template<class Entry>
struct BinaryNode {
    Entry data;
    BinaryNode<Entry> *left;
    BinaryNode<Entry> *right;

    BinaryNode() {
        data = NULL;
        left = NULL;
        right = NULL;
    }

    BinaryNode(const Entry &x) {
        data = x;
        left = NULL;
        right = NULL;
    }

    virtual Color get_color() const { return red; }

    virtual void set_color(Color c) {}
};

template<class Entry>
class BinaryTree {
public:
    BinaryTree() { this->root = NULL; }

    void test() { cout << "BT"; }

    bool isEmpty() { return root == NULL; }

    void preorder(void (*callback)(Entry &)) { recursive_preorder(this->root, callback); }

    void inorder(void (*callback)(Entry &)) { recursive_inorder(this->root, callback); }

    void postorder(void (*callback)(Entry &)) { recursive_postorder(this->root, callback); }

private:
    void recursive_preorder(BinaryNode<Entry> *sub_root, void (*callback)(Entry &)) {
        if (sub_root != NULL) {
            (*callback)(sub_root->data);
            recursive_preorder(sub_root->left, callback);
            recursive_preorder(sub_root->right, callback);
        }
    }

    void recursive_inorder(BinaryNode<Entry> *sub_root, void (*callback)(Entry &)) {
        if (sub_root != NULL) {
            recursive_inorder(sub_root->left, callback);
            (*callback)(sub_root->data);
            recursive_inorder(sub_root->right, callback);
        }
    }

    void recursive_postorder(BinaryNode<Entry> *sub_root, void (*callback)(Entry &)) {
        if (sub_root != NULL) {
            recursive_postorder(sub_root->left, callback);
            recursive_postorder(sub_root->right, callback);
            (*callback)(sub_root->data);
        }
    }

protected:
    BinaryNode<Entry> *root;
};

#endif //TEST_BINARY_TREE_H