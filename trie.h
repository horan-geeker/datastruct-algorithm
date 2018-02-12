//
// Created by hejunwei on 05/02/2018.
//

#ifndef TEST_TRIE_H
#define TEST_TRIE_H

const int num_chars = 28;

struct TrieNode {
    char *data;
    TrieNode *branch[num_chars];

    TrieNode();
};

class Trie {
public:

private:
    TrieNode *root;
};

#endif //TEST_TRIE_H
