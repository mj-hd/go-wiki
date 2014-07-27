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

INSERT INTO `pages` (`id`, `title`, `user`, `locked`, `created`, `modified`, `content`)
VALUES
	(1, 'index', 'anonymous', 0, '2014-07-22 20:48:49', '2014-07-26 10:12:55', 'An h1 header\r\n============\r\n\r\nParagraphs are separated by a blank line.\r\n\r\n2nd paragraph. *Italic*, **bold**, and `monospace`. Itemized lists\r\nlook like:\r\n\r\n  * this one\r\n  * that one\r\n  * the other one\r\n\r\nNote that --- not considering the asterisk --- the actual text\r\ncontent starts at 4-columns in.\r\n\r\n> Block quotes are\r\n> written like so.\r\n>\r\n> They can span multiple paragraphs,\r\n> if you like.\r\n\r\nUse 3 dashes for an em-dash. Use 2 dashes for ranges (ex., \"it\'s all\r\nin chapters 12--14\"). Three dots ... will be converted to an ellipsis.\r\n\r\n\r\n\r\nAn h2 header\r\n------------\r\n\r\nHere\'s a numbered list:\r\n\r\n 1. first item\r\n 2. second item\r\n 3. third item\r\n\r\nNote again how the actual text starts at 4 columns in (4 characters\r\nfrom the left side). Here\'s a code sample:\r\n\r\n    # Let me re-iterate ...\r\n    for i in 1 .. 10 { do-something(i) }\r\n\r\nAs you probably guessed, indented 4 spaces. By the way, instead of\r\nindenting the block, you can use delimited blocks, if you like:\r\n\r\n~~~\r\ndefine foobar() {\r\n    print \"Welcome to flavor country!\";\r\n}\r\n~~~\r\n\r\n(which makes copying & pasting easier). You can optionally mark the\r\ndelimited block for Pandoc to syntax highlight it:\r\n\r\n~~~python\r\nimport time\r\n# Quick, count to ten!\r\nfor i in range(10):\r\n    # (but not *too* quick)\r\n    time.sleep(0.5)\r\n    print i\r\n~~~\r\n\r\n\r\n\r\n### An h3 header ###\r\n\r\nNow a nested list:\r\n\r\n 1. First, get these ingredients:\r\n\r\n      * carrots\r\n      * celery\r\n      * lentils\r\n\r\n 2. Boil some water.\r\n\r\n 3. Dump everything in the pot and follow\r\n    this algorithm:\r\n\r\n        find wooden spoon\r\n        uncover pot\r\n        stir\r\n        cover pot\r\n        balance wooden spoon precariously on pot handle\r\n        wait 10 minutes\r\n        goto first step (or shut off burner when done)\r\n\r\n    Do not bump wooden spoon or it will fall.\r\n\r\nNotice again how text always lines up on 4-space indents (including\r\nthat last line which continues item 3 above). Here\'s a link to [a\r\nwebsite](http://foo.bar). Here\'s a link to a [local\r\ndoc](local-doc.html). Here\'s a footnote [^1].\r\n\r\n[^1]: Footnote text goes here.\r\n\r\nTables can look like this:\r\n\r\nsize  material      color\r\n----  ------------  ------------\r\n9     leather       brown\r\n10    hemp canvas   natural\r\n11    glass         transparent\r\n\r\nTable: Shoes, their sizes, and what they\'re made of\r\n\r\n(The above is the caption for the table.) Pandoc also supports\r\nmulti-line tables:\r\n\r\n--------  -----------------------\r\nkeyword   text\r\n--------  -----------------------\r\nred       Sunsets, apples, and\r\n          other red or reddish\r\n          things.\r\n\r\ngreen     Leaves, grass, frogs\r\n          and other things it\'s\r\n          not easy being.\r\n--------  -----------------------\r\n\r\nA horizontal rule follows.\r\n\r\n***\r\n\r\nHere\'s a definition list:\r\n\r\napples\r\n  : Good for making applesauce.\r\noranges\r\n  : Citrus!\r\ntomatoes\r\n  : There\'s no \"e\" in tomatoe.\r\n\r\nAgain, text is indented 4 spaces. (Alternately, put blank lines in\r\nbetween each of the above definition list lines to spread things\r\nout more.)\r\n\r\nHere\'s a \"line block\":\r\n\r\n| Line one\r\n|   Line too\r\n| Line tree\r\n\r\nand images can be specified like so:\r\n\r\n![example image](example-image.jpg \"An exemplary image\")\r\n\r\nInline math equations go in like so: $\\omega = d\\phi / dt$. Display\r\nmath should get its own line and be put in in double-dollarsigns:\r\n\r\n$$I = \\int \\rho R^{2} dV$$\r\n\r\nAnd note that you can backslash-escape any punctuation characters\r\nwhich you wish to be displayed literally, ex.: \\`foo\\`, \\*bar\\*, etc.\r\n\r\nDone.\r\n');

INSERT INTO `pages` (`id`, `title`, `user`, `locked`, `created`, `modified`, `content`)
VALUES
	(16, '_sidebar', 'anonymous', 0, '2014-07-23 19:40:17', '2014-07-26 13:12:59', '* [Home](/)\r\n');


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
