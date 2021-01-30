
-- 0 Success
-- 1 Fatal Error
-- 2 Connection or Server Bad
-- 3 Error occured in Script

DROP TABLE IF EXISTS loginuser;

CREATE TABLE IF NOT EXISTS loginuser 
( 
  id SERIAL PRIMARY KEY,
  first_name varchar(20) NOT NULL,
  last_name varchar(20) NOT NULL,
  email varchar(20) NOT NULL,
  password varchar(100) NOT NULL,
  isAdmin BOOLEAN DEFAULT FALSE,
  address1 varchar(20) DEFAULT NULL,
  address2 varchar(100) DEFAULT NULL,
  address3 varchar(100) DEFAULT NULL,
  phone_no varchar(100) DEFAULT NULL,
  create_user varchar(100) NOT NULL,
  create_datetime timestamp NOT NULL,
  update_user varchar(100) NOT NULL,
  update_datetime timestamp NOT NULL
);

TRUNCATE TABLE loginuser RESTART IDENTITY;

INSERT INTO loginuser VALUES
(0, 'Nayret', 'John', 'nay@gmail.com', '123', true, '1770044', '東京都練馬区', '上石神井1丁目2-24-8-810', '08077779999', 'system', NOW(), 'system', NOW()),
(1, 'Daniel', 'John', 'daniel@gmail.com', '123', false,  '1770044', '東京都練馬区', '上石神井1丁目2-24-8-810', '08077779999', 'system', NOW(), 'system', NOW());