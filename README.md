# Kompit's Golang Backend Engineer Interview Assignment

Welcome to the recruitment process for the Golang Back-end Engineer role. We are seeking candidates who have experience in developing scalable and efficient backend systems using Golang. Strong proficiency in Golang, along with knowledge of web frameworks, preferably Gin, is highly desirable for this role.

## 1. Getting Started

### 1.1. Prerequisites

To get started, you will need to have the following installed on your machine:

1. [Golang](https://golang.org/doc/install) >= 1.21
2. [Docker](https://docs.docker.com/get-docker/)
3. [Docker Compose](https://docs.docker.com/compose/install/)
4. [GNU Make](https://www.gnu.org/software/make/)
5. [oapi-codegen](https://github.com/deepmap/oapi-codegen)

    Install the latest version with:
    ```
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    ```
6. [mock](https://github.com/uber-go/mock)

    Install the latest version with:
    ```
    go install go.uber.org/mock/mockgen@latest
    ```
7. NodeJS & NPM to run the OpenAPI viewer.

### 1.2. Installation

1. Clone this repository to your local machine:

2. Run the following command to generate required files:

   ```bash
   make init
   ```

3. Run the following command to start the application:

   ```bash
   docker-compose up --build
   ```

   This will rebuild the application and the database. You can access the application at http://localhost:8080.

### 1.3. Directory Structure

The directory structure of this project is as follows:

```text
/
├── handler
│   ├── endpoints_test.go  Test cases for the endpoints.
│   ├── endpoints.go       Implementation of the endpoints.
│   ├── handler_test.go    Test cases for the handler.
│   ├── handler.go         The handler struct.
├── api.yml                The OpenAPI specification that you have to follow.
├── database.sql           The database schema that you have to implement.
├── main.go                The main function.
```

You are allowed to add new files and directories to the project, e.g. to add new packages for repositories or interactors/usecases.s


## 2. Assignment

### 2.1. API Implementation

You will need to implement the following endpoints:
1. `POST /competitions`
2. `GET /competitions/{id}`
3. `POST /competitions/{id}/join`
4. `POST /register`
5. `GET /users/{username}`

You can see the OpenAPI specification in [`api.yml`](api.yml) for more details on the endpoints, including the request and response payloads.

You can easily read the specification in browser by executing this command:

```bash
npx @redocly/cli preview-docs api.yml
```

### 2.2. Database Design

You will need to design the database schema for the application. The database schema is defined in `database.sql`. You will need to implement the schema in the PostgreSQL database. You can see database configuration in the the `docker-compose.yml` file.

To reset the database in the docker container, you can run the following command:

```bash
docker-compose down --volumes
```

### 2.3. Unit Test

You will need to write unit tests for the endpoints and handler. You can see the test cases in `handler/endpoints_test.go` and `handler/handler_test.go`. Should you write more implementation code, you will need to write unit tests to cover the new code, as we will take the unit test coverage as one of the judging criteria.

To run the unit tests, you can run the following command:

```bash
make test
```

## 3. Submission

To submit this assignment, you will need to create a new _public repository_ in your GitHub account and push the code to the repository. You will need to share the repository link with us.

## 4. Judging Criteria

We will judge your submission based on the following criteria:
1. **Functional Correctness**

    We will execute an automated test on your API server. The test will send HTTP requests to the URL `http://localhost:8080`. You can find a sample testing script in the `test_api.sh` file (Hint: We will use Postman). Please ensure that your API can receive and respond with the correct value or format.
2. **Code structure**

    Your code should have a clear and logical structure, with appropriate separation of concerns and modularization. Use packages and directories to organize your codebase, and ensure that each file has a specific purpose and responsibility. This will make your code easier to understand, maintain, and scale in the future.
3. **Database design**

    Designing an efficient and well-structured database schema is crucial for the success of your application. Take into account the relationships between objects and define appropriate tables, columns, and constraints to accurately store and retrieve data. Pay attention to normalization and indexing to optimize query performance and maintain data integrity.
4. **Unit test coverage**

    It is essential to have comprehensive unit tests to ensure the correctness of your code. Write unit tests to cover different scenarios and edge cases for each endpoint of your API server. Aim for a high test coverage to minimize the risk of bugs and ensure the stability of your application.

5. **Code readability**
    Writing clean and readable code is crucial for collaboration and maintainability. Use descriptive variable and function names, follow consistent formatting conventions, and include comments where necessary to explain the purpose and logic of your code. Adopting a consistent coding style will make it easier for others to understand and contribute to your project.