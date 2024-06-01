-- clear users to allow user_roles and user_status to be cleared
DELETE FROM users;

DELETE FROM user_roles;

DELETE FROM user_status;

DELETE FROM event_categories;

DELETE FROM organizations;
