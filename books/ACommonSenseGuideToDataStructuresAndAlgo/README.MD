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

### Chapter 6 exercises

### Chapter 7 exercises

### Chapter 8 exercises
1.  [intersection](../../algo_data/array/intersection.go)
2.  [find_duplicated](../../algo_data/array/find_duplicated.go)
3.  [find_the_missing_letter](../../algo_data/array/find_the_missing_letter.go)
4.  [non_first_duplicated](../../algo_data/array/non_first_duplicated.go)
