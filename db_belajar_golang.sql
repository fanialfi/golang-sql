CREATE DATABASE IF NOT EXISTS `db_belajar_golang`;

USE DATABASE `db_belajar_golang`;

CREATE TABLE IF NOT EXISTS `tb_student` (
  `id` varchar(5) NOT NULL,
  `name` varchar(255) NOT NULL,
  `age` int(11) NOT NULL,
  `grade` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `tb_student` (`id`,`name`,`age`,`grade`) VALUES
('B001','fani alfirdaus',17,1),
('B002','Jason Burne',18,2),
('B003','James Bond',19,3),
('B004','Ethan Hunt',20,4),
('B005','John Wick',21,5);

ALTER TABLE `tb_student` ADD PRIMARY KEY(`id`);