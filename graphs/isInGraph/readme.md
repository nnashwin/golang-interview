# A problem to determine if a string is located in a graph

## Problem Statement:

You have a Graph of strings which can be represented in any way.

The Nodes look like the following:

X C O O
X X X K
X S E I
X X X X

Write a function which takes the graph and a string as input and returns a boolean (true or false) value for whether or not the string can be found in the graph as output.

## Learning:

One of the biggest problems I had when asked this question had to deal with how to represent a graph.

Through some reading online and a bit of Khan Academy, I knew in the past that graphs can be represented in a number of ways, and I tried to represent it in a way I had heard before.

1. An Adjacency Matrix

An Adjacency Matrix is actually just a fancy word to describe a two-dimensional array.

Which looks like the following:
[
	[0, 1, 0, 1],
	[1, 0, 0, 0],
	[0, 0, 0, 0],
	[1, 0, 0, 0]
]

I originally tried to simply represent the graph in two arrays and put the strings as array indices, like in the following:

[
	[X, C, O, O],
	[X, X, X, K],
	[X, S, E, I],
	[X, X, X, X]
]

However, this was an incorrect understanding of how adjacency matrices work.

In adjacency matrices, the nested arrays themselves are the vertices in the graph, the number at any index is the weights of the edge connecting that particular vertice
to another, and the position in the nested array determines which vertices the edge connects together. 

So in the following example:

[
	[0, 1, 0, 1],
	[1, 0, 0, 0],
	[0, 0, 0, 0],
	[1, 0, 0, 0]
]

The first vertice, which is the first nested array, or arr[0], is connected to the second and 4th (arr[1] and arr[3] respectively) by edges which have a weight of 1.

These edges are also reflected in the 2nd vertice (arr[1]) and 4th vertice (arr[3]) by having a one in their 1st index (arr[1][0] and arr[3][0]).

2. Enter the linked list 
I later learned that I could also represent graphs as linked lists, and this method made a lot more sense to me.

Simple Linked lists are individual objects which have a value and a link (called a pointer in traditional cs and lower-level languages) to the next object of the same type.

Using a linked list, 

X C O O
X X X K
X S E I
X X X X

Can be represented 

X -> C -> O -> O -> X ... and so on.  (Notice the end of the first line O is jumping to the next row.)

Since in the problem statement it was mentioned that I can represent the graph anyway I want, representing it as a linked list is a standard way to solve the problem.

----- NEXT -----
