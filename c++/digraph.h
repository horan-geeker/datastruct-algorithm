//
// Created by hejunwei on 16/02/2018.
//

#ifndef TEST_DIGRAPH_H
#define TEST_DIGRAPH_H

#include <list>
#include <queue>

typedef int Vertex;

template<int graph_size>
class Digraph {
public:

//    void depth_sort(list <Vertex> &topological_order) {
//        bool visited[graph_size];
//        Vertex v;
//        for (int v = 0; v < this->count; ++v) visited[v] = false;
//        topological_order.clear();
//        for (int v = 0; v < this->count; ++v) {
//            if (!visited[v]) {
//                recursive_depth_sort(v, visited, topological_order);
//            }
//        }
//    }

//    void breadth_sort(list <Vertex> &topological_order) {
//        topological_order.clear();
//        Vertex v, w;
//        int predecessor_count[graph_size];
//        for (int v = 0; v < this->count; ++v) predecessor_count[v] = 0;
//        for (int v = 0; v < this->count; ++v) {
//            for (int i = 0; i < this->neighbors[v].size(); ++i) {
//                neighbors[v].retrieve(i, w);
//                predecessor_count[w]++;
//            }
//            queue<int> ready_to_process;
//            for (int v = 0; v < this->count; ++v) {
//                if (predecessor_count[v] == 0)
//                    ready_to_process.push(v);
//            }
//            while(!ready_to_process.empty()) {
//                ready_to_process.retrieve(v);
//                topological_order.insert(topological_order.size(), v);
//                for (int i = 0; i < this->neighbors[v].size(); ++i) {
//                    neighbors[v].retrieve(j, w);
//                    predecessor_count[w]--;
//                    if (predecessor_count[w] == 0)
//                        ready_to_process.push(w);
//                }
//                ready_to_process.serve();
//            }
//        }
//    }

private:
    int count;
    list <Vertex> neighbors[graph_size];

//    void recursive_depth_sort(Vertex v, bool visited[], list <Vertex> &topological_order) {
//        visited[v] = true;
//        int degree = this->neighbors[v].size();
//        for (int i = 0; i < degree; ++i) {
//            Vertex w;
//            neighbors[v].retrieve(i,w);
//            if (!visited[w]) {
//                recursive_depth_sort(w, visited, topological_order);
//            }
//        }
//        topological_order.insert(0, v);
//    }
};
#endif //TEST_DIGRAPH_H
