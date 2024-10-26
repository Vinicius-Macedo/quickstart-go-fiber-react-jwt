-- Tabela users
CREATE TABLE
  IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );

-- Tabela images
CREATE TABLE
  IF NOT EXISTS images (
    id SERIAL PRIMARY KEY,
    user_email TEXT NOT NULL UNIQUE,
    image_url TEXT NOT NULL,
    uploaded_at TIMESTAMP
    WITH
      TIME ZONE DEFAULT CURRENT_TIMESTAMP,
      CONSTRAINT fk_user_email FOREIGN KEY (user_email) REFERENCES users (email) -- Definição explícita da FK
  );