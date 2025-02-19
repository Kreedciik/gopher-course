CREATE TYPE reservation_status AS ENUM ('pending', 'confirmed', 'canceled');

CREATE TABLE restaurants (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    address TEXT NOT NULL,
    phone_number VARCHAR NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE reservations (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    restaurant_id UUID REFERENCES restaurants(id),
    reservation_time TIMESTAMP NOT NULL,
    status reservation_status NOT NULL
);

CREATE TABLE menus (
    id UUID PRIMARY KEY,
    restaurant_id UUID REFERENCES restaurants(id),
    name VARCHAR NOT NULL,
    description TEXT NOT NULL,
    price DOUBLE PRECISION NOT NULL
);

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    reservation_id UUID REFERENCES reservations(id),
    menu_item_id UUID REFERENCES menus(id),
    quantity INTEGER NOT NULL
);