# Load Balancer with Round-Robin Algorithm 
This Go program demonstrates a simple load balancer with a round-robin server aalgorithm. It creates virtual servers and simulates processing a large number of requests through the load balancer.

## Overview

The program consists of three main components:

1. **Load Balancer (`LoadBalancer`):** Manages the incoming requests and distributes them among virtual servers using a round-robin algorithm.

2. **Server (`Server`):** Represents a virtual server that performs work on incoming requests.

3. **Main Function (`main`):** Orchestrates the creation of the load balancer, virtual servers, and simulates the addition of requests to the load balancer. It then uses a round-robin algorithm to assign requests to servers.

## Components

### Load Balancer (`LoadBalancer`)

- `done` channel: Used to signal the termination of the load balancer.
- `mu` (sync.RWMutex): Provides thread-safe access to the request queue.
- `wg` (sync.WaitGroup): Ensures all requests are processed before program termination.
- `reqs` slice: Stores the IDs of added requests.

### Server (`Server`)

- `id`: Represents the unique identifier for each virtual server.

### Main Function (`main`)

1. Creates a load balancer.
2. Creates a specified number of virtual servers.
3. Simulates the addition of a large number of requests to the load balancer.
4. Implements a simple round-robin algorithm to distribute requests among servers.
5. Logs the results of server computations.

## Usage

To run the program, simply execute the following command:

```bash
go run main.go
```

The program will output logs indicating the creation of servers, addition of requests to the load balancer, and processing of requests by virtual servers. Feel free to explore and adapt the code for your specific use case. If you have any questions or suggestions, please don't hesitate to reach out.

