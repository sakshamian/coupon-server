# Coupon service

## Setup & Installation

- Create .env file in root directory and copy the contents:
    ```
    GO_PORT=0.0.0.0:7170
    DEBUG=true
    
    # database
    DB_USER=root
    DB_PASSWORD=12345678
    DB_HOST=mysqldb
    DB_NAME=coupon_db
    DB_PORT=3306
    MAX_OPEN_CONNECTIONS=200
    MAX_IDLE_CONNECTIONS=20
    
    # redis
    REDIS_HOST=redis:6379
    ```
- Run the app
    ```
    docker compose up -d --build
    ```

## Caching notes
  The application uses redis to cache coupons. When the API to get applicable coupons gets coupons from database, we also store them in redis. The key used to store coupon data is the unique coupon code. The api to validate a coupon checks the redis for coupon data, if it finds the coupon, we proceed with the validations. Otherwise, the coupon is fetched from database.

## Architectural design
  ![Coupon architecture](https://github.com/user-attachments/assets/0a9a7084-246e-4ef4-8c74-d509d0dea5ee)
  
## OpenAPI documentation link:
  - [Download Postman Collection](api/api-docs.json)

## Additonal notes:
  The application uses resterros package for error handling and responses package for response handling. These packages are written by me for maintaining consistency in responses and errors.
