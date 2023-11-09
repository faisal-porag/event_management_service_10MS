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

-- //////////////////////////////SAMPLE DATA///////////////////////////////////////// --

INSERT INTO events (title, start_at, end_at) VALUES
    ('Event 1', '2023-11-10 09:00:00', '2023-11-10 17:00:00'),
    ('Event 2', '2023-11-13 11:00:00', '2023-11-14 13:00:00'),
    ('Event 3', '2023-11-12 11:00:00', '2023-11-14 13:00:00'),
    ('Event 4', '2023-11-16 11:00:00', '2023-11-18 13:00:00'),
    ('Event 5', '2023-11-14 11:00:00', '2023-11-18 13:00:00'),
    ('Event 6', '2023-11-16 11:00:00', '2023-11-16 13:00:00'),
    ('Event 7', '2023-11-11 10:00:00', '2023-11-11 15:00:00'),
    ('Event 8', '2023-11-11 10:00:00', '2023-11-11 15:00:00'),
    ('Event 9', '2023-11-11 10:00:00', '2023-11-11 15:00:00'),
    ('Event 10', '2023-11-19 11:00:00', '2023-11-21 13:00:00');

INSERT INTO workshops (event_id, start_at, end_at, title, description) VALUES
    (1, '2023-11-10 10:00:00', '2023-11-10 12:00:00', 'Workshop 1', 'Description 1'),
    (2, '2023-11-11 11:00:00', '2023-11-11 13:00:00', 'Workshop 2', 'Description 2'),
    (2, '2023-11-11 11:00:00', '2023-11-11 13:00:00', 'Workshop 3', 'Description 3'),
    (2, '2023-11-13 11:00:00', '2023-11-14 13:00:00', 'Workshop 4', 'Description 4'),
    (1, '2023-11-12 11:00:00', '2023-11-14 13:00:00', 'Workshop 5', 'Description 5'),
    (1, '2023-11-16 11:00:00', '2023-11-18 13:00:00', 'Workshop 6', 'Description 6'),
    (1, '2023-11-14 11:00:00', '2023-11-18 13:00:00', 'Workshop 7', 'Description 7'),
    (1, '2023-11-16 11:00:00', '2023-11-16 13:00:00', 'Workshop 7', 'Description 8'),
    (1, '2023-11-15 11:00:00', '2023-11-16 13:00:00', 'Workshop 9', 'Description 9'),
    (2, '2023-11-11 11:00:00', '2023-11-11 13:00:00', 'Workshop 10', 'Description 10'),
    (2, '2023-11-11 11:00:00', '2023-11-11 13:00:00', 'Workshop 11', 'Description 11'),
    (1, '2023-11-19 11:00:00', '2023-11-21 13:00:00', 'Workshop 12', 'Description 12');


INSERT INTO reservations (name, email, workshop_id) VALUES
    ('Faisal Porag', 'faisal.porag@example.com', 1),
    ('Jane Smith1', 'jane.smith1@example.com', 2),
    ('Jane Smith2', 'jane.smith2@example.com', 2),
    ('Jane Smith3', 'jane.smith3@example.com', 2),
    ('Jane Smith4', 'jane.smith4@example.com', 2),
    ('Jane Smith5', 'jane.smith5@example.com', 1),
    ('Khan Smith6', 'khan.smith6@example.com', 1),
    ('Jane Smith7', 'jane.smith7@example.com', 1),
    ('Jon Smith8', 'jon.smith8@example.com', 1),
    ('Jane Smith9', 'jane.smith9@example.com', 2),
    ('Jane Smith10', 'jane.smith10@example.com', 2);



