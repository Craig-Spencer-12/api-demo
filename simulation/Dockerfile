FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    software-properties-common \
    && add-apt-repository ppa:sumo/stable \
    && apt-get update && apt-get install -y \
    sumo \
    sumo-tools \
    python3 \
    python3-pip \
    && rm -rf /var/lib/apt/lists/*


ENV SUMO_HOME=/usr/share/sumo
WORKDIR /app

# TODO: Should not need to install traci, should be in sumo/tools
RUN pip3 install traci requests --break-system-packages

COPY . .

CMD ["python3", "bridge.py"]