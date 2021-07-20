//
// Created by hejunwei on 08/02/2018.
//

#ifndef TEST_RED_BLACK_TREE_H
#define TEST_RED_BLACK_TREE_H

#include "binary_search_tree.h"
#include "b-tree.h"

enum RB_Code {
    okay, red_node, left_red, right_red, duplicate, rb_error
};

template<class Record>
struct RB_Node : public BinaryNode<Record> {
    Color color;

    RB_Node(const Record &new_entry) {
        color = red;
        this->data = new_entry;
//        left = NULL;
//        right = NULL;
    }

    RB_Node() {
        color = red;
//        left = NULL;
//        right = NULL;
    }

    void set_color(Color c) { color = c; }

    Color get_color() const { return color; }
};

template<class Record>
class RB_Tree : public BinarySearchTree<Record> {
public:
    ErrorCode insert(const Record &new_entry) {
        RB_Code status = rb_insert(this->root, new_entry);
        switch (status) {
            case red_node:
                this->root->set_color(black);
            case okay:
                return success;
            case duplicate:
                return duplicate_error;
            case right_red:
            case left_red:
                return internal_error;
        }
    }

private:
    RB_Code rb_insert(BinaryNode<Record> *&current, const Record &new_entry) {
        RB_Code status, child_status;
        if (current == NULL) {
            current = new RB_Node<Record>(new_entry);
            status = red_node;
        } else if (new_entry == current->data)
            return duplicate;
        else if (new_entry < current->data) {
            child_status = rb_insert(current->left, new_entry);
            status = modify_left(current, child_status);
        } else {
            child_status = rb_insert(current->right, new_entry);
            status = modify_right(current, child_status);
        }
        return status;
    }

    RB_Code modify_left(BinaryNode<Record> *&current, RB_Code &child_status) {
        RB_Code status = okay;
        BinaryNode<Record> *aunt = current->right;
        Color aunt_color = black;
        if (aunt != NULL) aunt_color = aunt->get_color();
        switch (child_status) {
            case okay:
                break;
            case red_node:
                if (current->get_color() == red) status = left_red;
                else status = okay;
                break;
            case left_red:
                if (aunt_color == black) status = rotate_right(current);
                else status = flip_color(current);
                break;
            case right_red:
                if (aunt_color == black) status = double_rotate_right(current);
                else status = flip_color(current);
                break;
        }
        return status;
    }

    RB_Code modify_right(BinaryNode<Record> *&current, RB_Code &child_status) {
        RB_Code status = okay;
        BinaryNode<Record> *left_child = current->left;
        switch (child_status) {
            case okay:
                break;
            case red_node:
                if (current->get_color() == red)
                    status = right_red;
                else
                    status = okay;
                break;
            case right_red:
                if (left_child == NULL)
                    status = rotate_left(current);
                else if (left_child->get_color() == red)
                    status = flip_color(current);
                else
                    status = rotate_left(current);
                break;
            case left_red:
                if (left_child == NULL)
                    status = double_rotate_left(current);
                else if (left_child->get_color() == red)
                    status = flip_color(current);
                else
                    status = double_rotate_left(current);
                break;
        }
        return status;
    }

    RB_Code rotate_left(BinaryNode<Record> *&current) {
        if (current == NULL || current->right == NULL) // impossible cases
            return rb_error;
        else {
            BinaryNode<Record> *right_tree = current->right;
            current->set_color(red);
            right_tree->set_color(black);
            current->right = right_tree->left;
            right_tree->left = current;
            current = right_tree;
        }
        return okay;
    }

    RB_Code rotate_right(BinaryNode<Record> *&current) {
        if (current == NULL || current->left == NULL) // impossible cases
            return rb_error;
        else {
            BinaryNode<Record> *left_tree = current->left;
            current->set_color(red);
            left_tree->set_color(black);
            current->left = left_tree->right;
            left_tree->right = current;
            current = left_tree;
        }
        return okay;
    }

    RB_Code flip_color(BinaryNode<Record> *current) {
        BinaryNode<Record> *left_child = current->left,
                *right_child = current->right;
        current->set_color(red);
        left_child->set_color(black);
        right_child->set_color(black);
        return red_node;
    }

    RB_Code double_rotate_right(BinaryNode<Record> *&current) {
        rotate_left(current->left);
        rotate_right(current);
        return okay;
    }

    RB_Code double_rotate_left(BinaryNode<Record> *&current) {
        rotate_right(current->right);
        rotate_left(current);
        return okay;
    }
};

#endif //TEST_RED_BLACK_TREE_H
