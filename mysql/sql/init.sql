CREATE DATABASE test;
use test;
CREATE TABLE users (id int primary key, name varchar(255));
insert into test.users (id, name) VALUES (1, 'taro');
insert into test.users (id, name) VALUES (2, 'jiro');
insert into test.users (id, name) VALUES (3, 'saburo');
insert into test.users (id, name) VALUES (4, 'shiro');
insert into test.users (id, name) VALUES (5, 'goro');

CREATE TABLE `test`.`products` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `product_id` CHAR(8) NOT NULL,
  `product_name` VARCHAR(255) NOT NULL,
  `price` DECIMAL NOT NULL,
  `register_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `product_id_UNIQUE` (`product_id` ASC));

CREATE TABLE IF NOT EXISTS `test`.`bigtable` (
 `col1` INT NULL,`col2` INT NULL,`col3` INT NULL,`col4` INT NULL,`col5` INT NULL
 ,`col6` INT NULL,`col7` INT NULL,`col8` INT NULL,`col9` INT NULL,`col10` INT NULL
 ,`col11` INT NULL,`col12` INT NULL,`col13` INT NULL,`col14` INT NULL,`col15` INT NULL
 ,`col16` INT NULL,`col17` INT NULL,`col18` INT NULL,`col19` INT NULL,`col20` INT NULL
 ,`col21` INT NULL,`col22` INT NULL,`col23` INT NULL,`col24` INT NULL,`col25` INT NULL
 ,`col26` INT NULL,`col27` INT NULL,`col28` INT NULL,`col29` INT NULL,`col30` INT NULL
 ,`col31` INT NULL,`col32` INT NULL,`col33` INT NULL,`col34` INT NULL,`col35` INT NULL
 ,`col36` INT NULL,`col37` INT NULL,`col38` INT NULL,`col39` INT NULL,`col40` INT NULL
 ,`col41` INT NULL,`col42` INT NULL,`col43` INT NULL,`col44` INT NULL,`col45` INT NULL
 ,`col46` INT NULL,`col47` INT NULL,`col48` INT NULL,`col49` INT NULL,`col50` INT NULL
 ,`col51` INT NULL,`col52` INT NULL,`col53` INT NULL,`col54` INT NULL,`col55` INT NULL
 ,`col56` INT NULL,`col57` INT NULL,`col58` INT NULL,`col59` INT NULL,`col60` INT NULL
 ,`col61` INT NULL,`col62` INT NULL,`col63` INT NULL,`col64` INT NULL,`col65` INT NULL
 ,`col66` INT NULL,`col67` INT NULL,`col68` INT NULL,`col69` INT NULL,`col70` INT NULL
 ,`col71` INT NULL,`col72` INT NULL,`col73` INT NULL,`col74` INT NULL,`col75` INT NULL
 ,`col76` INT NULL,`col77` INT NULL,`col78` INT NULL,`col79` INT NULL,`col80` INT NULL
 ,`col81` INT NULL,`col82` INT NULL,`col83` INT NULL,`col84` INT NULL,`col85` INT NULL
 ,`col86` INT NULL,`col87` INT NULL,`col88` INT NULL,`col89` INT NULL,`col90` INT NULL
 ,`col91` INT NULL,`col92` INT NULL,`col93` INT NULL,`col94` INT NULL,`col95` INT NULL
 ,`col96` INT NULL,`col97` INT NULL,`col98` INT NULL,`col99` INT NULL,`col100` INT NULL
);

