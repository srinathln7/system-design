# Everything you need to know about REST**

## Introduction:
- REST stands for Representational State Transfer.
- REST is a specification for communication between systems over HTTP.
- It is not a protocol, library, or framework.

## Key Concepts:
1. **Resource-Centric Architecture:**
   - In REST, everything is considered a resource.
   - Resources are entities in the system, e.g., a student, book, seller, etc.
   - Each resource has a unique identifier (URL).

2. **Client-Server Model:**
   - Standard client-server architecture where the client sends requests, and the server responds.
   - REST focuses on how clients and servers should communicate.

3. **Representation of Entities:**
   - REST emphasizes the representation of entities.
   - Commonly seen representations are JSON or XML.
   - Client demands a specific representation of the entity.

4. **Flexibility in Data Storage:**
   - REST doesn't enforce how data is stored in the server's database.
   - The server is free to choose its storage method, and it remains independent of REST.

5. **HTTP Methods (Verbs):**
   - REST leverages HTTP methods (GET, PUT, POST, DELETE) for actions on resources.
   - URL + HTTP method specifies the action and resource.
   - Multiplexing actions on the same resource with different methods.

6. **REST Over HTTP:**
   - Although REST doesn't mandate HTTP, it is commonly implemented over HTTP due to its widespread adoption.
   - HTTP methods align well with REST principles.

## Advantages of REST over HTTP:
1. **Existing Tooling:**
   - Widely adopted internet infrastructure.
   - Utilizes existing tools like browsers, clients (curl, Postman), web caches, monitoring tools, load balancers, and security controls.

2. **Standardization:**
   - HTTP methods and protocols are well-standardized.
   - Simplifies communication between clients and servers.

3. **Load Balancing:**
   - Leverages HTTP-based load balancers for uniform load distribution.

4. **Security:**
   - Built-in security features (SSL) for secure communication.
   - Existing security controls are applicable.

## Downsides of REST over HTTP:
1. **Consumption Challenges:**
   - Serialization and deserialization of data required.
   - Lack of native language support makes consumption more challenging.

2. **Repetitive Consumption Logic:**
   - Repetition of logic for handling serialization, deserialization, failures, timeouts, etc.
   - Individual consumers need to implement these details.

3. **Limited HTTP Verb Support:**
   - Some web servers may not support all HTTP verbs.
   - Choosing a server that supports all required verbs is crucial.

4. **Large HTTP Payloads:**
   - JSON, commonly used in REST, may lead to large payloads.
   - Not suitable for low-latency requirements.

5. **Protocol Switching Limitation:**
   - REST over HTTP doesn't allow switching to other protocols like UDP for specific performance requirements.
   - Limited flexibility in changing underlying protocols.

## Conclusion:
- REST is widely adopted but has some limitations.
- Choose REST based on specific use cases and requirements.
- Be aware of downsides while adopting REST over HTTP.

## Recommendation:
- Consider RPC (Remote Procedure Calls) for ultra-low-latency requirements.
- Choose the right specification based on the specific needs of the application.

## Remarks


**HTTP Methods (Verbs) in REST:**

REST relies on standard HTTP methods (also known as verbs) to perform actions on resources. The common HTTP methods are:
- **GET**: Retrieve information about the resource.
- **PUT**: Update or create a resource.
- **POST**: Create a new resource.
- **DELETE**: Remove a resource.

**Example: Managing a Student Resource**

Imagine you have a web application that manages student information. Each student is considered a resource, and you want to perform various actions on this resource.

1. **GET Method: Retrieve Student Information**
   - URL: `https://api.example.com/students/123`
   - Action: Retrieve information about the student with ID 123.
   - HTTP Method: `GET`

2. **PUT Method: Update Student Information**
   - URL: `https://api.example.com/students/123`
   - Action: Update the information of the student with ID 123.
   - HTTP Method: `PUT`

3. **POST Method: Create a New Student**
   - URL: `https://api.example.com/students`
   - Action: Create a new student with the provided information.
   - HTTP Method: `POST`

4. **DELETE Method: Remove a Student**
   - URL: `https://api.example.com/students/123`
   - Action: Delete the student with ID 123.
   - HTTP Method: `DELETE`

**URL + HTTP Method Specifies Action and Resource:**

In each example, the URL represents the resource (e.g., `/students/123`), and the HTTP method specifies the action to be taken on that resource (e.g., `GET`, `PUT`, `POST`, or `DELETE`).

- `GET` is used when you want to retrieve information.
- `PUT` is used when you want to update or create a resource.
- `POST` is used when you want to create a new resource.
- `DELETE` is used when you want to remove a resource.

**Multiplexing Actions on the Same Resource:**

REST allows for multiplexing actions on the same resource using different HTTP methods. For instance, the URL `/students/123` can be used for multiple actions:
- `GET /students/123`: Retrieve information about the student.
- `PUT /students/123`: Update the information of the student.
- `DELETE /students/123`: Remove the student.

By changing the HTTP method, you perform different actions on the same resource, and this flexibility is a key characteristic of RESTful design. It allows developers to express various operations on resources using a simple and consistent approach.