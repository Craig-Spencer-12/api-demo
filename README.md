# API-Demo
This project is to demo a full implentaion of a Go api, connected to a PostgreSQL database, using Kafka for sending events and deployed on AWS using Terraform.

## Usage
Use 3 seperate terminal windows if you want to see all the output

1. Consumer - This is were Kafka messages will be shown when a user is created
```bash
    make docker
    make consumer
```

2. API - This will show all the api traffic
```bash
    make server
```

3. Curl - This is how you interact with the api. You could swap this out with Postman
```bash
    make user
    make user
    make user
    make getUsers
```

## TODO
- Dockerize server
- Dockerize consumer
    - Test multiple consumers
- Add PostgreSQL database

- Deploy on AWS using Terraform