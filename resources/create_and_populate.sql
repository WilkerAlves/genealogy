CREATE TABLE `persons`
(
    `id`   int         NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb3;

INSERT INTO genealogy.persons (id, name)
VALUES (5, 'Bruce');
INSERT INTO genealogy.persons (id, name)
VALUES (6, 'Mike');
INSERT INTO genealogy.persons (id, name)
VALUES (7, 'Sonny');
INSERT INTO genealogy.persons (id, name)
VALUES (8, 'Phoebe');
INSERT INTO genealogy.persons (id, name)
VALUES (9, 'Anastasia');
INSERT INTO genealogy.persons (id, name)
VALUES (10, 'Martin');
INSERT INTO genealogy.persons (id, name)
VALUES (11, 'Dunny');
INSERT INTO genealogy.persons (id, name)
VALUES (12, 'Úrsula');
INSERT INTO genealogy.persons (id, name)
VALUES (13, 'Jacqueline');


CREATE TABLE `relationships`
(
    `parent`   int NOT NULL,
    `children` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO genealogy.relationships (parent, children)
VALUES (6, 5);
INSERT INTO genealogy.relationships (parent, children)
VALUES (6, 11);
INSERT INTO genealogy.relationships (parent, children)
VALUES (7, 6);
INSERT INTO genealogy.relationships (parent, children)
VALUES (8, 5);
INSERT INTO genealogy.relationships (parent, children)
VALUES (8, 11);
INSERT INTO genealogy.relationships (parent, children)
VALUES (9, 8);
INSERT INTO genealogy.relationships (parent, children)
VALUES (9, 12);
INSERT INTO genealogy.relationships (parent, children)
VALUES (10, 8);
INSERT INTO genealogy.relationships (parent, children)
VALUES (10, 12);
INSERT INTO genealogy.relationships (parent, children)
VALUES (12, 13);