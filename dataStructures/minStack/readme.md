# Min Stack

## Problem Definition
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

    push(x) -- Push element x onto stack.
    pop() -- Removes the element on top of the stack.
    top() -- Get the top element.
    getMin() -- Retrieve the minimum element in the stack. Example:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.

## Explanation
The actual stack we implement here is a housing struct of a linked list with nodes linked to each other through Prev field in the node.  We store a global min for readability sake, but we could just store the min
inside of the min at any particular time.  We can retrieve the current min either through querying the min at the current head (MinStack.Head.Min) or in our case from the global min (MinStack.Min).
Each time we push, we set the min in the node AND the global stack as the minimum of the val and the current min.
Each time we pop, we set the global min to what the node min is.

## Time Complexity
Since we are manipulating pointers, all operations happen in O(1) time.
