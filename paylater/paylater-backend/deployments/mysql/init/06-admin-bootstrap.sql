-- Bootstrap first admin (runs once on a fresh MySQL volume, after 02-admin-schema.sql).
-- Login: admin@paylater.local / Admin@123
-- Change this password after first login. Registration of further admins requires an ADMIN JWT.
-- Safe if other admins already exist: only inserts when this email is missing.
USE admin_db;

INSERT INTO admins (name, email, password)
SELECT 'Bootstrap Admin', 'admin@paylater.local', '$2a$10$N8gw78TqcLgcFSUbuuK0ze2ltSilG7gjljPrRezdTQ.F.RHtLeAtW'
WHERE NOT EXISTS (SELECT 1 FROM admins WHERE email = 'admin@paylater.local');
