CREATE TYPE user_role AS ENUM ('user', 'organization', 'admin');

ALTER TABLE users
ADD COLUMN role user_role DEFAULT 'user';
