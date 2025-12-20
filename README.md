# Truck Management Simulation
This project is a fleet data platform for trucking logistics. Vehicle telemetry is generated using the [Simulator for Urban MObility (SUMO)](https://eclipse.dev/sumo/) but could be seamlessly swapped out with real data from IoT devices.

Data flows through an event-driven pipeline for ingestion, processing, storage, and alerting, enabling real-time visibility and historical analysis of fleet operations. The architecture is designed to scale from local simulation to real-world deployment while keeping components decoupled and extensible.

## Architecture
![Diagram.png](docs/assets/diagram.png)

## Usage
```bash
docker compose up -d
```

>NOTE:
>Very much in progress so it won't work yet. Check out the early-demo branch for a working demo of some of the components

## TODO
- Change core service app into simple ingestor
- Update Business logic to work with SUMO data
- Add Redis
- (Future) Deploy on AWS using Terraform

## Journey
- Local demo api server
- Docker compose with multiple services (Kafka, Postgres)
- Dockerize main app
- Adapt to trucking use case (SUMO, Redis, Real Services) <- In Progress

## Dependencies
SUMO - https://eclipse.dev/sumo/
