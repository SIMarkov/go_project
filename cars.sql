CREATE DATABASE cars;

CREATE USER 'root'@'localhost' IDENTIFIED BY 'Slavil1991';
GRANT ALL PRIVILEGES ON cars.* TO 'root'@'localhost';
FLUSH PRIVILEGES;