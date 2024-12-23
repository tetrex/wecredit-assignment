# wecredit-assignment

### Implement mobile number/OTP-based authentication systems on a large scale and

#### handle the scenarios below.

1. Create new accounts and Store them in some Database.
2. Login using OTP to registered mobile number
3. Resend OTP if OTP expires
4. Get current User Details once Logged In.
5. Register API and Login API should have fingerprinting to differentiate if the User is logging In from different devices.
6. Deployment strategy for better operational Efficiency.

## Overview

This project is a backend API server for the wecredit-assignment. It provides various endpoints for authentication, user management, and health checks.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/tetrex/wecredit-assignment.git
   cd wecredit-assignment
   ```

2. Install dependencies:
   ```sh
   go mod download
   ```

## Check Makefile

> check makefile for other commands

## Starting the Local Server

1. Start the local server using Docker Compose:

   ```sh
   make dev
   ```

## Starting the Production Server

1. Build and start the production server using Docker Compose:

   ```sh
   docker compose -f docker-compose.prod.yml up -d --build
   ```

2. To stop the production server:

   ```sh
   docker compose -f docker-compose.prod.yml down
   ```

3. check docker-compose.swarm.yml for docker swarm config

## Viewing Swagger Documentation

1. Start the server (local or production).
2. Open your browser and navigate to [http://localhost:8000/docs/index.html](http://_vscodecontentref_/17) to view the Swagger documentation.

## Making API Requests

You can use tools like `curl` or Postman to make API requests. Below are some example requests:

1. Health Check:

   ```sh
   curl -X GET http://localhost:8000/
   ```

> visit swagger docs pages for api requests

## deployemnt using github action

> check .github folder

## Check Makefile

> Check the Makefile for other commands

## Migration Commands

- Apply all up migrations:

  ```sh
  make migrate-up
  ```

- Revert all down migrations:

  ```sh
  make migrate-down
  ```

- Create a new migration file:
  ```sh
  make migrate-new name=<migration_name>
  ```

## Usage of sqlc

1. Define your SQL queries in `.sql` files within the querys directory.
2. Configure `sqlc` by editing the sqlc.yaml file.
3. Generate the Go code by running:

```sh
make sqlc_gen
```

This will generate Go code for your SQL queries in the db/sqlc directory. You can then use these generated functions in your application code.
