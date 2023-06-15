CREATE DATABASE gorest;
USE gorest;
CREATE TABLE Customer (
  Customer_ID INT AUTO_INCREMENT PRIMARY KEY,
  Name VARCHAR(100) NOT NULL,
  Email VARCHAR(100),
  Phone_Number VARCHAR(10)
);

CREATE TABLE Menu (
  Item_ID INT AUTO_INCREMENT PRIMARY KEY,
  Name VARCHAR(100) NOT NULL,
  Description VARCHAR(100),
  Price INT NOT NULL
);

CREATE TABLE Orders (
  Order_ID INT AUTO_INCREMENT PRIMARY KEY,
  Customer_ID INT NOT NULL,
  Order_Time DATETIME NOT NULL,
  Price_Total INT,
  FOREIGN KEY (Customer_ID) REFERENCES Customer(Customer_ID)
);

CREATE TABLE Order_Item (
  Order_ID INT NOT NULL,
  Item_ID INT NOT NULL,
  Quantity INT NOT NULL,
  Price_Unit INT NOT NULL,
  FOREIGN KEY (Order_ID) REFERENCES Orders(Order_ID),
  FOREIGN KEY (Item_ID) REFERENCES Menu(Item_ID),
  PRIMARY KEY (Order_ID, Item_ID)
);

INSERT INTO Customer (Name, Email, Phone_Number)
VALUES  ('John Doe', 'johndoe@example.com', '1234567890'),
        ('Jane Smith', 'janesmith@example.com', '9876543210'),
        ('Michael Johnson', 'michaeljohnson@example.com', '5555555555');

INSERT INTO Menu (Name, Description, Price)
VALUES  ('Cheeseburger', 'Delicious beef patty topped with melted cheese', 10),
        ('Chicken Sandwich', 'Grilled chicken breast with fresh veggies', 8),
        ('Margherita Pizza', 'Classic pizza with tomato sauce and mozzarella cheese', 12),
        ('French Fries', 'Crispy golden fries', 5),
        ('Caesar Salad', 'Fresh lettuce, croutons, and Parmesan cheese', 7);

INSERT INTO Orders (Customer_ID, Order_Time, Price_Total)
VALUES  (1, NOW(), 27),
        (2, NOW(), 15),
        (3, NOW(), 22);

INSERT INTO Order_Item (Order_ID, Item_ID, Quantity, Price_Unit)
VALUES  (1, 1, 2, 10),
        (1, 4, 1, 5),
        (2, 2, 1, 8),
        (2, 3, 1, 12),
        (3, 1, 1, 10),
        (3, 5, 2, 7);

