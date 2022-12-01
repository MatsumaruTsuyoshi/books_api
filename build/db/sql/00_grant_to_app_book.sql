CREATE USER 'todo-app'@'%' IDENTIFIED BY 'todo-password';
GRANT SELECT,INSERT,UPDATE,DELETE ON book.* TO 'todo-app'@'%';