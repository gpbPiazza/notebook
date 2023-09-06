# Mastering Multithreading Programming with Go (Golang)


By James Cutajar. Course [Link](https://www.udemy.com/course/multithreading-in-go-lang/)

## Parallel computing
- [Amdahl's law](https://en.wikipedia.org/wiki/Amdahl%27s_law)
  - Amdahl's law is often used in parallel computing to predict the theoretical speedup when using multiple processors.
  - The Amdahl's law defines the optimization of a part of a process is limited by the fraction of time that the improved part is actually used.
  - The limiting the speed-up because of the sequential parts of our software
  ![Alt text](image-1.png)
  Definitions:
  - Where `s` is the quantity of time of speedup, how much faster the process will be if we add more processors to the process.
  - The execution time of the whole task before the improvement of the resources of the system is denoted as `T``. 
  - `p` is the fraction of the execution time of the task that would benefit from the improvement of the resources.
  p = TaskB/(TaskA + TaskB)
  ![Alt text](image.png)
  http://www.johngustafson.net/pubs/pub13/amdahl.htm
- [Gustafson's law](https://en.wikipedia.org/wiki/Gustafson%27s_law)
  - Gustafson's law addresses the shortcomings of Amdahl's law, which is based on the assumption of a fixed problem size, that is of an execution workload that does not change with respect to the improvement of the resources. Gustafson's law instead proposes that programmers tend to increase the size of problems to fully exploit the computing power that becomes available as the resources improve.
  ![Alt text](image-2.png)
  - Don't matter the sequential parts of our task if we increase the number of tasks/problems per processor because the relation is always linear. As long, you have an infinite problem size, and you increase the number of processors it doesn't really matter how much of your software is linear, how many sequential parts he has. But your speedup still depends on many parts of your taks is linear or parallel. More parts can be parallel, more speedup.