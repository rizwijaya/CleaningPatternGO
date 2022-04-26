# python3.6
from uplink import *
from downlink import *

def dataChecker(data, id):
    hd = data["data"][0:2]
    if hd == "11": #Jika sensor data
        insertHistory(data, id)
        uplinkDevice(data, id)
    elif hd == "22" or hd == "23" or hd == "24" or hd == "25": #jika alarm data
        uplinkAlarm(data, id)
    elif hd == "33": #jika config data
        uplinkConfig(data, id)
    elif hd == "AA": #jika set config data
        downlinkConfig(data, id)
    elif hd == "00" or hd == "01" or hd == "BB" or hd == "AB" or hd == "AC" or hd == "AD" or hd == "AF" or hd == "BA": #jika control
        downlinkRelay(data, id)
    elif hd == "AE": #jika set listrik
        downlinkCalibration(data, id)