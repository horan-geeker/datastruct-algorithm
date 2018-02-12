#ifndef TEST_LIBRARY_H
#define TEST_LIBRARY_H

#include <iostream>
#include <vector>
using namespace std;

template<class Entry>
void process_tree_node(Entry &n) {
    cout << n;
}

#endif