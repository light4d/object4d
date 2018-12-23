
CREATE SCHEMA `object4d`;
use object4d;

CREATE TABLE `object4d` (
  `lng` varchar(45) NOT NULL,
  `lat` varchar(45) NOT NULL,
  `h` varchar(45) NOT NULL DEFAULT 'R',
  `t` varchar(45) NOT NULL,
  `m` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `miniocon` (
  `id` int(11) NOT NULL,
  `endpoint` varchar(150) NOT NULL,
  `ak` varchar(250) NOT NULL,
  `sk` varchar(45) NOT NULL,
  `secure` tinyint(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
