CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(150) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(50),
    password_hash TEXT NOT NULL,
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);

CREATE TABLE properties (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    owner_id UUID NOT NULL REFERENCES users(id),
    category_id UUID NOT NULL REFERENCES categories(id),

    title VARCHAR(255) NOT NULL,
    description TEXT,

    transaction_type VARCHAR(20),

    price NUMERIC(15,2),

    bedrooms INT,
    bathrooms INT,

    city VARCHAR(100),

    status VARCHAR(20) DEFAULT 'pending',

    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE images (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    property_id UUID REFERENCES properties(id) ON DELETE CASCADE,
    image_url TEXT NOT NULL
);

CREATE TABLE contacts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    property_id UUID REFERENCES properties(id),
    client_id UUID REFERENCES users(id),
    message TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE visits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    property_id UUID REFERENCES properties(id),
    client_id UUID REFERENCES users(id),

    visit_date TIMESTAMP,
    status VARCHAR(20) DEFAULT 'pending'
);

CREATE TABLE deals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    property_id UUID REFERENCES properties(id),
    client_id UUID REFERENCES users(id),

    final_price NUMERIC(15,2),

    status VARCHAR(20) DEFAULT 'closed'
);

CREATE TABLE commissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    deal_id UUID UNIQUE REFERENCES deals(id),

    percentage NUMERIC(5,2),

    amount NUMERIC(15,2)
);

CREATE TABLE logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID REFERENCES users(id),

    action VARCHAR(255),

    created_at TIMESTAMP DEFAULT NOW()
);