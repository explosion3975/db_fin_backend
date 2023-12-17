CREATE TABLE customer_info(
    id    CHAR(10),
    name  CHAR(12),
    phone CHAR(16),
    address CHAR(30),
    age    INT,
    job    CHAR(12),
    join_date DATE,
    image  CHAR(10),
    permission    CHAR(1),
    purchase_status CHAR(6),
    PRIMARY KEY(id)
);
CREATE TABLE supplier_info(
    supplier_id CHAR(5),
    supplier_name    CHAR(16),
    PRIMARY KEY    (supplier_id)
);


CREATE TABLE customer_order_records(
    id    CHAR(10),
    ordered_product CHAR(16),
    supplier_name CHAR(16),
    unit    CHAR(6),
    order_date DATE,
    estimated_submission_date DATE,
    actual_submission_date DATE,
    number  double(8,2),
    unit_price double(8,2),
    order_amount   double(8,2), 
    supplier_id CHAR(5),
    PRIMARY KEY(id),
    FOREIGN KEY (id) REFERENCES customer_info(id),
    FOREIGN KEY (supplier_id) REFERENCES supplier_info(supplier_id)
);
CREATE TABLE company_procurement_info(
    id    CHAR(10),
    supplier_id CHAR(5),
    supplier_contact CHAR(12),
    ordered_product  CHAR(16),
    stock_location    CHAR(16),
    detail            CHAR(16),
    order_unit    CHAR(6),
    order_number    double(8,2),
    order_unit_price double(8,2),
    restock_date DATE,
  	FOREIGN KEY (supplier_id) REFERENCES         supplier_info(supplier_id)

);
CREATE TABLE company_receivables_info(
    id                CHAR(10),
    receivable_sum    double(8,2),
    remaining_balance double(8,2),
    customer_name     CHAR(12),
    PRIMARY KEY(id),
    FOREIGN KEY (id) REFERENCES customer_info(id)
);


