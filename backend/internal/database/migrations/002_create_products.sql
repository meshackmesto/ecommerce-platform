-- 002_create_products.sql
CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    stock_quantity INTEGER NOT NULL DEFAULT 0 CHECK (stock_quantity >= 0),
    category_id UUID REFERENCES categories(id) ON DELETE SET NULL,
    image_url VARCHAR(500),
    sku VARCHAR(100) UNIQUE,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for better performance
CREATE INDEX idx_products_category ON products(category_id);
CREATE INDEX idx_products_active ON products(is_active);
CREATE INDEX idx_products_name ON products(name);

-- Insert sample categories
INSERT INTO categories (name, description) VALUES 
('Electronics', 'Electronic devices and gadgets'),
('Clothing', 'Fashion and apparel'),
('Books', 'Books and educational materials'),
('Home & Garden', 'Home improvement and garden supplies');

-- Insert sample products
INSERT INTO products (name, description, price, stock_quantity, sku, image_url, category_id) VALUES 
('iPhone 14', 'Latest Apple smartphone with advanced features', 999.99, 50, 'IPHONE14-001', 'https://images.unsplash.com/photo-1592750475338-74b7b21085ab?w=300', 
 (SELECT id FROM categories WHERE name = 'Electronics')),
('MacBook Pro', 'High-performance laptop for professionals', 1299.99, 25, 'MBP-001', 'https://images.unsplash.com/photo-1496181133206-80ce9b88a853?w=300',
 (SELECT id FROM categories WHERE name = 'Electronics')),
('Cotton T-Shirt', 'Comfortable 100% cotton t-shirt', 29.99, 100, 'TSHIRT-001', 'https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=300',
 (SELECT id FROM categories WHERE name = 'Clothing')),
('Programming Book', 'Learn modern web development', 49.99, 30, 'BOOK-001', 'https://images.unsplash.com/photo-1544716278-ca5e3f4abd8c?w=300',
 (SELECT id FROM categories WHERE name = 'Books'));
