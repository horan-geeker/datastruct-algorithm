//
// Created by hejunwei on 05/02/2018.
//

#ifndef TEST_B_TREE_H
#define TEST_B_TREE_H

enum ErrorCode {
    not_present, success, overflow, duplicate_error, internal_error
};

template<class Record, int order>
struct B_Node {
    int count;
    Record data[order - 1];
    B_Node *branch[order];

    B_Node() {
        count = 0;
    };
};

template<class Record, int order>
class B_Tree {
public:
    B_Tree() { this->root = NULL; }

    ErrorCode search_tree(Record &target) { return this->recursive_search_tree(root, target); }

    ErrorCode insert(const Record &new_entry) {
        Record median;
        B_Node<Record, order> *right_branch, *new_root;
        ErrorCode result = this->push_down(this->root, new_entry, median, right_branch);
        if (result == overflow) {
            new_root = new B_Node<Record, order>;
            new_root->count = 1;
            new_root->data[0] = median;
            new_root->branch[0] = root;
            new_root->branch[1] = right_branch;
            root = new_root;
            result = success;
        }
        return result;
    }

    ErrorCode remove(const Record &target) {
        ErrorCode result;
        result = recursive_remove(this->root, target);
        if (this->root == NULL && root->count == 0) {
            B_Node<Record, order> *old_root = this->root;
            this->root = this->root->branch[0];
            delete old_root;
        }
        return result;
    }

private:
    B_Node<Record, order> *root;

    ErrorCode recursive_search_tree(B_Node<Record, order> *current, Record &target) {
        ErrorCode result = not_present;
        int position;
        if (current != NULL) {
            result = this->search_node(current, target, position);
            if (result == not_present) {
                result = this->recursive_search_tree(current->branch[position], target);
            } else {
                target = current->data[position];
            }
        }
        return result;
    }

    ErrorCode search_node(B_Node<Record, order> *current, const Record &target, int &position) {
        position = 0;
        while (position < current->count && target > current->data[position])
            position++;
        if (position < current->count && target == current->data[position])
            return success;
        else
            return not_present;
    }

    ErrorCode
    push_in(B_Node<Record, order> *current, const Record &entry, B_Node<Record, order> *right_branch, int position) {
        // shift all later data to the right
        for (int i = current->count; i > position; i--) {
            current->data[i] = current->data[i - 1];
            current->branch[i + 1] = current->branch[i];
        }
        current->data[position] = entry;
        current->branch[position + 1] = right_branch;
        current->count++;
    }

    ErrorCode push_down(
            B_Node<Record, order> *current,
            const Record &new_entry,
            Record &median,
            B_Node<Record, order> *&right_branch
    ) {
        ErrorCode result;
        int position;
        if (current == NULL) {
            median = new_entry;
            right_branch = NULL;
            result = overflow;
        } else {
            if (search_node(current, new_entry, position) == success) {
                result = duplicate_error;
            } else {
                Record extra_entry;
                B_Node<Record, order> *extra_branch;
                // 此处的递归被多次调用，同时修改着 extra_entry 的值，较深
                result = push_down(current->branch[position], new_entry, extra_entry, extra_branch);
                if (result == overflow) {
                    if (current->count < order - 1) {
                        result = success;
                        this->push_in(current, extra_entry, extra_branch, position);
                    } else {
                        this->split_node(current, extra_entry, extra_branch, position, right_branch, median);
                    }
                }
            }
        }
        return result;
    }

    void split_node(
            B_Node<Record, order> *current,
            const Record &extra_entry,
            B_Node<Record, order> *extra_branch,
            int position,
            B_Node<Record, order> *&right_half,
            Record &median
    ) {
        right_half = new B_Node<Record, order>;
        int mid = order / 2;
        if (position <= mid) {
            // case 1
            for (int i = mid; i < order - 1; ++i) {
                right_half->data[i - mid] = current->data[i];
                right_half->branch[i + 1 - mid] = current->branch[i + 1];
            }
            current->count = mid;
            right_half->count = order - 1 - mid;
            push_in(current, extra_entry, extra_branch, position);
        } else {
            // case 2
            mid++;
            for (int i = mid; i < order - 1; ++i) {
                right_half->data[i - mid] = current->data[i];
                right_half->branch[i + 1 - mid] = current->branch[i + 1];
            }
            current->count = mid;
            right_half->count = order - 1 - mid;
            push_in(right_half, extra_entry, extra_branch, position - mid);
        }
        median = current->data[current->count - 1];
        right_half->branch[0] = current->branch[current->count];
        current->count--;
    }

    ErrorCode recursive_remove(B_Node<Record, order> *current, const Record &target) {
        ErrorCode result;
        int position;
        if (current == NULL) result = not_present;
        else {
            if (search_node(current, target, position) == success) {
                // target is in the current node
                result = success;
                if (current->branch[position] != NULL) {
                    // not at a leaf node
                    copy_in_predecessor(current, position);
                    recursive_remove(current->branch[position], current->data[position]);
                } else remove_data(current, position);
            } else recursive_remove(current->branch[position], target); //not found continue recursive
            if (current->branch[position] != NULL)
                if (current->branch[position]->count < (order - 1) / 2)
                    restore(current, position);
        }
        return result;
    }

    void remove_data(B_Node<Record, order> *current, int position) {
        for (int i = position; i < current->count - 1; ++i) {
            current->data[i] = current->data[i + 1];
        }
        current->count--;
    }

    void copy_in_predecessor(B_Node<Record, order> *current, int position) {
        B_Node<Record, order> *leaf = current->branch[position];
        while (leaf->branch[leaf->count] != NULL)
            leaf = leaf->branch[leaf->count]; // move as far rightward as possible

        current->data[position] = leaf->data[leaf->count - 1];
    }

    void restore(B_Node<Record, order> *current, int position) {
        // right most branch
        if (position == current->count) {
            if (current->branch[position - 1]->count > (order - 1) / 2)
                move_right(current, position - 1);
            else
                combine(current, position);
        } else if (position == 0)
            if (current->branch[1]->count > (order - 1) / 2)
                move_left(current, 1);
            else
                combine(current, 1);
        else if (current->branch[position - 1]->count > (order - 1) / 2)
            move_right(current, position - 1);
        else if (current->branch[position + 1]->count > (order - 1) / 2)
            move_left(current, position + 1);
        else
            combine(current, position);
    }

    void move_left(B_Node<Record, order> *current, int position) {
        B_Node<Record, order> *left_branch = current->branch[position - 1];
        B_Node<Record, order> *right_branch = current->branch[position];

        left_branch->data[left_branch->count] = current->data[position - 1];
        left_branch->branch[++left_branch->count] = right_branch->branch[0];
        current->data[position - 1] = right_branch->data[0];
        right_branch->count--;
        for (int i = 0; i < right_branch->count; ++i) {
            right_branch->data[i] = right_branch->data[i + 1];
            right_branch->branch[i] = right_branch->branch[i + 1];
        }
        right_branch->branch[right_branch->count] = right_branch->branch[right_branch->count + 1];
    }

    void move_right(B_Node<Record, order> *current, int position) {
        B_Node<Record, order> *left_branch = current->branch[position];
        B_Node<Record, order> *right_branch = current->branch[position + 1];

        right_branch->branch[right_branch->count + 1] = right_branch->branch[right_branch->count];
        for (int i = right_branch->count; i > 0; --i) {
            right_branch->data[i] = right_branch->data[i - 1];
            right_branch->branch[i] = right_branch->branch[i - 1];
        }
        right_branch->count++;
        right_branch->data[0] = current->data[position]; // take entry from parent
        right_branch->branch[0] = left_branch->branch[left_branch->count--];
        current->data[position] = left_branch->data[left_branch->count];
    }

    void combine(B_Node<Record, order> *current, int position) {
        B_Node<Record, order> *left_branch = current->branch[position - 1];
        B_Node<Record, order> *right_branch = current->branch[position];

        left_branch->data[left_branch->count] = current->data[position - 1];
        left_branch->branch[++left_branch->count] = right_branch->branch[0];
        for (int i = 0; i < right_branch->count; ++i) {
            left_branch->data[left_branch->count] = right_branch->data[i];
            left_branch->branch[++left_branch->count] = right_branch->branch[i + 1];
        }
        current->count--;
        for (int i = position - 1; i < current->count; ++i) {
            current->data[i] = current->data[i + 1];
            current->branch[i + 1] = current->branch[i + 2];
        }
        delete right_branch;
    }

};

#endif //TEST_B_TREE_H
