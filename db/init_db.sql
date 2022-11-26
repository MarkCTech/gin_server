DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks(
  id         INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  completed BOOLEAN
);

INSERT INTO tasks
  (title, completed)
VALUES
  ('Clean Room', false),
  ('Read Book', false),
  ('Sleep', true);