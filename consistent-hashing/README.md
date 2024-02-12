# Consistent Hashing with Hash Ring

This Go program implements a simple consistent hashing algorithm using a hash ring. Consistent hashing is an algorithm often used in distributed systems to efficiently distribute data among a changing set of nodes. It is neither a service by itself nor it does the "actual" data transfer. It solves only the data ownership problem and not the data distribution problem. Here is "Why" consistent hashing is needed.

1. **Need for Even Data Distribution:** In large-scale distributed systems, data is distributed across multiple servers for horizontal scaling. To ensure predictable performance, it's crucial to evenly distribute data among these servers.

2. **Simple Hashing:** Initially, a common method for distributing data evenly is simple hashing. Each object's key is hashed using a hashing function, and then the modulo operation is performed on the hash result to determine which server the object belongs to.

3. **Drawbacks of Simple Hashing:** While simple hashing works well when the cluster size is fixed and data distribution is even, it becomes problematic when servers are added or removed. In such cases, most of the keys need to be redistributed, leading to performance issues.

4. **Introduction to Consistent Hashing:** Consistent hashing mitigates the issues of traditional hashing by ensuring that most objects remain assigned to the same server even as the number of servers changes.

5. **Core Concept of Consistent Hashing:**
   - Both servers and objects are hashed using the same hashing function, mapping them to the same range of values.
   - A hash ring is constructed by connecting both ends of the hash space, forming a ring structure.
   - Servers are placed on the hash ring, and objects are mapped onto the ring using their hashes.
   - To locate the server for a particular object, one traverses clockwise from the object's position on the ring until a server is found.

6. **Benefits of Consistent Hashing:**
   - Adding or removing servers requires redistributing only a fraction of the keys, minimizing data movement.
   - Ensures predictable performance even in dynamic environments where servers come and go frequently.

7. **Addressing Uneven Data Distribution:** While consistent hashing ensures even distribution in theory, in practice, segments of the ring may still be unevenly distributed. Virtual nodes are introduced to address this issue, where each server appears at multiple locations on the ring, representing virtual nodes.

8. **Real-world Applications:** Consistent hashing is used in various distributed systems, including NoSQL databases like Amazon DynamoDB and Apache Cassandra, content delivery networks like Akamai, and load balancers like Google Load Balancer.

Consistent hashing is a fundamental technique in distributed systems design, enabling scalability and fault tolerance while maintaining performance and efficiency.


## Implementation

### HashRing Struct

The `HashRing` struct contains the following attributes:

- `mu`: A read-write mutex to ensure thread safety.
- `nodes`: A sorted slice of node IDs representing the hash ring.
- `hashMap`: A map that associates node IDs with their corresponding node names.

### Functions

#### `newHashRing() *HashRing`

Creates and returns a new `HashRing` with initialized attributes.

#### `hashKey(key string) int`

Hashes a given key using the FNV-1a algorithm and returns the resulting hash as an integer.

#### `(hr *HashRing) AddNodeToRing(node string)`

Adds a new node to the hash ring. It computes the hash of the node and ensures it doesn't already exist. Nodes are added to both the `nodes` slice and the `hashMap`.

#### `(hr *HashRing) DeleteNodeFromRing(node string)`

Deletes a node from the hash ring. It finds the node in the `nodes` slice, removes it, and updates the `hashMap` accordingly.

#### `(hr *HashRing) GetNodeForKey(key string) int`

Finds the node responsible for a given key by hashing the key and searching for the corresponding node ID in the `nodes` slice.
This code defines the `GetNodeForKey` method for the `HashRing` struct. The purpose of this method is to determine which node in the hash ring is responsible for a given key using consistent hashing.

Here's a breakdown of the function:

1. `hr.mu.RLock()`: It acquires a read lock on the hash ring's mutex, ensuring that no other goroutine can modify the hash ring while this operation is in progress.

2. `defer hr.mu.RUnlock()`: This ensures that the read lock is released when the function completes, regardless of whether it returns normally or with an error.

3. `keyID := hashKey(key)`: It calculates the hash of the input key using the `hashKey` function, which employs the FNV-1a algorithm.

4. `index := sort.Search(len(hr.nodes), func(i int) bool {...})`: It performs a binary search on the sorted `hr.nodes` slice to find the index where the key's hash value would fit in the sorted order. The provided anonymous function returns `true` when the desired position is found.

5. If the index is equal to the length of `hr.nodes`, it means the key's hash value is greater than or equal to all nodes' hash values in the ring. In this case, the search "circles back" to the beginning of the ring by setting the index to 0.

6. `return hr.nodes[index]`: Finally, the function returns the node ID associated with the determined index in the `hr.nodes` slice. This represents the node responsible for the given key in the hash ring.

In summary, the `GetNodeForKey` method uses consistent hashing principles to find the node responsible for a specific key in the hash ring, providing a balanced distribution of keys among the available nodes.

### Main Function

The `main` function demonstrates the usage of the hash ring. It:

1. Creates a new hash ring.
2. Adds nodes ("nodeA," "nodeB," and "nodeC") to the hash ring.
3. Retrieves the node responsible for a test key.
4. Deletes a node ("nodeA") from the hash ring.
5. Retrieves the node responsible for the test key after deletion.

### Usage

To run the program, simply execute the following command:

```bash
go run main.go
```

### Note

- The program includes log statements to provide insights into the operations performed by the hash ring.

Feel free to modify the code and experiment with different nodes and keys to observe the consistent hashing behavior.




