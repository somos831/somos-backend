INSERT INTO user_roles (id, role_name) 
VALUES 
(1, "administrator"), (2, "editor"), (3, "member"), (4, "developer"), (5, "guest");

INSERT INTO user_status (id, status_name) 
VALUES 
(1, "active"), (2, "inactive"), (3, "locked"), (4, "pending"), (5, "suspended"), (6, "terminated"), (7, "banned");

INSERT INTO event_categories (id, name) 
VALUES 
(1, "networking"), (2, "conference"), (3, "seminar"), (4, "job fair"), (5, "community"), (6, "fundraiser"), (7, "virtual"), 
(8, "training"), (9, "recreational"), (10, "tradeshow"), (11, "other");

INSERT INTO organizations (id, name) 
VALUES 
(1, "SOMOS");
