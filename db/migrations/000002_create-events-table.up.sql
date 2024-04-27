CREATE TABLE IF NOT EXISTS organizations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL DEFAULT 'self'
);

CREATE TABLE IF NOT EXISTS images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    filename VARCHAR(255),
    url VARCHAR(255),
    alt VARCHAR(150),
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS locations (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL DEFAULT 'TBD',
    url VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS event_more_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    info VARCHAR(1500),
    url VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS events (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(1500),
    organization_id INT NOT NULL,
    img_id INT,
    location_id INT,
    price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    category_id INT NOT NULL,
    more_info_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (organization_id) REFERENCES organizations(id),
    FOREIGN KEY (img_id) REFERENCES images(id),
    FOREIGN KEY (location_id) REFERENCES locations(id),
    FOREIGN KEY (more_info_id) REFERENCES event_more_info(id),
    FOREIGN KEY (category_id) REFERENCES event_categories(id)
);
