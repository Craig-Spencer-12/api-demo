# import os
# import sys
# import time
# import requests
# import json

# # 1. Setup path for TraCI
# if 'SUMO_HOME' in os.environ:
#     sys.path.append(os.path.join(os.environ['SUMO_HOME'], 'tools'))
# else:
#     sys.exit("Please declare environment variable 'SUMO_HOME'")

# import traci 

# def run_bridge():
#     # 2. Connect to the waiting SUMO GUI
#     # Ensure you ran: sumo-gui -c sim.sumocfg --remote-port 8813
#     try:
#         traci.init(port=8813)
#         print("Successfully connected to SUMO!")
#     except Exception as e:
#         print(f"Connection failed: {e}")
#         return

#     while traci.simulation.getMinExpectedNumber() > 0:
#         # Advance simulation
#         traci.simulationStep()

#         # Get all vehicle IDs
#         truck_ids = traci.vehicle.getIDList()

#         for t_id in truck_ids:
#             # 3. Compatible GPS Data Collection
#             # Get raw X,Y and convert to Geo (Lat/Lon)
#             x, y = traci.vehicle.getPosition(t_id)
#             lon, lat = traci.simulation.convertGeo(x, y)
            
#             speed = traci.vehicle.getSpeed(t_id)

#             # 4. Construct Telemetry Payload
#             # Flattened to make Go unmarshaling easier
#             payload = {
#                 "truck_id": t_id,
#                 "lat": lat,
#                 "long": lon,
#                 "speed": round(speed, 2),
#                 "timestamp": time.time()
#             }

#             print(json.dumps(payload))

#             # 5. Push to Go Server
#             try:
#                 # Short timeout so the simulation doesn't lag if Go is slow
#                 requests.post("http://localhost:8080/ingest", json=payload, timeout=0.05)
#             except requests.exceptions.RequestException:
#                 # Silent fail if Go server is not reachable
#                 pass

#     traci.close()
#     print("Simulation finished.")

# if __name__ == "__main__":
#     run_bridge()


import os
import sys
import time
import requests
import json
import traci

def run_bridge():
    # 1. SMART CONFIG: Are we in Docker or on Desktop?
    # Docker usually sets environment variables we can check
    IS_DOCKER = os.environ.get("AM_I_IN_DOCKER", "false").lower() == "true"

    if IS_DOCKER:
        # CLI version for the server/automated tests
        sumo_binary = "sumo"
        config_path = os.path.join("sumo", "sim.sumocfg")
        ingestor_url = "http://ingestor:8080/ingest"
    else:
        # GUI version for your demo!
        sumo_binary = "sumo-gui"
        config_path = os.path.join("simulation", "sumo", "sim.sumocfg")
        ingestor_url = "http://localhost:8080/ingest"

    try:
        # Start SUMO
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

            print(json.dumps(payload))

            try:
                requests.post(ingestor_url, json=payload, timeout=0.1)
            except requests.exceptions.RequestException:
                pass

    traci.close()

if __name__ == "__main__":
    run_bridge()