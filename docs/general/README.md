# Genral Concepts

## Network protocols

The video discusses the top 8 most popular network protocols:

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

- **Introduction to CDN:** Explains that a CDN, or content delivery network, has been around since the late 90s and was initially developed to speed up the delivery of static HTML content globally.

- **Evolution and Purpose:** Highlights that CDNs have evolved and are now essential whenever HTTP traffic is served, improving web service performance by bringing content closer to users.

- **CDN Infrastructure:** Describes the deployment of servers, known as Point of Presence (PoPs), worldwide to ensure users can access fast-edge servers. Introduces technologies like DNS-based routing and Anycast to direct user requests to the closest PoP.

- **Functionality of Edge Servers:** Explains that edge servers act as reverse proxies with content caches, reducing the load on origin servers by serving cached static content and transforming content into optimized formats on-the-fly.

- **Security and Availability Benefits:** States that modern CDNs provide security against DDoS attacks due to their large network capacity and improved availability by distributing content across many PoPs, making them resilient to hardware failures.

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

  