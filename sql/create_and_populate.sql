DROP DATABASE IF EXISTS genealogy;

CREATE DATABASE genealogy;

USE genealogy;

DROP TABLE IF EXISTS persons;

CREATE TABLE `persons`
(
    `id`   int         NOT NULL AUTO_INCREMENT,
    `name` varchar(256) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO genealogy.persons (id, name)
VALUES (1, 'Bruce');
INSERT INTO genealogy.persons (id, name)
VALUES (2, 'Mike');
INSERT INTO genealogy.persons (id, name)
VALUES (3, 'Sonny');
INSERT INTO genealogy.persons (id, name)
VALUES (4, 'Phoebe');
INSERT INTO genealogy.persons (id, name)
VALUES (5, 'Anastasia');
INSERT INTO genealogy.persons (id, name)
VALUES (6, 'Martin');
INSERT INTO genealogy.persons (id, name)
VALUES (7, 'Dunny');
INSERT INTO genealogy.persons (id, name)
VALUES (8, 'Ursula');
INSERT INTO genealogy.persons (id, name)
VALUES (9, 'Jacqueline');

DROP TABLE IF EXISTS relationships;

CREATE TABLE `relationships`
(
    `parent`   int NOT NULL,
    `children` int NOT NULL,
    PRIMARY KEY (`parent`, `children`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO genealogy.relationships (parent, children)
VALUES (2, 1);
INSERT INTO genealogy.relationships (parent, children)
VALUES (2, 7);
INSERT INTO genealogy.relationships (parent, children)
VALUES (3, 2);
INSERT INTO genealogy.relationships (parent, children)
VALUES (4, 1);
INSERT INTO genealogy.relationships (parent, children)
VALUES (4, 7);
INSERT INTO genealogy.relationships (parent, children)
VALUES (5, 4);
INSERT INTO genealogy.relationships (parent, children)
VALUES (5, 8);
INSERT INTO genealogy.relationships (parent, children)
VALUES (6, 4);
INSERT INTO genealogy.relationships (parent, children)
VALUES (6, 8);
INSERT INTO genealogy.relationships (parent, children)
VALUES (8, 9);