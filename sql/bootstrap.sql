/*this sql file will be mapped into the init.d of the 
mysql server so it creates the tables as appropriate*/
CREATE DATABASE planet;
USE planet;

CREATE TABLE planet_users(
    user_id VARCHAR(50) NOT NULL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL
);

CREATE TABLE planet_groups(
    group_name VARCHAR(50) NOT NULL PRIMARY KEY
);

CREATE TABLE plant_group_members(
    user_id VARCHAR(50) NOT NULL,
    group_id VARCHAR(50) NOT NULL,
    FOREIGN KEY(user_id) REFERENCES planet_users(user_id),
    FOREIGN KEY(group_id) REFERENCES planet_groups(group_name),
    PRIMARY KEY pk_group_members (user_id, group_id)
);

INSERT INTO planet_groups (group_name) VALUES ("Admin");
INSERT INTO planet_users (first_name, last_name, user_id) VALUES ("Casey", "Manus", "cmanus");
INSERT INTO plant_group_members(user_id, group_id) VALUES("cmanus","Admin");
GRANT ALL on planet.* to app;