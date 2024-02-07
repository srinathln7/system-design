# API Composition Pattern in Microservices

## **Introduction to API Composition:**
- Microservices architecture may require interacting with multiple services to fulfill a specific use case.
- API Composition is a pattern that addresses the need to query multiple microservices for compiling a final response.

## **Implementing API Composition:**
- API Composition involves placing a middleman (composer) between the user and microservices to hide complexities.
- API Gateway is a common component acting as a composer, placed in front of microservices.
- Popular API Gateway solutions include KrakenD and Kong.
- The API Gateway forwards requests to multiple microservices, aggregates responses, and returns them to the user.

## Advantages of API Composition
1. **Simplicity:**
   - API Composition is simple to implement by adding an API Gateway.
2. **Single Point of Interaction:**
   - Users interact with a single entry point (API Gateway), simplifying the user experience.
3. **Hiding Implementation Complexity:**
   - The pattern hides the details of interactions with multiple services from end-users.
4. **Security and Limiting:**
   - The API Gateway serves as a point for implementing security, limiting, and caching.
5. **Covering Bad Design Decisions:**
   - Allows for modifying responses in the composer to hide or improve bad design decisions.
6. **Facilitating Microservices Transition:**
   - Useful when transitioning from a monolith to microservices, providing an abstraction layer.

## Disadvantages of API Composition
1. **Data Size Consideration:**
   - Handling large data responses from multiple services can be resource-intensive.
2. **Availability Challenges:**
   - Dependent on the availability of all services; downtime in one service affects the overall availability.
3. **Transactional Data Consistency:**
   - Achieving transactional data consistency across multiple services is challenging.
4. **Management Overhead:**
   - The API Gateway (Composer) needs proper management, maintenance, and monitoring.
5. **Potential Bottleneck:**
   - The API Gateway might become a bottleneck under a high load, requiring careful monitoring.

## Conclusion:
- API Composition is a valuable pattern for simplifying microservices interactions.
- Understanding its advantages and disadvantages is crucial for effective implementation.
- Considerations include data size, availability, transactional consistency, and the management overhead of the API Gateway.

These notes cover key points regarding API Composition in Microservices, emphasizing its implementation, advantages, and potential challenges.