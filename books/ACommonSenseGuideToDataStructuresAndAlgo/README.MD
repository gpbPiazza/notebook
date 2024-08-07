# A Common-Sense Guide to Data Structures and Algorithms
##### Level Up Your Core Programming Skills

By Jay Wengrow

### Chapter 1 exercises

1 - For an array with 100 elements

1.a - To Read from an array always take one step
1.b - To Search for a none existing value in an array would take 100 steps, since we need to iterate in each value of the array
1.c - To insert at the beginning would take 101 steps, since the computer must shift to the right all elements in the memory and insert the new element at the beginning.
1.d - To insert at the end of array would take one step, since the computer always knows the start and the end of array.
1.e - To delete at the beginning would take 100 steps, 1 to delete the element and 99 to shift all elements to the left.
1.f - To delete at the end would take one step, since the computer knows the end of the array and don't need to shift elements to this operation.

2 - For an array-based set with 100 elements
2.a - To Read as the same of the array would take one step.
2.b - To Search, the same quantity of steps of the array.
2.c - To Insert a new value at the beginning we must provide a full search in the set to verify if the given value is no duplicated, so this operation would take 100 steps, and we need to shift all the elements to the right plus this operation would take more 100 steps and to insert the new value 1 step. So the all operation would take 201 steps.
2.d - To Insert a new value at the end we must provide a full search for the same reason at letter C, since we want insert at the end we don't need to shift elements sou this operation would take 100 steps from the search and 1 step to insert totalizing 101.
2.e - To Delete at the beginning would take 1 step to delete and 99 steps to shift all elements to the left, totalizing 100 steps.
2.f - To delete at the end would take one step.

3. It would take N steps to count the frequency of the given value in an array. Where N is the size of the given array.

### Chapter 2 exercises

1. Since a linear search has to iterate into each element in the array to find the number eight would take 4 steps.
2. A binary search would take one step, because the linear search always start verifying at the middle of the array and the number 8 is in middle position.
3. The number of steps for a binary search in array of 100,000 would be 16 or 17, we can verify this by dividing 100,000 by 2 until reach the zero value of the vision, the number of division is the number of steps.

### Chapter 3 exercises

1. Is O of one, in constant time. The number of steps of this algorithm doesn't change by the given data.
2. Is O of N, the algorithm iterates in all elements of the array and sum their values, if the given array doubled his size the number of steps will too.
3. Is O of Log2(N), since the quantity of inputted data doubles and the number of steps only increase by one.
4. Is O of 1, all operations in the algorithm are in constant time, so if the arrays double of size and also in the worst case we always will have the same quantity of steps.

### Chapter 4 exercises

1.
| N Elements | O(N) | O(log2 N) | O(Nˆ2)   |
| -------- | ------ | --------- | -------- |
| 100      | 100    | 7         | 10000    |
| 2000     | 2000   | 11        | 4000000  |

2. Since the algorithm has O(Nˆ2) notation and takes 256 steps to conclude the operation the size of array can be found as root square of 256 that is 16 array size.
3. The algorithm show iterates in the array one time and inside the first foor loop do another iteration having this implementation will result in O(Nˆ2) time complexity. Because if the array size grows the number of steps will grows quadratically.
4. 
```python
  def greatestNumber(array):
    greatestNum = array[0]
    for i in array:
      if greatestNum < array[i]:
        greatestNum = array[i]

    return greatestNum
```
### Chapter 5 exercises
1. An Algorithm that takes 4 N + 16 steps would have 4*O(n)+16*O(1) time complexity, since Big O ignores constants we would have O(n).
2. An algorithm that takes 2 Nˆ2 steps would be 2*O(nˆ2) time complexity, since Big O ignores constants we would have O(nˆ2).
3. O(n)
4. O(n)
5. $ O(n)*O(nˆ(1/2)) = O(nˆ3/2)) $ simplifying we would have O(nˆ2)
### Chapter 6 exercises
1. An algorithm that takes 3Nˆ2+2N+1 steps would have time complexity O(nˆ2) since we can ignore constants and Big O only cares about the worst case notation.
2. For N+logN steps would be O(n) same response from question 1.
3. Worst case: O(nˆ2), medium case: O(nˆ2 /2) and best case: O(1).
4. Insert the return statement in side of the if condition when found X to return immediately and stop iterating.

### Chapter 7 exercises
1. BestCase O(1), Medium O(N/4), Worst O(N/2) → O(N)
2. O(N) since N is the total array length of both arrays.
3. B: O(1), Medium O((N*M)/2), Worst O(N*M).
4. O(nˆ3)
5. O(log n)
### Chapter 8 exercises
1.  [intersection](../../algo_data/array/intersection.go)
2.  [find_duplicated](../../algo_data/array/find_duplicated.go)
3.  [find_the_missing_letter](../../algo_data/array/find_the_missing_letter.go)
4.  [non_first_duplicated](../../algo_data/array/non_first_duplicated.go)
### Chapter 9 exercises

### Chapter 10 exercises
1. Base case is the if low>high puts low
2. Would do factorial of 10 to infinite, the base case would never be reached with (n - 2) in the recursion call, to fix we should change -2 to -1.
3. Base case would be if low == high return low
4. 
```go
func PrintNums(nums []a slice of int_OR_[]int) {
  for i := range nums {
    if isSlice(nums[i]) {
      PrintNums(nums)
    }else {
      fmt.Println(nums[i])
    }
  }
}
```
### Chapter 11 exercises

1. [count_all_chars](../../algo_data/array/count_all_chars.go)
2. [only_even_numbers](../../algo_data/array/only_even.go)
3. [get_triangular_number_in_given_position](../../algo_data/sequences/triangular_numbers.go)
4. [index_of_x](../../algo_data/array/index_of_x.go)
5. [unique_paths](../../algo_data/matrix/unique_paths.go)

### Chapter 12 exercises
1. [add_until_100](../../algo_data/array/add_until_100.go)
2. [golomb_with_memo](../../algo_data/sequences/golomb.go)
3. [uniquePaths_with_memo](../../algo_data/matrix/unique_paths.go)

### Chapter 13 exercises

1. [greatest_product_treee_numbers](../../algo_data/sorting/greatest_product_tree_number.go)
2. [find_missing_number](../../algo_data/sorting/find_greatest_number.go)
3. [find_greatest_number](../../algo_data/sorting/find_greatest_number.go)

### Chapter 14 exercises

1. [print_list_method_Print](../../algo_data/linkedlist/list.go)
2. [rever_print_double_linked_List](../../algo_data/linkedlist/double_linked_list.go)
3. [last_node_from_linked_list](../../algo_data/linkedlist/list.go)
4. [reverse_linked_list](../../algo_data/linkedlist/list.go)
5. [delete_self_node_linked_List](../../algo_data/linkedlist/list.go)

### Chapter 15 exercises

1. [1,5,9,2,4,10,6,3,8]
    (1)
      \
      (5)
   /        \
  (2)       (9)
    \       /   \
    (4)    (6)     (10)  
    /         \
  (3)         (8)
2. 1000 well balanced nodes, the number of steps to find a element would be 10 steps. Because the time complexity to search
in a binary search tree in the bestcase if O(Log(n)) wich means by each iteration we cut in the half the quantity of elements to search, so:
1000/2 = 500
500/2 = 250
250/2 = 125
125/2 = 62,5
62,5/2 = 31,25
31,25/2 = 15.625
15.625/2 = 7.8
7.8/2 = 3.9
3.9/2 = 1.9
1.9/2 = 0.97
3. [find_greatest_val](../../algo_data/trees/binary_search_tree/find_greatest.go)
4. Show sequence of prints of the following function
```go
  func TrasnverseAndPrint(node *Node) {
    if node == nil {
      return
    }
    fmt.Println(node.Val)
    TrasnverseAndPrint(node.Left)
    TrasnverseAndPrint(node.Right)
  }
```
- Moby dick
- Great Expectations
- Alice in Wonderland
- Lord of the files
- Robinson Crusoe
- Pride and Pejudice
- The odyssey

5. Show sequence of prints of the following function
```go
  func TrasnverseAndPrint(node *Node) {
    if node == nil {
      return
    }
    TrasnverseAndPrint(node.Left)
    TrasnverseAndPrint(node.Right)
    fmt.Println(node.Val)
  }
```
- Alice in Wonderland 
- Lord of the files
- Great Expectations
- Pride and Pejudice
- The odyssey
- Robinson Crusoe
- Moby dick


### Chapter 16 exercises

1. 	
  		    11
			 /    \
		  10	    8
    /   \    /  \
    6    9  7    4
  / \   / \
 2   1  3  5

2. 
  		   9
			 /    \
		  6	      8
    /   \    /  \
   3    5  7    4
  / \   
 2   1  

3. a heap with [55,22,34,10,2,99,68]
     99
   /   \ 
  22   68
 / \   / \
10 2  34 55

Delete order is: [99,68,55,34,22,10,2]
By doing this we juset dsicovered a new sort algorithm, heapsort.
The Main ideia is Insert the elements into the heap and delete them from the heap inserting into a array.
Takes O (N log N)

### Chapter 17 exercises

1. 
- "tag"
- "tan"
- "tank"
- "tap"
- "today"
- "total"
- "we"
- "well"
- "went"

2. [TestTrie_CHAPTER17_EXERCISE2](../../algo_data/tries/trie_test.go)
3. [TransversePrintln](../../algo_data/tries/trie.go)
4. [AutocorrectWord](../../algo_data/tries/trie.go)

### Chapter 18 exercises

1. if the user search nail in this graph probably the function of recommendation would return 
"nail polish", "needles", "pins" and "hammer".
2. A, B, E, J, F, O, C, G, K, D, H, L, M, I, N, P
3. A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P 
4. [bfsSearch](../../algo_data/graphs/all_connected_graph.go)
5. [shortestPath](../../algo_data/graphs/all_connected_graph.go)

### Chapter 19 exercises

1. The wordBuilder code has space complexity of O(nˆ2), because the units of memory grow as the input data grow and we have a nasted foor loop inserting N elements in the new array.
2. The reverseArray code has space complexity of O(n), because the units of memory grow as the input data grow, 
this code only has O(n) space complexity because the code implementation creates a new array to implement the algorithm.
3. [reverse](../../algo_data/array/reverse.go)
4.
| Version | Time Complexity | Space Complexity |
|---------|-----------------|------------------|
| v1.0    | O(n)            | O(n)             |
| v2.0    | O(n)            | O(1)             |
| v3.0    | O(n)            | O(n)             |


### Chapter 20 exercises

1.[find_who_plays_in_both_sports](../../algo_data/array/find_common_players.go)
2.[find_missing_int](../../algo_data/array/find_missing_int.go)
3.[find_monst_profit](../../algo_data/array/find_the_most_proftiable_sell.go)
4.[find_highest_product](../../algo_data/array/find_highest_product.go)
5. 