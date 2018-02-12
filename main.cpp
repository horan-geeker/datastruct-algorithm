#include "library.h"
#include "binary_tree.h"
#include "binary_search_tree.h"
#include "b-tree.h"
#include "red-black-tree.h"
#include "test.h"

using namespace std;

int main() {
//    BinaryTree<int> *binary_tree = new BinaryTree<int>;
//    binary_tree->inorder(process_tree_node<int>);
    BinarySearchTree<char> *bst = new BinarySearchTree<char>;
    bst->insert('a');
    bst->insert('b');
    bst->insert('c');
    bst->insert('d');
    bst->insert('e');
    bst->insert('r');
//    bst->remove(7);
//    bst->inorder(process_tree_node<char>);

    //solution();

    B_Tree<char, 5> *b_tree = new B_Tree<char, 5>;
    b_tree->insert('a');
    b_tree->insert('b');
    b_tree->insert('c');
    b_tree->insert('d');
    b_tree->insert('e');
    b_tree->insert('f');
    b_tree->insert('g');
    b_tree->insert('h');
    b_tree->remove('h');

    RB_Tree<char> *rb_tree = new RB_Tree<char>;
    rb_tree->insert('c');
    rb_tree->insert('o');
    rb_tree->insert('r');
    rb_tree->insert('n');
    rb_tree->insert('f');
    rb_tree->insert('l');
    rb_tree->insert('a');
    rb_tree->insert('k');
    rb_tree->insert('e');
    rb_tree->insert('s');

    auto *solution = new Solution;
    vector<int> a;
    a.push_back(0);
    a.push_back(1);
    a.push_back(2);
    a.push_back(3);
    a.push_back(4);
    a.push_back(5);
    for (int i : a) {
        cout<< i;
    }
    solution->sortedArrayToBST(a);
}