#  Introduction to RPC - Remote Procedure Calls

## Key Concepts:
1. **RPC Overview:**
   - RPC stands for Remote Procedure Calls, which are widely adopted for inter-service communication over a network.
   - The core highlight of RPCs is to make a network call look like a local function call.
   - Abstraction of complexities like serialization, decentralization, and transport is a key feature of RPC.

2. **RPC Components and Communication:**
   - Services in RPC can communicate like a client-server model over various communication protocols such as TCP, UDP, etc.
   - Traditional communication involves a client making a request to the server, and the server responds, typically over protocols like HTTP or others.

3. **Use Case:**
   - Example scenario: Authentication service needs to communicate with the Notification service.
   - Traditional REST approach involves making HTTP calls, leading to repetitive code and potential language-specific issues.
   - RPC provides a solution by abstracting these communication complexities.

4. **RPC Code Example:**
   - RPC allows writing code that looks like a local function call but hides the complexity of remote network calls.
   - Example: `def login(email, password):` with an internal call to `notification_otp(email)` that looks like a local function call.

5. **Challenges with Traditional Approaches:**
   - In REST, making HTTP calls involves repeating code for each service interaction.
   - Handling issues like network reliability, retries, and error handling becomes cumbersome.

6. **RPC Purpose and Solution:**
   - RPC aims to standardize communication between services, irrespective of the language or communication protocol in use.
   - Abstracts out repetitive tasks and provides a unified approach for communication.

7. **Stub in RPC:**
   - Stub is a key concept in RPC, responsible for converting method requests and responses into a format understandable by both client and server.
   - Handles marshalling and unmarshalling of data between different services.

8. **Stub Generation:**
   - Stub generators automatically create code for client and server based on the interface definition language (IDL) used.
   - IDL defines the types, methods, and structures shared between services.

9. **Communication Protocols in RPC:**
   - RPC allows flexibility in choosing communication protocols, including TCP, UDP, HTTP, or even custom protocols.
   - It abstracts the transport layer, allowing developers to focus on business logic.

10. **Advantages of RPC:**
    - Makes remote calls appear local, enhancing developer productivity.
    - Strong API contracts, reducing errors and ensuring clarity in communication.
    - Supports modern languages and auto-generates client libraries.
    - Abstracts out mundane tasks, offering performance optimizations, security, and efficiency.

11. **Concerns and Challenges:**
    - Stub regeneration required on signature changes.
    - Testing RPC can be non-trivial and requires careful consideration.
    - Initial setup and getting started might be challenging.
    - Limited browser support for pure RPC implementation.

12. **Conclusion:**
    - RPC is a powerful paradigm for inter-service communication, providing a standardized, efficient, and language-agnostic approach.

