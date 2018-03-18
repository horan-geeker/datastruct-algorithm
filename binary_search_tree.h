//
// Created by hejunwei on 01/02/2018.
//

#ifndef TEST_BINARY_SEARCH_TREE_H
#define TEST_BINARY_SEARCH_TREE_H

#include "binary_tree.h"

template<class Record>
class BinarySearchTree : public BinaryTree<Record> {
public:
    int depth() {
        return this->recursive_depth(this->root);
    }

    int recursive_depth(BinaryNode<Record>*& node) {
        int deep = 0;
        if (node != NULL) {
            int lchilddeep = recursive_depth(node->left);
            int rchilddeep = recursive_depth(node->right);
            deep = lchilddeep > rchilddeep ? lchilddeep + 1 : rchilddeep + 1;
        }
        return deep;
    }

    void print_node_by_level() {
        int parentSize = 1, childSize = 0;
        BinaryNode<Record> * temp;
        queue<BinaryNode<Record> *> q;
        q.push(this->root);
        do
        {
            temp = q.front();
            cout << temp->data << "  ";
            q.pop();

            if (temp->left != NULL)
            {
                q.push(temp->left);
                childSize ++;
            }
            if (temp->right != NULL)
            {
                q.push(temp->right);
                childSize ++;
            }

            parentSize--;
            if (parentSize == 0)
            {
                parentSize = childSize;
                childSize = 0;
                cout << endl;
            }

        } while (!q.empty());
    }

    /*
     * 如果找到了返回true，找不到返回false
     */
    bool tree_search(Record &target) const {
        return recursive_search(this->root, target);
    }

    void travel_by_depth(void (*callback)(Record &)) {
        recursive_travel(this->root, callback);
    }

    void recursive_travel(BinaryNode<Record> *&sub_root, void (*callback)(Record &)) {
        if (sub_root != NULL) {
            (*callback)(sub_root->data);
            BinaryNode<Record> *parent = sub_root;
        }
    }

    /*
     * 如果插入重复的key，会返回false
     */
    bool insert(const Record &new_data) {
        return search_and_insert(this->root, new_data);
    }

    bool search_and_insert(BinaryNode<Record> *&sub_root, const Record &new_data) {
        if (sub_root == NULL) {
            sub_root = new BinaryNode<Record>(new_data);
            return true;
        } else if (new_data < sub_root->data) {
            search_and_insert(sub_root->left, new_data);
        } else if (new_data > sub_root->data) {
            search_and_insert(sub_root->right, new_data);
        } else {
            return false;
        }
    }

    bool remove(const Record &target) {
        return search_and_destroy(this->root, target);
    }

    bool search_and_destroy(BinaryNode<Record> *&sub_root, Record target) {
        if (sub_root == NULL || sub_root->data == target) return remove_root(sub_root);
        else if (sub_root->data > target) return search_and_destroy(sub_root->left, target);
        else return search_and_destroy(sub_root->right, target);
    }

    /*
     * 删除传进来的 sub_root 节点
     */
    bool remove_root(BinaryNode<Record> *&sub_root) {
        // 空节点
        if (sub_root == NULL) return false;

        BinaryNode<Record> *to_delete = sub_root;
        // 右子树为空，直接用左子树
        if (sub_root->right == NULL) sub_root = sub_root->left;
            // 左子树为空，用右子树
        else if (sub_root->left == NULL) sub_root = sub_root->right;
            // 两边子树都存在
        else {
            // 暂时定为左子树
            to_delete = sub_root->left;
            BinaryNode<Record> *parent = sub_root;
            // 看左子树的右边是否存在, 并且一直往右边找
            while (to_delete->right != NULL) {
                parent = to_delete;
                to_delete = to_delete->right;
            }
            sub_root->data = to_delete->data; //copy to_delete value to root
            // 当左子树没有右边的话，to_delete->left 是 NULL 或者 有值，交给 sub_root 的 left 就好了
            if (parent == sub_root) sub_root->left = to_delete->left;
            else parent->right = to_delete->left;
        }
        delete to_delete;
        return true;
    }

private:
    bool recursive_search(BinaryNode<Record> *sub_root, const Record target) const {
        if (sub_root == NULL || sub_root->data == target) {
            return sub_root;
        } else if (sub_root->data > target) {
            return recursive_search(sub_root->left, target);
        } else {
            return recursive_search(sub_root->right, target);
        }
    }
};

#endif //TEST_BINARY_SEARCH_TREE_H