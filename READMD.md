# test_1  
This program finds the maximum path sum in a given triangle of numbers using Dynamic Programming (DP). The approach follows the Bottom-Up DP Method to efficiently compute the result.  

### Steps:  
1. Start from the second-last row:  
   - Process the triangle from bottom to top.  

2. Update each value based on possible paths:  
   - For each element, update its value by adding the maximum of its two downward neighbors.  
   - This ensures that each element stores the best possible sum from that point downward.  

3. Repeat until reaching the top row:  
   - The first element (top of the triangle) will contain the maximum path sum.  

# test_2  
This solution applies Iterative Constraint Propagation, a method similar to Relaxation Algorithms (e.g., Bellman-Ford). The goal is to iteratively update values based on local constraints (`R`, `L`, `=`) until the sequence stabilizes.  

### Steps:  
1. Forward Pass:  
   - Process the string from left to right.  
   - If `R` is found, ensure the right value is at least one greater than the left.  
   - If `=`, enforce equality between adjacent values.  

2. Backward Pass:  
   - Process the string from right to left.  
   - If `L` is found, ensure the left value is at least one greater than the right.  
   - If `=`, enforce equality again.  

3. Iterative Correction:  
   - Repeat the process until no further changes occur, ensuring all constraints are satisfied.  

4. Normalization:  
   - Adjust the smallest value in the sequence to zero, preserving relative differences.  

# test_3  
This program fetches, processes, and counts words from the Bacon Ipsum API using a concurrent worker-based approach for performance optimization.  

### Steps:  
1. Fetching Data (`FetchBaconIpsum`)  
   - Makes an HTTP GET request to the Bacon Ipsum API.  
   - Uses context with timeout (5s) to prevent long-running requests.  
   - Reads the response and returns it as a string.  

2. Word Splitting (`splitWords`)  
   - Uses `strings.FieldsFunc()` for efficient word splitting.  
   - Removes non-letter characters while preserving words like `"t-bone"`.  

3. Counting Words (`ExtractMeatWords`)  
   - Converts words to lowercase to avoid case mismatches.  
   - Uses `sync.Map` (thread-safe) to store word counts.  
   - Spawns multiple worker Goroutines (2 × CPU cores) for parallel processing.  
   - Uses atomic operations to update counts efficiently without locks.  

4. Concurrent Processing (Worker Pool)  
   - A buffered channel (`wordChan`) queues words before sending them to workers.  
   - Each worker:  
     - Reads a word from the channel.  
     - Checks if the word exists in `sync.Map`.  
     - Uses `atomic.AddInt32` to update the count safely.  
   - The main Goroutine waits (`wg.Wait()`) until all workers finish.  

### Performance Optimization Techniques:  
1. Parallel Processing with Goroutines → Uses multiple workers (2 × NumCPU) to speed up processing.  
2. Atomic Operations (`sync/atomic`) → Eliminates the need for locks, making updates faster and non-blocking.  
3. Buffered Channels → Improves efficiency by reducing Goroutine blocking.  
4. Optimized Word Splitting → Uses `strings.FieldsFunc()` instead of regex for better speed and memory usage.