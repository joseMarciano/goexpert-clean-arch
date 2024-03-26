CREATE DATABASE IF NOT EXISTS orders;
USE orders;


CREATE TABLE IF NOT EXISTS orders.orders (
                        id VARCHAR(50) PRIMARY KEY,
                        price DOUBLE,
                        tax DOUBLE,
                        final_price DOUBLE
);