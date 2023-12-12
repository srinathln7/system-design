# Challenge

Setup a mySQL DB and create a schema for social network (users, posts, photos, profiles, following etc.) and insert data in (users and profile) in one transaction.

## Open MySQL server

```
sudo mysql -u root -p 
```

## Create and switch to the new database

```
CREATE DATABASE srinath_system_design;
```

To switch,

```
USE srinath_system_design;
```

### Create a social network schema

```

-- Create the schema

-- Create the tables
CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profiles (
    profile_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNIQUE,
    full_name VARCHAR(100),
    bio TEXT,
    avatar_url VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- Start a transaction
START TRANSACTION;

-- Insert data into the users table
INSERT INTO users (username, email) VALUES
    ('user1', 'user1@example.com'),
    ('user2', 'user2@example.com');

-- Get the last inserted user ID
SET @user1_id = LAST_INSERT_ID();

-- Insert data into the profiles table
INSERT INTO profiles (user_id, full_name, bio, avatar_url) VALUES
    (@user1_id, 'User One', 'Welcome to my profile!', 'avatar1.jpg');

-- Commit the transaction
COMMIT;


```


### Query the Tx

```

SELECT * FROM profiles;

SELECT * FROM users;

```

### Start another Tx


```
START TRANSACTION;

INSERT INTO profiles (user_id, full_name, bio, avatar_url) VALUES
    (@user2_id, 'User two', 'Welcome to my profile!', 'avatar2.jpg');

<Now kill the process before COMMIT> 

```

Constraints, cascades, and triggers are database concepts that contribute to maintaining data integrity and consistency in MySQL. Let's explore each of them with simple examples:

### 1. Constraints:

Constraints are rules that define the relationships between columns in a table. They help enforce data integrity by preventing the entry of invalid data. Common constraints include:

- **Primary Key Constraint:** Ensures that a column or a set of columns uniquely identifies each row in a table.
  
- **Foreign Key Constraint:** Establishes a link between two tables, ensuring referential integrity.

- **Unique Constraint:** Ensures that all values in a column are unique.

- **Check Constraint:** Enforces a condition that must be true for each row.

#### Example:

```sql
-- Create tables with constraints
CREATE TABLE users (
    user_id INT PRIMARY KEY,
    username VARCHAR(50) UNIQUE
);

CREATE TABLE posts (
    post_id INT PRIMARY KEY,
    user_id INT,
    content TEXT,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
```

In this example, the `users` table has a primary key constraint on the `user_id` column, and the `posts` table has a foreign key constraint referencing the `user_id` column in the `users` table.

### 2. Cascades:

Cascading actions define what should happen when a referenced row is updated or deleted. For example, in a foreign key relationship, you might specify that if a referenced row is deleted, the corresponding rows in the referencing table should also be deleted (CASCADE). This helps maintain referential integrity.

#### Example:

```sql
-- Create a table with a CASCADE constraint
CREATE TABLE orders (
    order_id INT PRIMARY KEY,
    user_id INT,
    order_date DATE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
```

In this example, when a user is deleted from the `users` table, all corresponding orders in the `orders` table with the matching `user_id` will also be deleted due to the `ON DELETE CASCADE` clause.

### 3. Triggers:

Triggers are special stored procedures that are automatically executed (or "triggered") in response to specific events, such as INSERT, UPDATE, or DELETE operations on a table. They are used to enforce business rules or perform additional actions when certain changes occur in the database.

#### Example:

```sql
-- Create a trigger to update a timestamp on user profile changes
CREATE TRIGGER before_update_profile
BEFORE UPDATE ON profiles
FOR EACH ROW
SET NEW.updated_at = NOW();
```

In this example, a trigger is created to automatically update the `updated_at` timestamp column in the `profiles` table every time a row is updated.

### How They Help in Consistency:

- **Constraints:** Ensure that data adheres to predefined rules, preventing the entry of inconsistent or invalid data.

- **Cascades:** Maintain referential integrity by defining actions to be taken when referenced rows are updated or deleted, preventing orphaned or inconsistent data.

- **Triggers:** Allow you to automate actions based on specific events, helping enforce additional business rules or maintain consistency.

Together, these mechanisms contribute to the overall consistency and integrity of the data in a MySQL database, reducing the risk of errors and ensuring that the data remains accurate and reliable.


### Isolation Levels

#### Repeatable Reads

Provides consistent views within the transaction no matter if the other transaction commits a new value.

```
select @@transaction_ISOLATION;
+-------------------------+
| @@transaction_ISOLATION |
+-------------------------+
| REPEATABLE-READ         |
+-------------------------+

describe users;

```



Open two terminals side-to-side and start firing two Txs concurrently

Terminal 1 

```
SELECT * FROM users WHERE user_id=1;

+---------+----------+-------------------+---------------------+
| user_id | username | email             | created_at          |
+---------+----------+-------------------+---------------------+
|       1 | user1    | user1@example.com | 2023-12-12 09:05:09 |
+---------+----------+-------------------+---------------------+


UPDATE users SET email='user1@isolation.com' WHERE user_id=1;


<Repeat query>

COMMIT;

```



Terminal 2

```
SELECT * FROM users WHERE user_id=1;

+---------+----------+-------------------+---------------------+
| user_id | username | email             | created_at          |
+---------+----------+-------------------+---------------------+
|       1 | user1    | user1@example.com | 2023-12-12 09:05:09 |
+---------+----------+-------------------+---------------------+


<AFTER Tx1 COMMIT>

SELECT * FROM users WHERE user_id=1;

```

You will still see the same value. Once you `COMMIT` or `ROLLBACK` the transaction, and fire the query again, we should be able to see the latest updated value set by Tx1.


#### Serializable Reads

```
SET SESSION TRANSACTION ISOLATION LEVEL SERIALIZABLE;
UPDATE users SET email='user1@example.com' WHERE user_id=1;
```

#### Read Committed

```
SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED;
UPDATE users SET email='user1@readcommitted.com' WHERE user_id=1;
```

#### Read Uncommitted

```
SET SESSION TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;
UPDATE users SET email='user1@example.com' WHERE user_id=1;
```


