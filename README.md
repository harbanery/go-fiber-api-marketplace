# API CRUD Marketplace

This is an application for performing CRUD (Create, Read, Update, Delete) operations on a Marketplace. This application is built using Postman to run it.

## Technology Stack

This API is built using the following technologies:

- Go: The programming language used to develop the backend logic.

- Fiber v2: A web framework for Golang used to handle HTTP requests and responses.

- Air by Cosmtrek: A live reload tool for Go applications, used to automatically rebuild and restart the server during development.

## Installation

1. Clone this repository to your local directory.

2. Open Postman.

## Running the Application

1. Open Postman.

2. Import the Postman collection provided in this repository.

3. Set environment variables if needed.

4. Start accessing the API endpoints.

## Endpoints

Here is a list of available endpoints:

1. Users

   - GET /users: Get a list of all users.

   - GET /users/{id}: Get a user by ID.

   - POST /users: Add a new user.

   - PUT /users/{id}: Update user information.

   - DELETE /users/{id}: Delete a user.

2. Sellers

   - GET /sellers: Get a list of all sellers.

   - GET /sellers/{id}: Get a seller by ID.

   - POST /sellers: Add a new seller.

   - PUT /sellers/{id}: Update seller information.

   - DELETE /sellers/{id}: Delete a seller.

3. Products

   - GET /products: Get a list of all products.

   - GET /products/{id}: Get a product by ID.

   - POST /products: Add a new product.

   - PUT /products/{id}: Update product information.

   - DELETE /products/{id}: Delete a product.

4. Categories

   - GET /categories: Get a list of all categories.

   - GET /categories/{id}: Get a category by ID.

   - POST /categories: Add a new category.

   - PUT /categories/{id}: Update category information.

   - DELETE /categories/{id}: Delete a category.

## Reference

Feel free if you want to check it out:

[Air by Cosmtrek](https://github.com/cosmtrek/air)

[GoFiber by Zaki Maliki](https://zakimaliki.gitbook.io/gofiber)

[Mockup Web Figma](https://www.figma.com/design/F2wIb9WHG4kntUkbFC39OB/Mockup-Web?node-id=0-1&t=yXyuxr2edx3A68Wx-0)
