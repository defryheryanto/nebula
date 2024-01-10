# Nebula Dashboard
Nebula is a simple monitoring dashboard designed to provide comprehensive log searching capabilities.
We're aiming to have more comprehensive monitoring functionalities like Application Performance Monitoring (APM), Http Logging, and even Span Tracing.

## Setting Up Application
### Pre-requisites
1. [Go](https://go.dev/)
2. [MongoDB](https://www.mongodb.com/)
3. [PostgreSQL](https://www.postgresql.org/)

### Running the app
1. Build go binary under `./cmd/dashboard'
2. Ensure that variables in `.env.example` is exists and correct in your environment variables
3. Run/Build go binary under `./database/migrate.go`
4. Run the app go binary (output from step 1)
