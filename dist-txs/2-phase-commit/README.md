# Distributed Transactions and Two-Phase Commit Protocol

## Introduction
Distributed transactions are a complex but crucial aspect of designing systems involving multiple services. This set of study notes is based on a YouTube video discussing distributed transactions and the Two-Phase Commit Protocol. The speaker uses the example of Zomato's 10-minute food delivery to illustrate the concepts.

### Video Overview
- **Title:** Distributed Transactions: Two-Phase Commit Protocol
- **Link:** [YouTube Video](https://www.youtube.com/watch?v=7FgU1D4EnpQ)
- **Duration:** Approximately 21 minutes

## Concepts Covered

### 1. Overview of Distributed Transactions
- **Challenge:** Distributed transactions involve multiple services coordinating to ensure that a series of steps either all happen or none happen.
- **Example:** Zomato's 10-minute food delivery is used as a case study to understand the challenges in coordinating microservices.

### 2. Zomato's 10-Minute Food Delivery
- **Business Approach:** Zomato purchases food in bulk based on data predicting demand, allowing them to guarantee 10-minute food delivery.
- **Engineering Challenges:** Ensuring availability of both food and delivery partners at the time of order placement.

### 3. Distributed Transaction Scenario
- **Example Flow:**
  1. User places an order.
  2. Order service talks to store service to reserve the food.
  3. Order service talks to delivery service to reserve a delivery partner.
  4. Both reservations must succeed for the order to be guaranteed.

### 4. Two-Phase Commit Protocol
- **Phases:**
  - **Prepare Phase:** Reserving items (food and delivery partner).
  - **Commit Phase:** Assigning or booking reserved items for the order.

- **Advantages:**
  - Guarantees atomic transactions.
  - Ensures isolation by preventing others from assigning reserved items.

- **Disadvantages:**
  - Pathetically slow due to sequential operations.
  - Prone to deadlock, requiring deadlock detection algorithms.

### 5. Handling Failures
- **Reservation Timeout:** To prevent perpetual reservations, a timer is set for each reservation.
- **Transaction Failure:** If either reservation fails, the entire transaction fails.
- **Commitment:** Successful reservations lead to the commit phase, where items are assigned to the order.

## Advantages and Disadvantages
### Advantages
1. **Atomic Transactions:** Guarantees that either all steps in a transaction happen or none happen.
2. **Isolation:** Ensures exclusivity by preventing others from accessing reserved items.

### Disadvantages
1. **Slow Performance:** The protocol's sequential nature makes it slow, impacting throughput.
2. **Deadlock Prone:** The separate reservation and commit phases increase the likelihood of creating cyclic dependencies, leading to deadlocks.


## Conclusion
Understanding distributed transactions and protocols like the Two-Phase Commit Protocol is crucial for designing reliable and consistent systems in a distributed environment. These notes provide an overview of the theoretical concepts discussed in the video, and the practical simulation in the next part promises to deepen this understanding.