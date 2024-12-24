# Bookstore Service

The Bookstore Service is a gRPC-based application that manages book-related operations. This guide covers how to build, run, and test the service using Docker and `grpcurl`.

## Prerequisites

- Docker installed on your system.
- `grpcurl` installed for testing the gRPC service.

## Running the Service with Docker

1. **Build the Docker image**:

   ```sh
   docker build -t <docker_image_name>
   ```

2. **Run the Docker container**:

   ```sh
   docker run -d -e SECRET_KEY=<your-secret-key> -p 8080:8080 <docker_image_name>
   ```

   Replace `<your-secret-key>` with your actual secret key. The service will now be available on `localhost:8080`.

3. **Stop and Remove the Container**:

   ```sh
   docker stop <docker_image_name>
   docker rm <docker_image_name>
   ```

## Testing the Service with `grpcurl`

To test the service, ensure the server is running and accessible at `localhost:8080`.

1. **First Register on the servive and then login to obtain access token and use it in subsequent requests**:

   ```sh
   grpcurl -d '{"username": "testuser", "password": "testpassword"}' -plaintext localhost:8080 bookstore.BookService/Register
   ```

   This will register a user

   ```sh
   grpcurl -d '{"username": "testuser", "password": "testpassword"}' -plaintext localhost:8080 bookstore.BookService/Login
   ```

   This will respond with an access token



2. **Call a Specific Method**: Replace `<Method>` with the method name and `<Payload>` with the request payload in JSON format:

   ```sh
   grpcurl -plaintext -h "authorization":<access token> -d '<Payload>' localhost:8080 bookstore.BookService/Login/<Method>
   ```

## Notes

- Ensure port `8080` is available before running the service.
- Update the `dockerfile` if additional dependencies or configurations are required.


