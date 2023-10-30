INSERT INTO workplace (name, adress, phone, mail, type) VALUES 
("Paris","25 Avenue Foch","01 69 25 30 10","paris@decantez-vous.com","head office"),
("Nantes","20 Rue Crébillon","02 40 41 90 80","nantes@decantez-vous.com","store"),
("Marseille","15 Avenue Robert Schuman","04 40 31 92 80","marseille@decantez-vous.com","store"),
("Lille","78 Rue de la Barre","03 20 31 92 80","lille@decantez-vous.com","store"),
("Bordeaux","78 Rue de la Barre","05 30 10 06 80","bordeaux@decantez-vous.com","store"),
("Strasbourg","78 Rue de la Barre","03 20 81 56 85","strasbourg@decantez-vous.com","store"),
("Lyon","78 Rue de la Barre","05 20 31 52 80","lyon@decantez-vous.com","store");

INSERT INTO profession (name, description, salary) VALUES 
("Director", "directs the company's actions", 8000),
("RH", "improve working conditions and organization", 3500),
("Community Marketing", "create adds for the company", 2200),
("Commercial", "search product et supplier", 2800),
("Site Manager", "directs the site", 3200),
("Team Leader", "directs the vendors", 2500),
("Vendor", "sell the product", 1800);

INSERT INTO product (name) VALUES 
("vin"),
("rhum"),
("whisky");

INSERT INTO supplier (firstName, lastName, idProduct, phone, adress, mail) VALUES 
('Élise', 'Dubois', 1, '06 95 36 41 54', '123 Rue de la Paix, Paris', 'elise.dubois@decantez-vous.com'),
('Alexandre', 'Guy', 2, '06 95 34 21 54', '456 Avenue des Champs-Élysées, Paris', 'alex.guy@decantez-vous.com'),
('Sophie', 'Leroux', 3, '06 25 36 41 57', '789 Boulevard Saint-Germain, Paris', 'sophie.leroux@decantez-vous.com'),
('Lucas', 'Dupont', 1, '06 55 36 42 34', '101 Avenue de la République, Lyon', 'lucas.dupont@decantez-vous.com'),
('Manon', 'Laurent', 2, '06 21 85 41 54', '234 Rue du Faubourg, Marseille', 'manon.laurent@decantez-vous.com'),
('Mathieu', 'Lefebvre', 3, '06 95 65 75 01', '456 Avenue de la Liberté, Lyon', 'mathieu.lefebvre@decantez-vous.com'),
('Juliette', 'Fontaine', 1, '06 15 16 41 11', '234 Avenue des Roses, Toulouse', 'juliette.fontaine@decantez-vous.com'),
('Louis', 'Henry', 2, '06 95 00 21 50', '567 Boulevard de Étoile, Marseille', 'louis.henry@decantez-vous.com'),
('Zoé', 'Dufresne', 3, '06 95 30 20 60', '678 Avenue de la Joie, Bordeaux', 'zoe.dufresne@decantez-vous.com'),
('Marie', 'Mone', 1, '06 95 40 20 50', '123 Rue de la Rue', 'marie.mone@decantez-vous.com'),
('Luc', 'Martin', 2, '06 95 30 70 50', '456 Avenue de Avenue', 'luc.martin@decantez-vous.com'),
('Sophie', 'Leroy', 3, '06 95 30 25 50', '789 Boulevard du Boulevard', 'sophie.leroy@decantez-vous.com'),
('Thomas', 'Girard', 1, '06 95 32 20 52', '456 Rue de la Montagne', 'thomas.girard@decantez-vous.com'),
('Laura', 'Petit', 2, '06 95 38 20 88', '890 Avenue de la Forêt', 'laura.petit@decantez-vous.com'),
('Pierre', 'Moreau', 3, '06 95 21 21 50', '567 Rue de la Plage', 'pierre.moreau@decantez-vous.com'),
('Camille', 'Lefevre', 1, '06 95 69 20 69', '123 Avenue du Soleil', 'camille.lefevre@decantez-vous.com'),
('Nicolas', 'Rousseau', 2, '06 95 23 22 50', '456 Rue de la Lune', 'nicolas.rousseau@decantez-vous.com'),
('Charlotte', 'Dubois', 3, '06 95 81 82 83', '789 Avenue de l Étoile', 'charlotte.dubois@decantez-vous.com'),
('Antoine', 'Martin', 1, '06 95 30 28 58', '890 Boulevard de la Mer', 'antoine.martin@decantez-vous.com');

INSERT INTO employee (firstName, lastName, phone, mail, password, birthDate, hireDate, iban, idWorkplace, isWorking, idProfession) VALUES
('John','Doe','06 24 12 36 98','john.doe@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1990-01-15','2017-04-11','FR761234567RTYNBVG567890123',1,1,1),
('Jane','Smith','06 85 69 45 25','jane.smith@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1985-07-22','2021-05-15','FR89MLO365MLK982013000',1,0,2),
('Bob','Johnson','06 58 45 12 36','bob.johnson@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1995-03-05','2023-10-17','FR29NWBK60161331926819',1,1,2),
('Martin','Meunier','06 58 45 12 36','martin.meunier@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','15-07-1994','2020-01-10','FR29NWBDSF523M31926819',1,1,3),
('Lea','Come','06 42 45 11 36','lea.come@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1994-07-15','2020-01-10','FR29LKJUYHGTRFD1926819',1,1,4),
('Camille','Lecomte','06 12 34 56 78','camille.lecomte@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1990-05-12','2023-01-15','FR7612345678901234567890123',2,1,5),
('Théo','Moulin','07 23 45 67 89','theo.moulin@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1988-07-30','2022-07-22','FR89370400440532013000',3,1,5),
('Léa','Durand','06 11 22 33 44','lea.durand@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1995-03-05','2022-05-18','FR3514501234123456789012345',4,1,5),
('Louis','Petit','07 12 34 56 78','louis.petit@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1987-10-20','2021-12-05','FR2812345678901234567890123',5,1,5),
('Emma','Moreau','06 54 32 10 98','emma.moreau@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1992-11-15','2021-04-10','FR7612345678901234567890123',6,1,5),
('Hugo','Lefebvre','07 45 67 89 01','hugo.lefebvre@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1989-06-30','2020-11-15','FR6812345678901234567890123',7,1,5),
('Lucas','Dupont','06 12 34 56 78','lucas.dupont@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1990-05-12','2023-01-15','FR7612345678901234567890123',2,1,6),
('Emma','Martin','07 23 45 67 89','emma.martin@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1988-07-30','2022-07-22','FR89370400440532013000',3,1,6),
('Louis','Dubois','06 11 22 33 44','louis.dubois@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1995-03-05','2022-05-18','FR3514501234123456789012345',4,1,6),
('Léa','Lefèvre','07 12 34 56 78','lea.lefevre@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1987-10-20','2022-08-18','FR2812345678901234567890123',5,1,6),
('Hugo','Roux','06 54 32 10 98','hugo.roux@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1992-11-15','2021-12-05','FR7612345678901234567890123',6,1,6),
('Chloé','Girard','07 45 67 89 01','chloe.girard@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1989-06-30','2021-04-10','FR6812345678901234567890123',7,1,6),
('Maxime','Martin','06 12 34 56 78','maxime.martin@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1990-05-12','2023-01-15','FR7612345678901234567890123',2,1,7),
('Manon','Lefebvre','07 23 45 67 89','manon.lefebvre@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1988-07-30','2022-07-22','FR89370400440532013000',2,1,7),
('Antoine','Moreau','06 11 22 33 44','antoine.moreau@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','05-03-1995','2022-05-18','FR3514501234123456789012345',3,1,7),
('Inès','Dupont','07 12 34 56 78','ines.dupont@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1987-10-20','2021-12-05','FR2812345678901234567890123',3,0,7),
('Noé','Girard','06 54 32 10 98','noe.girard@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1992-11-15','2021-04-10','FR7612345678901234567890123',4,1,7),
('Zoé','Moulin','07 45 67 89 01','zoe.moulin@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1989-06-30','2020-11-15','FR6812345678901234567890123',4,0,7),
('Louis','Bertrand','06 12 34 56 78','louis.bertrand@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1994-04-05','2022-09-15','FR7612345678901234567890123',5,1,7),
('Léa','Roux','07 23 45 67 89','lea.roux@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1993-01-20','2021-07-22','FR89370400440532013000',5,1,7),
('Ethan','Lecomte','06 11 22 33 44','ethan.lecomte@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1991-11-05','2020-05-18','FR3514501234123456789012345',6,1,7),
('Emma','Berger','07 12 34 56 78','emma.berger@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1996-02-20','2019-12-05-','FR2812345678901234567890123',6,1,7),
('Chloé','Lefèvre','06 54 32 10 98','chloe.lefevre@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1997-09-15','2019-04-10','FR7612345678901234567890123',7,1,7),
('Hugo','Robin','07 45 67 89 01','hugo.robin@decantez-vous.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','1989-08-30','2018-11-15','FR6812345678901234567890123',7,0,7);


INSERT INTO relationWorkplaceSupplier (idSupplier, idWorkplace ) VALUES
(1,2),(1,4),(1,5),(2,2),(3,4),(3,5),(4,6),(4,7),(5,3),(5,6),(6,2),(7,7),(8,5),(9,4),
(9,5),(10,3),(11,6),(12,6),(12,7),(13,3),(14,5),(15,2),(15,7),(16,4),(17,4),(18,2),(19,5);

INSERT INTO relationEmployee (idEmployee, idReferent) VALUES 
(1,1),(2,1),(3,1),(4,1),(5,1),(6,1),(7,1),(8,1),(9,1),(10,1),(11,1),
(12,6),(13,7),(14,8),(15,9),(16,10),(17,11),(18,12),(19,12),(20,13),(21,13),
(22,14),(23,14),(24,15),(15,25),(16,26),(27,16),(28,18);