# Topological Sort of a Dag

## Explanation
The Topological Sort of a DAG (a directed acyclic graph) is one where each directed edge that can be labeled uv (the two vertices that are connected by this edge), the u comes before the v in any ordering made.
If we use the idea of installing dependencies in a project as an example, you can image that each of the nodes in a graph is a dependency that other nodes maybe rely on before being able to be installed.

Let's say we have a graph of 4 dependencies (I'm making up these dependencies on the fly, they might not work this way):
1. NodeJS
2. C++
3. Rust
4. ExpressJS

In order to install Express JS correctly, I need to install the other dependencies that ExpressJS has BEFORE Express can be installed.

Let's say that Express JS has the Dependency of NodeJS.  That can be represented by the following ascii art:
```
|-------|       |------|  
|       |       |      |
|Express| <---- |NodeJS|
|-------|       |------|
```

So in order for Express to be Installed, NodeJS must be installed first.

In this example, let's say that NodeJS requires two dependencies: Rust and C++.
In order to install NodeJS, Rust and C++ need to be installed as well.

More Ascii Art:
```
                |------|
             -  |      |
            -   | Rust |
|-------|  -    |------|  
|       | <
|NodeJS | <
|-------|  -    |------|
            -   |      |
             -  | C++  |
                |------|
```

In this world, both Rust and C++ must be installed to install NodeJS, which in turn is required to use ExpressJS.

A topological sort of the graph would return one of the following:
[Rust, C++, NodeJS, ExpressJS]

[C++, Rust, NodeJS, ExpressJS]

Which would return the correct order of dependencies that are required to install ExpressJS.


## Implementation
The algorithm I chose to implement Topological Sort in this instance is based off of DFS and can be found [here](https://en.wikipedia.org/wiki/Topological_sorting#Depth-first_search).
The reason I loop through the array of nodes and run visit on each node is because these abstract tasks are not directly connected.  If I just inputted one node, the DFS would only figure out which tasks are required (children) of that particular node.
With Topological Sort, we can feed an entire graph of disparate tasks that might not have direct relationships with each other and still ensure that they are processed in the correct order.

The input to the function is the graph represented as a map and the output is a slice of strings that are ordered from left to right. This order represents the tasks on the left being required to be completed before the tasks on the right. 

## Most Realest Application I can think of right now
I represented the graph as a JSON file that contains nodes as a HashMap in which each node has a name, a value, and the children that are dependent on this node. In most real life programming scenarios I have seen, the dependencies would not know the dependent libraries, rather the dependent libraries would know which
dependencies they have.  This would alter the relationships in my previous explanation by changing the meaning of each edge to be read as which depencies are REQUIRED to install the library.
I.E. NodeJS would know that it requires ExpressJS before it can install.  Rust and C++ know that they require NodeJS to install, etc. 

We can easily modify the algorithm to calculate this order by having the sorted nodes append to the end of the array in the visit function instead of the beginning.
```
...other logic

*sortedNodes = append([]string{n}, *sortedNodes...)
// will become
*sortedNodes = append(*sortedNodes, n)
```
In this example, the `children` field in the json file will really represent the idea of `dependencies`. 
You could also just reverse the sorted array after running the same logic, but it's more fun to figure out a way to do it as we parse through the graph.
