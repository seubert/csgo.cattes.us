CREATE TABLE `music` ( `id` int(11) unsigned NOT NULL AUTO_INCREMENT, `file_name` varchar(255) NOT NULL DEFAULT '', `uploader` varchar(255) NOT NULL DEFAULT '', `artist` varchar(255) DEFAULT NULL, `title` varchar(255) DEFAULT NULL, `genre` varchar(255) DEFAULT NULL, PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8;