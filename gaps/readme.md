


Let p be a prime.
For any n-set D, we want to find an a such that ad \in E_p for all d in D.
le z be such that az = 1.
Then we want some z such that D is a subset of zE_p. 

If E_p is connected then zE_p satisfies the three gap condition and the higher gap conditions.
There are p such sets and they all satisfy the higher gap condition.



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
