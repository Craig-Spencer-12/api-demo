import os
import sys
import time
import requests
import json
import traci

def run_bridge():
    IS_DOCKER = os.environ.get("AM_I_IN_DOCKER", "false").lower() == "true"

    if IS_DOCKER:
        sumo_binary = "sumo"
        config_path = os.path.join("sumo", "sim.sumocfg")
        ingestor_url = "http://ingestor:8080/ingest"
    else:
        sumo_binary = "sumo-gui"
        config_path = os.path.join("simulation", "sumo", "sim.sumocfg")
        ingestor_url = "http://localhost:8080/ingest"

    try:
        traci.start([sumo_binary, "-c", config_path])
        print(f"Started {sumo_binary} successfully!")
    except Exception as e:
        print(f"Failed to start simulation: {e}")
        return

    while traci.simulation.getMinExpectedNumber() > 0:
        traci.simulationStep()
        
        truck_ids = traci.vehicle.getIDList()
        for t_id in truck_ids:
            x, y = traci.vehicle.getPosition(t_id)
            lon, lat = traci.simulation.convertGeo(x, y)
            speed = traci.vehicle.getSpeed(t_id)

            payload = {
                "truck_id": t_id,
                "lat": lat,
                "long": lon,
                "speed": round(speed, 2),
                "timestamp": time.time()
            }

            # Debug
            # print(json.dumps(payload))

            try:
                requests.post(ingestor_url, json=payload, timeout=0.1)
            except requests.exceptions.RequestException:
                pass

    traci.close()

if __name__ == "__main__":
    run_bridge()