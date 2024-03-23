CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100),
  email VARCHAR(100),
  password VARCHAR(100)
);


INSERT INTO users (name, email, password) VALUES
('Mateus', 'mateus@jesus.com', '123456'),
('João', 'joao@jesus.com', '123456'),
('Maria', 'maria@jesus.com', '123456');
