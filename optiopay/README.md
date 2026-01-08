# Instructions, task

Bureaucr.at is a typical hierarchical organisation. Claire, its CEO, has a hierarchy of employees reporting to her and each employee can have a list of other employees reporting to him/her. An employee with at least one report is called a Manager.

Your task is to implement a corporate directory for Bureaucr.at with an interface to find the closest common Manager (i.e. farthest from the CEO) between two employees. You may assume that all employees eventually report up to the CEO.

Here are some guidelines:

- Resolve ambiguity with assumptions.
- The directory should be an in-memory structure.
- A Manager should link to Employees and not the other way around.
- We prefer that implementation is done in Go.
- How the program takes its input and produces its output is up to you.

# Solution

- recursively go down the tree
- count employees matching employee1 and employee2 in every subtree (including subtree top node)
- find the node with counter=2
- bingo!
