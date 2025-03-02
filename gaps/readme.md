
The problem is simple - we have a procedure that takes in a node and generates
left/right nodes if it exists. Hence, we are able to generate a tree.

Requirements:

A = { S in L | D is disjoint from S }
B = { S in L | D is a subset of S }
There is a bijection A -> B taking a set to its complement.
Hence |A| = |B|.
X = { S in L | D is not a subset of S }
Then |X| = |L| - |B| = |L| - |A|.
A is a subset of X
|A| <= |X| = |L| - |A|
|A| <= |L|/2


An upper bound on |X| <=> a lower bound on |A|
k <= |A| => L - k >= L - |A| = |X|
We want a really high lower bound on |A|.
|A| <= |L|/2

|X| = |L|-|A|>=|L|/2
