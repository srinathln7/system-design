# General Concepts

## HTTP Status Codes

1. **2xx - Success**:
   - **200 OK**: Indicates that the request was successful.
   - **201 Created**: Indicates that a new resource was successfully created.
   - **204 No Content**: Indicates that the request was successful, but there is no content to return.

2. **4xx - Client Error**:
   - **400 Bad Request**: Indicates that the server could not understand the request due to invalid syntax.
   - **401 Unauthorized**: Indicates that authentication is required and has failed or has not been provided.
   - **403 Forbidden**: Indicates that the server understood the request but refuses to authorize it.
   - **404 Not Found**: Indicates that the requested resource could not be found on the server.

3. **429 - Too Many Requests**:
   - Indicates that the user has sent too many requests in a given amount of time.

4. **5xx - Server Error**:
   - **500 Internal Server Error**: Indicates that the server encountered an unexpected condition that prevented it from fulfilling the request.
   - **502 Bad Gateway**: Indicates that the server, while acting as a gateway or proxy, received an invalid response from the upstream server it accessed in attempting to fulfill the request.
   - **503 Service Unavailable**: Indicates that the server is currently unable to handle the request due to temporary overloading or maintenance of the server.

5. **3xx - Redirection**:
   - **301 Moved Permanently**: Indicates that the requested resource has been permanently moved to a new location.
   - **302 Found**: Indicates that the requested resource temporarily resides under a different URL.
   - **304 Not Modified**: Indicates that the resource has not been modified since the version specified by the request headers, resulting in no content being returned.

6. **1xx - Informational**:
   - **101 Switching Protocols**: Indicates that the server understands and is willing to comply with the client's request to change the protocol being used on this connection.

**Pro-Tips**:
- Start with the basics when tackling status code issues: check request headers and body, ensure the correct HTTP method is used, and confirm the endpoint URL.
- Use tools like Postman or Insomnia for testing requests and inspecting responses.
- Check server logs for additional insights if available.

These status codes are essential for understanding the outcome of HTTP requests and troubleshooting issues in web applications, mobile apps, or APIs.


## HTTPS:

1. **Introduction to HTTPS**:
   - HTTPS encrypts data sent over the Internet, making it unreadable by anyone other than the sender and receiver.
   - HTTPS is an extension of HTTP, with data sent in encrypted form using TLS (Transport Layer Security).

2. **TLS Handshake**:
   - **Step 1**: The browser establishes a TCP connection with the server.
   - **Step 2**: TLS handshake begins with the client sending a "Client Hello" message to the server, specifying TLS version and supported cipher suites.
   - **Step 3**: The server selects the TLS version and cipher suite and sends them back to the client in a "Server Hello" message along with its certificate, including the public key.
   - **Step 4**: The client generates an encryption key, encrypts it with the server's public key, and sends it back to the server.
   - **Step 5**: The server decrypts the key using its private key, establishing a shared encryption key for secure communication.

3. **Symmetric Encryption**:
   - The shared encryption key is used for symmetric encryption, where data encrypted on one end can only be decrypted by the other end.
   - Symmetric encryption is more efficient for bulk data transmission compared to asymmetric encryption.

4. **TLS Versions**:
   - TLS 1.2 and TLS 1.3 are discussed, with TLS 1.3 optimized to reduce the number of network round trips required for the handshake.

5. **Asymmetric Encryption**:
   - RSA is used as an example of asymmetric encryption for securely exchanging the symmetric session key.
   - Diffie-Hellman is mentioned as a more common method in TLS 1.3 for exchanging session keys without transmitting public keys over the network.

6. **Conclusion**:
   - The video concludes by encouraging viewers to learn more about system design through books and newsletters.

Overall, the transcript provides a clear explanation of how HTTPS works and the TLS handshake process, including key concepts like encryption, cipher suites, and asymmetric encryption methods.

## Network protocols

1. **HTTP (Hypertext Transfer Protocol)**:
   - HTTP is the backbone of web development, using a request-response model for clients to request resources from servers.
   - It defines methods like GET, POST, PUT, and DELETE for different operations on the server.

2. **HTTPS (Hypertext Transfer Protocol Secure)**:
   - HTTPS builds on HTTP by adding encryption through Transport Layer Security (TLS), ensuring data confidentiality and authentication.
   - It prevents man-in-the-middle attacks and is essential for secure websites, such as those used for banking and shopping.

3. **HTTP/3**:
   - HTTP/3 aims to improve speed and security by using QUIC, built on UDP rather than TCP, to optimize performance and minimize lag.
   - It brings encryption by default at the transport layer and speeds up initial connection setup.

4. **WebSocket**:
   - WebSocket provides full-duplex bidirectional communication on a single TCP connection, enabling real-time collaboration and live data streams.
   - It supports sending small messages instantly with low overhead, ideal for chat, gaming, or real-time updates.

5. **TCP (Transmission Control Protocol)**:
   - TCP prioritizes reliability over speed, ensuring robust performance through error checking, transmission controls, and ordered data delivery.
   - It adapts to network conditions with flow control and retransmissions.

6. **UDP (User Datagram Protocol)**:
   - UDP focuses on speed over reliability, minimizing overhead for use cases like gaming, voice, and streaming.
   - It lacks reliability compared to TCP but is often combined with application-layer integrity checks for a balance between speed and reliability.

7. **SMTP (Simple Mail Transfer Protocol)**:
   - SMTP is the standard for transferring email messages between mail servers, essential for configuring mail services and ensuring proper email delivery.

8. **FTP (File Transfer Protocol)**:
   - FTP allows efficient uploading and downloading of files between hosts and remains ubiquitous for file-based workflows, especially in financial institutions.

**Pros**:
- Each protocol serves specific purposes, enabling efficient communication and data exchange over networks.
- Understanding these protocols allows developers to build fast, secure systems tailored to their requirements.

**Cons**:
- Some protocols may have limitations or vulnerabilities that need to be addressed, such as the potential for data corruption in UDP or security risks in HTTP without encryption.

In conclusion, these network protocols play crucial roles in enabling communication and data exchange across the internet, providing the foundation for various applications and services we use daily.

## CDN

- **Introduction to CDN:** A CDN is a geographically distributed group of servers that caches content close to end users. It allows for quick transfer of assets needed for loading Internet content, including HTML pages, JavaScript files, stylesheets, images, and videos. The primary purpose of a CDN is to reduce latency in communication between websites and users. CDNs improve efficiency by introducing intermediary servers between clients and website servers, reducing web traffic, bandwidth consumption, and improving user experience.

- **Evolution and Purpose:** Highlights that CDNs have evolved and are now essential whenever HTTP traffic is served, improving web service performance by bringing content closer to users. The majority of web traffic today is served through CDNs, including traffic from major sites like Facebook, Netflix, and Amazon. CDNs can also help protect websites against common malicious attacks like Distributed Denial of Service (DDOS) attacks.

- **CDN Infrastructure:** Describes the deployment of servers, known as Point of Presence (PoPs), worldwide to ensure users can access fast-edge servers. Introduces technologies like DNS-based routing and Anycast to direct user requests to the closest PoP.

**Difference from Web Hosting**: While a CDN does not host content and cannot replace the need for proper web hosting, it caches content at the network edge to improve website performance. Many websites opt for CDNs to address performance issues encountered with traditional hosting services.

- **Functionality of Edge Servers:** Explains that edge servers act as reverse proxies with content caches, reducing the load on origin servers by serving cached static content and transforming content into optimized formats on-the-fly.
 
- **Benefits of Using a CDN**:
   - **Improving website load times**: By distributing content closer to website visitors, CDNs reduce load times, lower bounce rates, and increase visitor engagement.
   - **Reducing bandwidth costs**: CDNs optimize caching and other techniques to reduce data transfer from origin servers, thus reducing hosting costs.
   - **Increasing content availability and redundancy**: CDNs handle large traffic volumes and hardware failures better than origin servers.
   - **Improving website security**: CDNs offer DDoS mitigation, security certificate improvements, and other security optimizations. 

   ### **Improvement of Website Load Times**:
   - CDNs reduce distance between users and website resources, utilize hardware and software optimizations, and reduce data transfer through file size reduction and TLS/SSL certificate optimization.

   ### **Reliability and Redundancy**:
   - CDNs ensure uptime through load balancing, intelligent failover, and Anycast routing to transfer traffic to available data centers in case of hardware failures or technical issues.

   ### **Data Security**:
   - CDNs protect data with TLS/SSL certificates to ensure authentication, encryption, and integrity.

   ### **Bandwidth Expense**:
   - CDNs reduce bandwidth costs by minimizing origin server requests and optimizing data transfer. For example, Cloudflare CDN reduces origin requests and bandwidth    costs. 

- **How CDNs Work**:
   - CDNs are a network of servers aimed at delivering content quickly, cheaply, reliably, and securely.
   - They place servers at Internet exchange points (IXPs) to improve speed and connectivity.
   - CDNs make optimizations on client/server data transfers, place data centers strategically, enhance security, and manage various types of failures and Internet congestion.
   - Work on principles of **caching, dynamic acceleration, and edge logic computations** to optimize data delivery.

- **Usecases**
   - High-speed content delivery for global, whole-site experiences.
   - Real-time streaming for rich media files.
   - Multi-user scaling to support large numbers of concurrent users, as seen in gaming platforms like King.    

## Forward vs Reverse proxy

- **Proxy Types:** Explains the two common types of proxies: forward proxy and reverse proxy.
  
- **Forward Proxy (Client to Internet):**
  - Acts as a middleman between client machines and the internet.
  - Protects client's online identity by hiding their IP address.
  - Can be used to bypass browsing restrictions, often implemented by institutions.
  - Allows blocking access to specific content through filtering rules.
  - Requires client configuration, but transparent proxies streamline the process for large institutions.

- **Reverse Proxy (Internet to Web Servers):**
  - Sits between the internet and web servers, handling requests on behalf of clients.
  - Protects a website by hiding its IP addresses, making DDoS attacks more challenging.
  - Enables load balancing to distribute traffic among multiple web servers, preventing overload.
  - Caches static content, improving performance by quickly serving locally cached versions.
  - Manages SSL encryption, freeing up origin servers from computationally expensive SSL handshakes.
  - Commonly deployed in multiple layers, with examples like Cloudflare as an edge service and additional layers like API gateways or load balancers.

- **Deployment Layers:** Describes the typical deployment layers in modern websites, involving edge services, API gateways or load balancers, and clusters of web servers.

## API Architecture styles

The video discusses the top 6 most popular API architecture styles, each with its own design philosophy and use cases:

1. **SOAP (Simple Object Access Protocol)**:
   - SOAP is a mature, comprehensive, and XML-based protocol commonly used in industries like financial services and payment gateways where security and reliability are paramount.
   - However, SOAP can be overkill for lightweight mobile apps or quick prototypes due to its complexity and verbosity.

2. **RESTful APIs (Representational State Transfer)**:
   - RESTful APIs are easy to implement and use HTTP methods for communication.
   - They are popular for web services like Twitter or YouTube, but may not be the best fit for real-time data or highly connected data models.

3. **GraphQL**:
   - GraphQL is both an architectural style and a query language that allows clients to request specific data as needed, reducing over-fetching or under-fetching.
   - It is flexible and efficient, making it suitable for applications with complex data requirements, but it has a steep learning curve and may be overkill for simpler applications.

4. **gRPC**:
   - gRPC is a modern, high-performance RPC (Remote Procedure Call) framework that uses Protocol Buffers.
   - It is favored for microservices architectures due to its efficiency, but it may pose challenges with browser clients due to limited browser support.

5. **WebSocket**:
   - WebSocket provides real-time, bidirectional, and persistent connections, making it ideal for applications requiring low-latency data exchange, such as live chat or real-time gaming.
   - However, using WebSocket for applications that don't require real-time data may introduce unnecessary overhead.

6. **Webhook**:
   - Webhook is an event-driven architecture style that uses HTTP callbacks and asynchronous operations.
   - It is useful for scenarios like GitHub, where systems need to be notified of events like new commits, but may not be suitable for synchronous communication or immediate response requirements.

**Pros**:
- Each architecture style offers specific benefits tailored to different project requirements.
- They provide flexibility and scalability, allowing developers to choose the most suitable architecture for their application's needs.

**Cons**:
- Choosing the wrong architecture style for a project can lead to unnecessary complexity or inefficiency.
- Some styles may have a steeper learning curve or limited support in certain environments, requiring additional resources for implementation and maintenance.

In conclusion, there is no one-size-fits-all solution when it comes to API architecture styles. Developers should carefully evaluate their project requirements and choose the style that best fits their needs for efficiency, scalability, and ease of implementation.

### GraphQL vs REST

1. **Definition:** GraphQL provides a schema of the data in the API and allows clients to request exactly what they need. It sits between clients and backend services, aggregating multiple resource requests into a single query and supporting mutations and subscriptions.

2. **Similarities with REST:**
   - Both GraphQL and REST send HTTP requests and receive HTTP responses.
   - They both use URLs for making requests and can return JSON responses.

3. **Differences from REST:**
   - With GraphQL, clients specify the exact resources and fields they want, while REST APIs may include related resources in the response without client control.
   - GraphQL uses a schema to define available resources instead of URLs, allowing complex queries to fetch additional data based on schema-defined relationships.
   - While REST requests can be made using common tools like cURL or web browsers, GraphQL requires heavier tooling support on both client and server sides.

4. **Drawbacks of GraphQL:**
   - GraphQL requires a larger upfront investment in tooling support compared to REST, especially for simple CRUD APIs.
   - It's more difficult to cache GraphQL responses effectively compared to REST due to its single point of entry and default use of HTTP POST, limiting the leverage of HTTP caching mechanisms.
   - GraphQL's flexibility in querying specific data poses risks of unexpected resource-intensive queries, which can potentially bring down backend services if not properly managed.

5. **Considerations for Using GraphQL:**
   - Software engineering involves trade-offs, and there's no one right answer. Organizations should carefully consider the trade-offs, complexities, and risks associated with adopting GraphQL for high-scale production use.

Overall, GraphQL offers powerful features for flexible data querying but requires careful consideration and management of its complexities and potential drawbacks, especially in high-scale production environments.

## API testing

- **Introduction to API Testing:**
  - Highlights the importance of API testing as a critical backend process ensuring the functionality of digital systems.

- **Nine Types of API Testing:**
  1. **Smoke Testing:**
     - Quick test drive for APIs.
     - Basic tests to confirm functionality before launch.

  2. **Functional Testing:**
     - Similar to QA testing for software programs.
     - Detailed test cases to ensure all API functionalities meet specific requirements.

  3. **Integration Testing:**
     - Ensures seamless communication between diverse APIs and cloud services.
     - Verifies coordination among different APIs, such as airline reservation and payment systems.

  4. **Regression Testing:**
     - Replays existing test cases against updated API versions to check for defects.
     - Enables developers to iterate without worrying about collateral damage.

  5. **Load Testing:**
     - Evaluates real-world performance by subjecting the API to high user volumes.
     - Ensures responsiveness and stability during peak usage times.

  6. **Stress Testing:**
     - Simulates extreme traffic spikes and edge conditions beyond normal usage.
     - Prepares APIs for maintaining responsiveness and stability during challenging scenarios.

  7. **Security Testing:**
     - Probes APIs for weaknesses or loopholes that could lead to unauthorized access or cyber threats.
     - Fortifies APIs against attacks to protect sensitive data.

  8. **UI Testing:**
     - Ensures APIs contribute to a seamless user experience.
     - Fast response times directly correlate to positive UI perceptions, e.g., in ride-sharing apps.

  9. **Fuzz Testing:**
     - Throws malformed, unexpected, or randomly generated data at APIs to identify edge cases.
     - Hardens systems against unpredictable real-world usage.


## API optimization techniques

The video discusses seven optimization techniques to improve API performance:

1. **Caching**:
   - Caching stores the result of expensive computations to reuse them later without redoing the computation.
   - Frequently accessed endpoints with the same request parameters can benefit from caching in Redis or Memcached, reducing database hits and improving speed.

2. **Connection Pooling**:
   - Maintaining a pool of open connections instead of opening a new database connection for each API call can improve throughput.
   - Reusing connections reduces the overhead of establishing new connections for each request.

3. **Avoiding N+1 Query Problems**:
   - The N+1 query problem occurs when accessing data and related entities, resulting in excessive database queries.
   - Fetching data in a single query or optimizing queries to reduce round trips to the database can improve efficiency.

4. **Pagination**:
   - Breaking large API responses into smaller, manageable pages using pagination parameters like limit and offset reduces data transfer and client-side load.

5. **Using Lightweight JSON Serializers**:
   - Fast serialization libraries minimize the time spent converting data into JSON format, improving response times for JSON responses.

6. **Compression**:
   - Enabling compression on large API response payloads reduces the amount of data transferred over the network, improving performance.
   - Efficient compression algorithms like Brotli and CDN support for compression further enhance performance.

7. **Asynchronous Logging**:
   - Asynchronous logging involves placing log entries into an in-memory buffer by the main application thread while a separate logging thread writes the log entries.
   - This reduces the time taken to write logs, particularly in high-throughput systems, but there's a small risk of losing logs if the application crashes before logs are written.

**Pros**:
- These optimization techniques can significantly improve API performance, making applications faster and more efficient.
- They address common performance bottlenecks and reduce latency, enhancing the user experience.

**Cons**:
- Implementing some of these techniques, such as caching and connection pooling, may introduce additional complexity to the application.
- Asynchronous logging carries a risk of losing logs if the application crashes before logs are written, requiring careful implementation and monitoring.

Overall, these optimization techniques provide valuable strategies for improving API performance when implemented judiciously and tailored to specific application requirements.

### RESTful API Design Tips

1. **Clear Naming:** Use straightforward and logical names for APIs. Choose descriptive names that convey the purpose of the API, such as using plurals for resource collections.

2. **Reliability through Idempotent APIs:** Ensure that making the same API call multiple times has the same effect as calling it once. Idempotent APIs prevent duplicate actions and ensure consistency, especially for operations like creating, reading, updating, and deleting resources.

3. **Versioning:** Implement versioning in APIs to support backward compatibility and allow for updates without breaking existing applications. Versioning helps developers using older versions of the API to upgrade at their own pace.

4. **Pagination:** Control the amount of data returned by APIs by implementing pagination. This prevents overwhelming API consumers with large data sets and enhances performance by fetching data in smaller chunks.

5. **Clear Query String for Sorting and Filtering:** Use clear query strings for sorting and filtering API data. This makes response data easy to understand and allows for additional sorting and filtering criteria to be added over time without breaking existing integrations.

6. **Security:** Prioritize security when designing APIs, especially for sensitive credentials like API keys. Leverage HTTP headers over URLs for passing credentials and implement robust access controls and encryption to protect against potential security threats.

7. **Keep Cross-Resource References Simple:** Use clear linking between connected resources to avoid cluttering with long query strings. Direct paths make associations clear for developers using the API.

8. **Plan for Rate Limiting:** Implement rate limiting to protect APIs from overload and abuse. Set request quotas based on various dimensions like source IP addresses, user accounts, or endpoint categories to ensure fair usage and protect infrastructure from denial-of-service (DoS) attacks.

By following these best practices, we can create robust, reliable, and secure APIs that are a pleasure to work with and can be relied upon by companies and developers for long-term use.


## Deployment Strategies
- **Introduction to Deployment Strategies:**
  - Emphasizes the importance of deploying code to production and the challenges involved.

- **Big Bang Deployment:**
  - Deploys all changes at once, causing downtime.
  - Requires thorough preparation and testing.
  - Typically used for major updates or intricate database upgrades.

- **Rolling Deployment:**
  - Incrementally updates different parts of the system over time.
  - Reduces downtime and allows early issue detection.
  - Balances risk and user impact but lacks targeted rollouts.

- **Blue-Green Deployment:**
  - Maintains two identical production environments, allowing seamless transitions and easy rollbacks.
  - Doubles infrastructure and resource needs, introducing complexity.

- **Canary Deployment:**
  - Tests new versions on a small subset of servers or users before full rollout.
  - Offers safety and control, with targeted rollouts based on user-specific criteria.
  - Requires careful monitoring and automated testing.

- **Feature Toggle:**
  - Manages specific new features within the application by introducing toggles to turn them on or off.
  - Offers excellent control over new features and allows for targeted user testing.
  - Can add complexity to the codebase if not managed properly.

- **Conclusion and Discussion Prompt:**
  - Encourages considering the characteristics of the application and user expectations when choosing a deployment strategy.


## Shipping code to production

1. **User Feedback and Requirements Gathering**: The product team collects user feedback and requirements, which are then broken down into smaller work items or user stories.

2. **Sprint Planning**: The product and engineering teams plan sprints typically lasting 1-2 weeks, during which developers pick up work items. Larger projects may span multiple sprints.

3. **Design and Architecture**: For larger initiatives, there's often an RFC (Request for Comments) or design document process to align on high-level architecture and technical approach upfront.

4. **Development**: Developers use Git or similar version control systems to create feature branches for building new functionality. They also develop migration scripts for database schema changes, which require careful design and testing.

5. **Code Review and Merge**: Developers open pull requests for code review, catching issues early through knowledge sharing. After approval, code is merged into the main branch after running comprehensive unit tests.

6. **Continuous Integration and Deployment (CI/CD)**: CI/CD pipelines, powered by tools like GitHub Actions or Jenkins, automatically build, test, and deploy code through multiple environments like dev, test, and staging. This ensures consistent validation and reduces surprises in production.

7. **Quality Assurance (QA) and User Acceptance Testing (UAT)**: QA engineers validate functionality, run regressions, security scans, and performance tests. UAT involves the product team, QA, and developers validating features together before release candidates proceed to production.

8. **Production Rollout and Monitoring**: Techniques like canary releases and feature flags are used to slowly roll out changes and reduce risk in production. SREs monitor metrics, logs, and traffic for any production issues, prioritizing and fixing bugs as needed. Product and engineering teams also monitor analytics to ensure the feature works as expected and aligns with business metrics.

In summary, the process involves iterative development, testing, and incremental rollout, with a focus on collaboration, automation, and continuous improvement.


## CI/CD 

CI/CD, or Continuous Integration and Continuous Delivery, is a set of practices aimed at automating the software development process from initial code commit to deployment. Let's break down the key points mentioned in the video:

1. **Continuous Integration (CI)**:
   - CI involves automating the process of merging code changes into the shared repository frequently, often triggered by each commit.
   - The CI process runs a series of tasks to ensure that the commit is safe to merge into the main branch, typically including building and testing.
   - Maintaining a good CI process requires a set of robust tests with sufficient coverage, which can be challenging but essential for ensuring code quality.
   - Common CI tools include GitHub Actions, Buildkite, Jenkins, CircleCI, and TravisCI.
   - Test tools, such as Jest for JavaScript unit testing and Playwright or Cypress for integration testing, are essential components of the CI process.

2. **Continuous Deployment (CD)**:
   - CD involves automatically deploying new code changes to production environments.
   - Real continuous deployment is challenging and less common than CI, especially for complex stateful systems.
   - Many teams practice CD only for stateless systems like APIs or web servers.
   - Techniques like feature flags and canary deployments help minimize risks associated with continuous deployment.
   - Manual deployment processes are still prevalent for complex stateful systems, requiring dedicated platform teams and fixed deploy cadences.
   - Tools like GitHub Actions, Buildkite, Jenkins, and infrastructure-specific tools such as ArgoCD for Kubernetes are commonly used for CD tasks.

**Pros**:
- Automation of build, test, and deployment processes saves time and reduces manual errors.
- Continuous integration ensures that code changes are tested early and frequently, leading to better code quality.
- Continuous deployment enables faster and more reliable delivery of features to end-users.

**Cons**:
- Implementing CI/CD requires initial setup and ongoing maintenance, which can be time-consuming.
- Real continuous deployment may not be feasible for all types of systems, especially complex stateful systems.
- Over-reliance on automation may lead to neglecting manual testing and oversight, potentially introducing hidden bugs.

Overall, CI/CD is a powerful practice that can significantly improve the efficiency and reliability of the software development lifecycle, but its effectiveness depends on factors such as system complexity and team resources.

## Architecture patterns

- **Introduction to Software Architecture:**
  - Compares software architecture to building foundations, highlighting its importance in application development.

- **Layered Architecture:**
  - Separates system components into distinct layers (e.g., presentation, business logic, data access).
  - Promotes separation of concerns and abstraction, facilitating easier maintenance and scalability.

- **Event-Driven Architecture:**
  - Facilitates communication between loosely coupled components through events.
  - Example includes Command Query Responsibility Segregation (CQRS), promoting decoupling of read and write operations.

- **Microkernel Architecture:**
  - Emphasizes separating core system functionality into a microkernel and extending functionality through plugins or add-ons.
  - Promotes extensibility, ease of maintenance, and fault isolation.

- **Microservices Architecture:**
  - Decomposes applications into small, independent services communicating via APIs.
  - Promotes agility and rapid innovation but introduces complexity in inter-service communication and data consistency.

- **Monolithic Architecture:**
  - Bundles all application components into a single codebase and runs as a single unit.
  - Simplifies development and deployment, commonly used in startups and smaller applications.
  - 'Modular monolith' approach retains benefits while emphasizing clear modular boundaries for easier maintenance and scalability.

- **Conclusion and Discussion Prompt:**
  - Encourages considering specific challenges, requirements, and contexts when choosing an architecture pattern.
  

## DevOps vs SRE vs Platform Engineering

- **Introduction to DevOps, SRE, and Platform Engineering:**
  - DevOps aims to bridge the gap between development and operations teams, fostering collaboration throughout the software development lifecycle.
  - SRE (Site Reliability Engineering), pioneered by Google, focuses on building resilience into complex systems using software engineering approaches.
  - Platform Engineering involves developing robust platforms that enable developers to efficiently create high-quality software aligned with business objectives.

- **DevOps:**
  - "You build it, you run it" principle exemplifies DevOps, where teams responsible for building software also deploy and maintain it.
  - Roles often blend and evolve, with individuals adapting and incorporating elements from various domains.

- **SRE (Site Reliability Engineering):**
  - SREs ensure software systems are robustly built, function as designed, and can handle real-world demands using software engineering approaches.
  - They wear multiple hats, sometimes overlapping with Platform Engineering by building tools to simplify processes for developers.

- **Platform Engineering:**
  - Platform Engineering, exemplified by teams like Netflix, focuses on developing reliable platforms that enable rapid software development aligned with business objectives.
  - It facilitates developers' focus on crafting features while taking care of the underlying infrastructure and processes.

- **Common Vision and Collaboration:**
  - DevOps, SRE, and Platform Engineering share a vision of enhancing collaboration, automation, and efficiency in software development and operations.
  - Professionals should adapt and evolve with the industry, with these roles co-existing to contribute to a robust and efficient tech ecosystem.

- **Conclusion and Discussion Prompt:**
  - Encourages professionals to focus less on rigid job titles and more on adapting and evolving with the industry.


  ## Process vs Threads

  - **Process vs Thread:**
  - **Definition:**
    - A program is an executable file containing code or processor instructions stored on disk.
    - When loaded into memory and executed by the processor, it becomes a process.
    - A process includes resources managed by the operating system, such as processor registers, memory pages, etc.
    - Each process has its own memory address space, ensuring isolation from other processes.
  - **Threads:**
    - A thread is the unit of execution within a process.
    - A process has at least one thread (main thread) but can have multiple threads.
    - Each thread has its own stack, while threads within a process share a memory address space.
    - Threads can communicate via shared memory space, but one misbehaving thread can affect the entire process.
  - **Context Switching:**
    - Context switching involves switching one process or thread out of the CPU to allow another to run.
    - It is costly, involving saving/loading registers, memory page switching, and updating kernel data structures.
    - Switching between threads is generally faster than between processes due to shared memory address space.
    - Mechanisms like fibers and coroutines aim to minimize context switching costs further, often cooperatively scheduled by the application.
  - **Conclusion:**
    - Processes and threads have distinct characteristics and roles in program execution.
    - Understanding their differences is crucial for efficient resource management and program design.

## Bloom Filters

1. **Definition:** A Bloom filter is a probabilistic data structure that provides a "probably yes" or "firm no" answer to whether an element is in a set. It can have false positives (saying an element is present when it's not) but no false negatives (saying an element is absent when it's actually present). Once an item is put in the Bloom filter it cannot be removed.

2. **Tradeoff:** Bloom filters trade memory efficiency for occasional false positives. They consume less memory compared to data structures like hash tables, which provide a perfect answer but with higher memory usage.

3. **Use Cases:** Bloom filters are used in various applications:
   - NoSQL databases use them to reduce disk reads for keys that don't exist.
   - Content delivery networks (CDNs) use them to prevent caching "one-hit-wonders" and improve caching hit rates.
   - Web browsers have used them to identify malicious URLs, although more efficient solutions are now required due to the increasing number of malicious URLs.
   - Password validators use them to prevent users from using weak passwords.

4. **How It Works:** Bloom filters use hash functions to set bits in a bit array. When inserting an element, multiple hash functions are applied to generate indexes, and corresponding bits are set to 1. When querying an element's existence, the same hash functions are applied, and if all corresponding bits are set to 1, the element is probably in the set.

5. **False Positives:** Bloom filters can produce false positives when an element is not present in the set but the filter mistakenly indicates its presence. The frequency of false positives can be controlled by choosing an appropriate size for the Bloom filter based on the expected number of entries.

Overall, Bloom filters provide a space-efficient way to check for membership in a set, although they may occasionally yield false positives. This tradeoff makes them suitable for scenarios where memory efficiency is critical, and occasional errors can be tolerated.