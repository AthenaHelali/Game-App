CREATE TABLE users (
    id int not null primary key AUTO_INCREMENT,
    name varchar(255) not null ,
    phone_number varchar(255) not null unique,
    password text not null ,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);