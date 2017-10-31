# go-example-api

RESTful API of products, following TDD and using dev/test PostgreSQL databases

# how to use
1. create *products* table

```sql
CREATE TABLE products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)
```

2. set environment variables

* DEV_DB_USERNAME
* DEV_DB_PASSWORD
* DEV_DB_NAME
* TEST_DB_USERNAME
* TEST_DB_PASSWORD
* TEST_DB_NAME

3. run `go test -v`

**complete**
