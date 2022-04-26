# python3.6

import random
from paho.mqtt import client as mqtt_client
import json
from db import dataChecker as dataChecker

broker = 'mqtt.antares.id'
port = 1883
topic = "/oneM2M/resp/antares-cse/0b6b7781f24344d5:b0ec30902bc42bdf/json"
# generate client ID with pub prefix randomly
client_id = f'python-mqtt-{random.randint(0, 100)}'

def connect_mqtt() -> mqtt_client:
    def on_connect(client, userdata, flags, rc):
        if rc == 0:
            print("Connected to MQTT Broker!")
        else:
            print("Failed to connect, return code %d\n", rc)

    client = mqtt_client.Client(client_id)
    client.on_connect = on_connect
    client.connect(broker, port)
    return client

dataList = []

def cekData(payload):
    parse = json.loads(payload)
    con = parse["m2m:rsp"]["pc"]["m2m:cin"]["con"]
    id = parse["m2m:rsp"]["pc"]["m2m:cin"]["pi"]
    dataList.append(json.loads(con))
    for data in dataList:
        dataChecker(data, id)

def subscribe(client: mqtt_client):
    def on_message(client, userdata, msg):
        cekData(msg.payload)
        #print(f"Received `{msg.payload.decode()}` from `{msg.topic}` topic")

    client.subscribe(topic)
    client.on_message = on_message


def run():
    client = connect_mqtt()
    subscribe(client)
    client.loop_forever()

if __name__ == '__main__':
    run()
