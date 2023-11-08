CREATE TABLE events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL
);

CREATE TABLE workshops (
   id INT AUTO_INCREMENT PRIMARY KEY,
   event_id INT NOT NULL,
   start_at TIMESTAMP NOT NULL,
   end_at TIMESTAMP NOT NULL,
   title VARCHAR(255) NOT NULL,
   description TEXT,
   FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE TABLE reservations (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  workshop_id INT NOT NULL,
  FOREIGN KEY (workshop_id) REFERENCES workshops(id)
);




