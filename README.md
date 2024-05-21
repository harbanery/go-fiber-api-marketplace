# API CRUD Marketplace

This application serves as a robust platform for managing Marketplace activities through a set of CRUD (Create, Read, Update, Delete) operations. Built with flexibility and efficiency in mind, it leverages the capabilities of Go programming language along with a suite of essential tools and packages to deliver a seamless experience. Postman is utilized for testing and running the API endpoints, ensuring smooth operation and easy integration.

## Technology Stack

This API is built using the following packages:

- Go: The programming language used to develop the backend logic.
- Fiber v2: A web framework for Golang used to handle HTTP requests and responses.
- Air by Cosmtrek: A live reload tool for Go applications, used to automatically rebuild and restart the server during development.
- Bluemonday by microcosm-cc: A package for sanitizing HTML to prevent cross-site scripting (XSS) attacks.
- JWT-Go v5 by golang-jwt: A package for working with JSON Web Tokens (JWT) for authentication.
- Validator v10 by go-playground: A library for validating structs and fields.
- Cloudinary: A cloud service for managing media assets like images and videos.

## Installation

1. Clone this repository to your local directory.
2. Make sure you have Go installed. If not, you can download it from [here](https://go.dev/dl/).
3. Install Fiber v2, Air by Cosmtrek, JWT-Go v5, Validator v10, and bluemonday packages by running the following commands:

   ```powershell
      go get -u github.com/gofiber/fiber/v2
      go get -u github.com/cosmtrek/air
      go get -u github.com/golang-jwt/jwt/v5
      go get -u github.com/go-playground/validator/v10
      go get -u github.com/microcosm-cc/bluemonday
   ```

4. Set up Cloudinary account, install the packages, and get your API credentials.

   ```powershell
      go get -u github.com/cloudinary/cloudinary-go
   ```

## Running the Application

1. Open Postman.
2. Import the Postman collection provided in this repository.
3. Set environment variables if needed.
4. Start accessing the API endpoints.

<!-- ## Endpoints

Here is a list of available endpoints:

1. Auth

   - GET /register: Get a list of all users.

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

3. Customers

   - GET /sellers: Get a list of all sellers.

   - GET /sellers/{id}: Get a seller by ID.

   - POST /sellers: Add a new seller.

   - PUT /sellers/{id}: Update seller information.

   - DELETE /sellers/{id}: Delete a seller.

4. Products

   - GET /products: Get a list of all products.

   - GET /products/{id}: Get a product by ID.

   - POST /products: Add a new product.

   - PUT /products/{id}: Update product information.

   - DELETE /products/{id}: Delete a product.

5. Categories

   - GET /categories: Get a list of all categories.

   - GET /categories/{id}: Get a category by ID.

   - POST /categories: Add a new category.

   - PUT /categories/{id}: Update category information.

   - DELETE /categories/{id}: Delete a category. -->

## Reference

Feel free if you want to check it out:

[Mockup Web Figma](https://www.figma.com/design/F2wIb9WHG4kntUkbFC39OB/Mockup-Web?node-id=0-1&t=yXyuxr2edx3A68Wx-0)

[GoFiber Gitbook by Zaki Maliki](https://zakimaliki.gitbook.io/gofiber)

[Fiber v2](https://docs.gofiber.io)

[Air by Cosmtrek](https://github.com/cosmtrek/air)

[Jwt-Go](https://github.com/golang-jwt/jwt)

[Validator v10](https://github.com/go-playground/validator)

[Refresh Tokens Auth by Tanveer Hassan](https://medium.com/monstar-lab-bangladesh-engineering/jwt-auth-in-go-part-2-refresh-tokens-d334777ca8a0)

[bluemonday](https://github.com/microcosm-cc/bluemonday)

[Cloudinary Go SDK](https://cloudinary.com/documentation/go_integration)
