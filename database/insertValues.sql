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
 ('Élise', 'Dubois', 1, '06 95 36 41 54', '123 Rue de la Paix, Paris', 'elise.dubois@email.com'),
  ('Alexandre', 'Guy', 2, '06 95 34 21 54', '456 Avenue des Champs-Élysées, Paris', 'alex.guy@email.com'),
  ('Sophie', 'Leroux', 3, '06 25 36 41 57', '789 Boulevard Saint-Germain, Paris', 'sophie.leroux@email.com'),
  ('Lucas', 'Dupont', 1, '06 55 36 42 34', '101 Avenue de la République, Lyon', 'lucas.dupont@email.com'),
  ('Manon', 'Laurent', 2, '06 21 85 41 54', '234 Rue du Faubourg, Marseille', 'manon.laurent@email.com'),
  ('Mathieu', 'Lefebvre', 3, '06 95 65 75 01', '456 Avenue de la Liberté, Lyon', 'mathieu.lefebvre@email.com'),
  ('Juliette', 'Fontaine', 1, '06 15 16 41 11', '234 Avenue des Roses, Toulouse', 'juliette.fontaine@email.com'),
  ('Louis', 'Henry', 2, '06 95 00 21 50', '567 Boulevard de Étoile, Marseille', 'louis.henry@email.com'),
  ('Zoé', 'Dufresne', 3, '06 95 30 20 60', '678 Avenue de la Joie, Bordeaux', 'zoe.dufresne@email.com'),
  ('Marie', 'Mone', 1, '06 95 40 20 50', '123 Rue de la Rue', 'marie.mone@email.com'),
('Luc', 'Martin', 2, '06 95 30 70 50', '456 Avenue de Avenue', 'luc.martin@email.com'),
('Sophie', 'Leroy', 3, '06 95 30 25 50', '789 Boulevard du Boulevard', 'sophie.leroy@email.com'),
('Thomas', 'Girard', 1, '06 95 32 20 52', '456 Rue de la Montagne', 'thomas.girard@email.com'),
('Laura', 'Petit', 2, '06 95 38 20 88', '890 Avenue de la Forêt', 'laura.petit@email.com'),
('Pierre', 'Moreau', 3, '06 95 21 21 50', '567 Rue de la Plage', 'pierre.moreau@email.com'),
('Camille', 'Lefevre', 1, '06 95 69 20 69', '123 Avenue du Soleil', 'camille.lefevre@email.com'),
('Nicolas', 'Rousseau', 2, '06 95 23 22 50', '456 Rue de la Lune', 'nicolas.rousseau@email.com'),
('Charlotte', 'Dubois', 3, '06 95 81 82 83', '789 Avenue de l Étoile', 'charlotte.dubois@email.com'),
('Antoine', 'Martin', 1, '06 95 30 28 58', '890 Boulevard de la Mer', 'antoine.martin@email.com');

INSERT INTO employee (firstName, lastName, phone, mail, password, birthDate, hireDate, iban, idRelation, idWorkplace, isWorking, idProfession) VALUES

('John','Doe','06 24 12 36 98','john.doe@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','15-01-1990','11-04-2017','FR761234567RTYNBVG567890123',1,1,1,1),
('Jane','Smith','06 85 69 45 25','jane.smith@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','22-07-1985','15-05-2021','FR89MLO365MLK982013000',2,1,0,2),
('Bob','Johnson','06 58 45 12 36','bob.johnson@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','05-03-1995','17-10-2023','FR29NWBK60161331926819',3,1,1,2),
('Martin','Meunier','06 58 45 12 36','martin.meunier@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','15-07-1994','10-01-2020','FR29NWBDSF523M31926819',4,1,1,3),
('Lea','Come','06 42 45 11 36','lea.come@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','15-07-1994','10-01-2020','FR29LKJUYHGTRFD1926819',5,1,1,4),
('Camille','Lecomte','06 12 34 56 78','camille.lecomte@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','12-05-1990','15-01-2023','FR7612345678901234567890123',6,2,1,5),
('Théo','Moulin','07 23 45 67 89','theo.moulin@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','30-07-1988','22-07-2022','FR89370400440532013000',7,3,1,5),
('Léa','Durand','06 11 22 33 44','lea.durand@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','05-03-1995','18-05-2022','FR3514501234123456789012345',8,4,1,5),
('Louis','Petit','07 12 34 56 78','louis.petit@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','20-10-1987','05-12-2021','FR2812345678901234567890123',9,5,1,5),
('Emma','Moreau','06 54 32 10 98','emma.moreau@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','15-11-1992','10-04-2021','FR7612345678901234567890123',10,6,1,5),
('Hugo','Lefebvre','07 45 67 89 01','hugo.lefebvre@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','30-06-1989','15-11-2020','FR6812345678901234567890123',11,7,1,5),
('Lucas','Dupont','06 12 34 56 78','lucas.dupont@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','12-05-1990','15-01-2023','FR7612345678901234567890123',12,2,1,6),
('Emma','Martin','07 23 45 67 89','emma.martin@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','30-07-1988','22-07-2022','FR89370400440532013000',13,3,1,6),
('Louis','Dubois','06 11 22 33 44','louis.dubois@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','05-03-1995','18-05-2022','FR3514501234123456789012345',14,4,1,6),
('Léa','Lefèvre','07 12 34 56 78','lea.lefevre@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','20-10-1987','18-08-2022','FR2812345678901234567890123',15,5,1,6),
('Hugo','Roux','06 54 32 10 98','hugo.roux@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','15-11-1992','05-12-2021','FR7612345678901234567890123',16,6,1,6),
('Chloé','Girard','07 45 67 89 01','chloe.girard@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','30-06-1989','10-04-2021','FR6812345678901234567890123',17,7,1,6),
('Maxime','Martin','06 12 34 56 78','maxime.martin@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','12-05-1990','15-01-2023','FR7612345678901234567890123',18,2,1,7),
('Manon','Lefebvre','07 23 45 67 89','manon.lefebvre@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','30-07-1988','22-07-2022','FR89370400440532013000',19,2,1,7),
('Antoine','Moreau','06 11 22 33 44','antoine.moreau@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','05-03-1995','18-05-2022','FR3514501234123456789012345',20,3,1,7),
('Inès','Dupont','07 12 34 56 78','ines.dupont@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','20-10-1987','05-12-2021','FR2812345678901234567890123',21,3,0,7),
('Noé','Girard','06 54 32 10 98','noe.girard@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','15-11-1992','10-04-2021','FR7612345678901234567890123',22,4,1,7),
('Zoé','Moulin','07 45 67 89 01','zoe.moulin@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','30-06-1989','15-11-2020','FR6812345678901234567890123',23,4,0,7),
('Louis','Bertrand','06 12 34 56 78','louis.bertrand@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','05-04-1994','15-09-2022','FR7612345678901234567890123',24,5,1,7),
('Léa','Roux','07 23 45 67 89','lea.roux@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','20-01-1993','22-07-2021','FR89370400440532013000',25,5,1,7),
('Ethan','Lecomte','06 11 22 33 44','ethan.lecomte@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','05-11-1991','18-05-2020','FR3514501234123456789012345',26,6,1,7),
('Emma','Berger','07 12 34 56 78','emma.berger@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','20-02-1996','05-12-2019','FR2812345678901234567890123',27,6,1,7),
('Chloé','Lefèvre','06 54 32 10 98','chloe.lefevre@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','15-09-1997','10-04-2019','FR7612345678901234567890123',28,7,1,7),
('Hugo','Robin','07 45 67 89 01','hugo.robin@example.com','154f93d7ee57efef9aecb9b46f179a9aec8b90edb63775a6571c5275f06928a7','30-08-1989','15-11-2018','FR6812345678901234567890123',29,7,0,7);



