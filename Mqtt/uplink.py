# python3.6
import mysql.connector

# Konfigurasi Database
def connector():
    mydb = mysql.connector.connect(
        host="localhost",
        user="root",
        password="",
        database="tamaskapju"
    )
    return mydb

def hexD(hex):
    return int(hex, 16)

def deviceTime(time):
    return str(hexD(time[4:8])) + "-" + str(hexD(time[2:4])) + "-" + str(hexD(time[0:2])) + " " + str(hexD(time[8:10])) + ":" + str(hexD(time[10:12])) + ":" + str(hexD(time[12:14]))

def insertHistory(data, id):
    mydb = connector()
    mycursor = mydb.cursor()
    sql = "INSERT INTO history (id_device, energy, power, voltage, ampere, powerF, LightSensor, battery, relay_status, brightness, device_time) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)"
    val = (id, hexD(data["data"][2:14])/10000, hexD(data["data"][14:18]), hexD(data["data"][18:22]), hexD(data["data"][22:26])/100, hexD(data["data"][26:28])/100, hexD(data["data"][28:32]), hexD(data["data"][32:36])/100, hexD(data["data"][36:38]), hexD(data["data"][38:40]), deviceTime(data["data"][40:54]))
    mycursor.execute(sql, val)
    mydb.commit()
    return

def uplinkDevice(data, id):
    mydb = connector()
    mycursor = mydb.cursor()
    sql = "UPDATE devices SET energy = %s, power = %s, voltage = %s, ampere = %s, powerF = %s, LightSensor = %s, battery = %s, relay_status = %s, brightness = %s, device_time = %s WHERE id_device = %s"
    val = (hexD(data["data"][2:14])/10000, hexD(data["data"][14:18]), hexD(data["data"][18:22]), hexD(data["data"][22:26])/100, hexD(data["data"][26:28])/100, hexD(data["data"][28:32]), hexD(data["data"][32:36])/100, hexD(data["data"][36:38]), hexD(data["data"][38:40]), deviceTime(data["data"][40:54]), id)
    mycursor.execute(sql, val)
    mydb.commit()
    return
def uplinkAlarm(data, id):
    mydb = connector()
    mycursor = mydb.cursor()

    if data["data"][0:2] == "25":
        sql = "UPDATE device SET ampere = %s, battery = %s, alarm_type = %s, device_time = %s WHERE id_device = %s"
        val = (hexD(data["data"][2:6]), hexD(data["data"][6:10]), hexD(data["data"][10:12]), hexD(data["data"][12:26]), id)
    else:    
        if data["data"][0:2] == "22":
            val1 = "relay_status" 
        elif data["data"][0:2] == "23":
            val1 = "brightness" 
        elif data["data"][0:2] == "24":
            val1 = "restart_status" 
        sql = "UPDATE devices SET "+ val1 + " = %s, battery = %s, alarm_type = %s, device_time = %s WHERE id_device = %s"
        val = (hexD(data["data"][2:4]), hexD(data["data"][4:8]), hexD(data["data"][8:10]), hexD(data["data"][10:24]), id)
    
    mycursor.execute(sql, val)
    mydb.commit()
    return
def uplinkConfig(data, id):
    mydb = connector()
    mycursor = mydb.cursor()
    sql = "UPDATE device_conf SET houron_hour = %s, minon_hour = %s, houroff_hour = %s, houron_min = %s, interval_loop = %s, interval_send = %s, mode = %s WHERE id_device = %s"
    val = (hexD(data["data"][2:4]), hexD(data["data"][4:6]), hexD(data["data"][6:8]), hexD(data["data"][8:10]), hexD(data["data"][8:16]), hexD(data["data"][16:24]), hexD(data["data"][24:26]), id)
    mycursor.execute(sql, val)
    mydb.commit()
    return