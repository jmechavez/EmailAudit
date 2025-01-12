-- Drop and create `users` table
DROP TABLE IF EXISTS users;
CREATE TABLE users (
  user_id SERIAL PRIMARY KEY,
  fname VARCHAR(100) NOT NULL,
  lname VARCHAR(100) NOT NULL,
  id_no VARCHAR(100) NOT NULL,
  email VARCHAR(255) NOT NULL,
  status INT NOT NULL -- Changed to INT for 0 and 1
);

-- Insert data into the `clients` table
INSERT INTO users (fname, lname, id_no, email, status)
VALUES
  ('John', 'Doe', '12345', 'john.doe@example.com', 1),  -- 1 for active
  ('Jane', 'Smith', '67890', 'jane.smith@example.com', 0), -- 0 for inactive
  ('Alice', 'Johnson', '11223', 'alice.johnson@example.com', 1); -- 1 for active

