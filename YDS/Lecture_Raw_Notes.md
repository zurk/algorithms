# Yandex Data School. Algorithms and data structures. Part 1.
https://yandexdataschool.com/edu-process/algorithms1

## Lecture 1. Complexity and computation models. Amortized costs.

When we try to build good algorithm we think about:
1. Time
2. Memory

Time and memory can be measured differently. Depend from the model we use.
We will use Word RAM Model. RAM is Random access memory. 
You can imagine it as a huge array (RAM[i] is one word with length of w (32/64) bits). 
Also we have CPU and it has list of elementary operations: load from RAM[i], store to RAM[i], arithmetic operations, etc.

**Example.** Count ones in binary number x. 
Complexity: O(w)

**Example.** Sort(a_1, a_n). 
Complexity? Worst case complexity. f(n) = \max_{input \in Input_n} f(input)
Complexity: O(n \log n)
for RAM model you have faster but useless methods.  (????) But too big constant. O(n \log \log (n) ) 
See also: Radix Sort O(n \cdot w).

Avarage case complexity: f(n) = \avg_{input \in Input_n} f(input). You need something like prior probability for input and calculate Expectation. Useless (?). Better to use randomized complexity: No need to use prior probability, but you should use randomized algothims. You need pure random operation in CPU.

f(n) = \max_{input} \avg_r {f(r, input)}, where r -- number of random calls. It is better than avarage case complexity because you can rerun algorithm for the same input.

## About sort complexity. Decision trees.
We need one more model: Decision trees. It is useful for lower bound estimation.

**Example**. Sort(a_1, \dots, a_n) \rightarrow S_n (permutation).
RAM model is phisical. We have CPU and RAM. And uniform. No difference for input size. It is not true for decision trees models.
You need to define A_1, A_2, \dots, A_n, \dots algorithms. A_i is tree to sort i items with 2 Nodes type: indernal nodes with comparation and leafs with answer from S_n (permutation set). When you "run" your tree, you just follow the steps for particular tree A_i. So, for this model

f(n) = \Omega(n \log n) (\Exsist \alpha > 0 : f(n) \ge \alpha n \log n)

Proof:
\max in f(n) is equal to height of the tree, but you should have an input to reach this leaf.
You can say that if Height(A_n) \le \beta than \max_{input} f(input) \le \beta.

\forall \pi \in S_n you should find leaf in A_n for permutation \pi. So, at least you should have n! leaves. Let's run algothim for all permutations and choose the deapest leaf from all reached leafs. So Height(A_n) at least \log(n!) = \Omega (x \log x).

So. what about RAM model?
If your program satisfies some conditions it corresponds to Decision tree model. 
Conditions:
1. You can only use non-random algorithms and can only compare two values in array. 
Easy to understand that it leads to decision tree model.

## Lecture 2. Amortized (redused) costs
It is more related to data structures.

**Example.** std::vector<T>
To extend vector size (push_back) your worst case cost in equal to vector size O(n) because you need to copy all vector.
But you need it in rear cases. 
So, Amortized cost for push_back can be considered as O(1) for multiplicative capasity updates with \lambda multiplier. Why?
**Bank method:**
We have several push_back pb_1, \cdot, pb_n and it will cost O(n).
And we are paid \alpha n $ in total. \alpha for each push_back.

Consider \alpha = 3, \lambda=2. Easy case: we have anough vector capasity.

Easy to understand that 
0.5 * cap \le size \le cap

It is important to know how many money you have on each step. It is called potential function.
We want to store for 2$ in each cell after size index in vector.
Ok, we have new push_back with 3$. if we have anough capasity we put element for 1$ and put 2$ more to next cell without money. 
In case of size = capasity we spend our money to copy array to new place. That's it. 1$ is payment for any constant operation.

**Method of potentials:**
The same, but formal:
s - states. a - actions
s_0 ---a_0---> s_1 ---a_1---> ... ---a_{n-1}---> s_n
c(a_i) is true action a_i cost.
Consider \phi: S \rigtharrow \mathbb{R} -- we need to find the best one.

Next, consider redused cost: \hat c (a_i) = c(a_i) + \phi(s_i) - \phi(s_{i+1}).
Now, if you sum \hat c(a_o) + \dots + \hat a_{n-1} = c(a_0) + \dots + c(a_{n-1}) + [\phi(s_0) - \phi(s_n)].
So, c(a_o) + \dots + c(a_{n-1}) = \hat c(a_0) + \dots + \hat a_{n-1}  + [\phi(s_{i+1} - \phi(s_{0}]
You need to find \phi such that \hat c(a_i) \le T and [\phi(s_n) - \phi(s_0)] \le 0.
The fact leads to c(a_o) + \dots + c(a_{n-1}) \le T n
T can be not constant in general case, for example \log n.

For bank case we need to take minus money in the bank as \phi.

** Example: **
How to make queue from stacks only?
to build queue you need to stacks. one for tail (S_T) and one for head (S_H).

tail(enqueue) |_| |_| |_| |_| |_| |_| |_| |_| |_| |_| head(dequeue)

                      S_T <----- | -----> S_H

1. Only enqueue (2$ for each):
              |$| |$| |$| |$| |$||
2. If we do dequeue we need to relocate everything to S_H (1$ for copy-paste, 1$ for dequeue):
                                 ||_| |_| |_| |_| |_|           

So, if we want to estimate complexity we have the same problem as for std::vector<T>.

... Skip explanation ...

This example can be useful.
How to find max in stack in a fast way?
You need second stack (because of pop action)

|a_1| |a_2| |a_3| |a_4| |a_5|
|m_1| |m_2| |m_3| |m_4| |m_5| m5 = \max(a_5, m_4)

It is not the same for queue. But you can model queue with 2 stacks! Will be the same constant redused costs. 

Deamortizetion. You need to rebuild data structure to go from amortized costs to worst case.
Key: you perform your hard action before the case, so you need to do it in paralell way with easy steps.
It is important for real-time systems.

## Mutability

You have data structure and actions of two types: mutators (change structure state) and accessors (just get something from structure).
There are exsists immutable data structures.

**Example:** Immutable stack (IS).
Each mutator action returns new stack. 

Immutable Stack   (IS)  | Mutable stack (MS) 
------------------------|------------------------
1. bool isEmpty() const | 1. bool isEmpty() const
2. IS push(T)           | 2. void push(T)
3. <T, IS> pop()        | 3. T pop()

for exapmle in functional languages you just do not have state conseption and you each time have new element.
To build efficient immutable data sctructure you should use references.
Immutable leads to persistance. It can be useful to build efficent algorithms. 
Easy to implement Immutable Stack, not so easy for queue.
We can build queue from two stacks, so it is possible. But computationaly uneffisient. You need to Deamortize it, and will have ~5 stacks for it O_o.

## Quick sort.
Stable property: if a_i = a_j for i < j => \pi^{-1}(a_i) < \pi^{-1}(a_j)
a_1, \dots, a_n => \forall l \in L, r \in R such that l < r
you need to choose pivot = a_i first. Better to use randomized approach and choose i randomly. 



