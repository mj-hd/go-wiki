CREATE TABLE `pages` (
	  `id` int(11) NOT NULL AUTO_INCREMENT,
	  `title` varchar(128) NOT NULL,
	  `user` varchar(64) NOT NULL,
	  `locked` tinyint(1) NOT NULL DEFAULT '0',
	  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	  `modified` timestamp NULL DEFAULT NULL,
	  `content` longtext,
	  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;

CREATE TABLE `users` (
	  `id` int(11) NOT NULL AUTO_INCREMENT,
	  `name` varchar(64) NOT NULL,
	  `caption` varchar(128) NOT NULL DEFAULT '',
	  `level` int(11) NOT NULL DEFAULT '0',
	  `registered` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	  `address` varchar(128) NOT NULL DEFAULT '',
	  `password` varchar(60) DEFAULT '',
	  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
