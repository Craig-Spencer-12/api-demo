# Truck Management Simulation
This project is a fleet data platform for trucking logistics. Vehicle telemetry is generated using the [Simulator for Urban MObility (SUMO)](https://eclipse.dev/sumo/) but could be seamlessly swapped out with real data from IoT devices.

Data flows through an event-driven pipeline for ingestion, processing, storage, and alerting, enabling real-time visibility and historical analysis of fleet operations. The architecture is designed to scale from local simulation to real-world deployment while keeping components decoupled and extensible.

## Architecture
![Diagram.png](docs/assets/diagram.png)

## Usage

### Fully Containerized
Simulation runs containerized in CLI mode. Good for performance but less visual.

```bash
docker compose --profile sim --profile migrate up -d
```

>NOTE: The `--profile migrate` is only necessary when running for the first 
time as it sets up the postgres schemas.

### Locally Run Simulator
Simulation runs loccally next to Docker stack. This lets SUMO run in gui mode showing full visual interface.

>NOTE: Install SUMO here - https://eclipse.dev/sumo/

```bash
docker compose up -d
```

```bash
python3 simulation/bridge.py
```

## TODO
- Create more services based on example-service
- Crete querying service for user facing postgres interactions
- Alerting system for certain events (ie. Exessive speeding)
- Improve visuals of sim (Try different maps and vehicle amounts)
- (Future) Deploy on AWS using Terraform

## Journey
- Local demo api server
- Docker compose with multiple services (Kafka, Postgres)
- Dockerize main app
- Adapt to trucking use case (SUMO, Redis, Real Services)

## Dependencies
SUMO - https://eclipse.dev/sumo/
