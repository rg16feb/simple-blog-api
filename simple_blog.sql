-- Use the below code to create a database
CREATE DATABASE simple_blog;
USE simple_blog;

CREATE TABLE blog_posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL
);
