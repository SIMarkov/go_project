CREATE DATABASE cars;

CREATE USER 'root'@'localhost' IDENTIFIED BY 'SM';
GRANT ALL PRIVILEGES ON cars.* TO 'root'@'localhost';
FLUSH PRIVILEGES;
