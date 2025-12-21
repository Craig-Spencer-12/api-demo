FROM ubuntu:latest

# Prevent interactive prompts
ENV DEBIAN_FRONTEND=noninteractive

# 1. Install prerequisites and add the official SUMO PPA
RUN apt-get update && apt-get install -y \
    software-properties-common \
    && add-apt-repository ppa:sumo/stable \
    && apt-get update && apt-get install -y \
    sumo \
    sumo-tools \
    python3 \
    python3-pip \
    && rm -rf /var/lib/apt/lists/*

# 2. Set SUMO_HOME
ENV SUMO_HOME=/usr/share/sumo
WORKDIR /app

# 3. Install Python dependencies
RUN pip3 install traci requests --break-system-packages

COPY . .

CMD ["python3", "bridge.py"]