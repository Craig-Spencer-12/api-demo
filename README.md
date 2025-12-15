# API-Demo
This project is to demo a full implentaion of a Go api, connected to a PostgreSQL database, using Kafka for sending events.

## Usage
Spin up services
```bash
    docker compose up -d
```
Curl - This is how you interact with the api. You could swap this out with Postman
```bash
    make user bob
    make getUsers
```

## TODO
- Test multiple consumers
- Deploy on AWS using Terraform
- Make user flow diagram
- Make architecture doc
- Create more concrete use case (trucking logistics)
- Create better way to interact with Kafka
