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

def downlinkConfig(data, id):
    mydb = connector()
    mycursor = mydb.cursor()
    sql = "UPDATE device_conf SET houron_hour = %s, minon_hour = %s, houroff_hour = %s, houron_min = %s, interval_loop = %s WHERE id_device = %s"
    val = (hexD(data["data"][2:4]), hexD(data["data"][4:6]), hexD(data["data"][6:8]), hexD(data["data"][8:10]), hexD(data["data"][10:12]), id)
    mycursor.execute(sql, val)
    mydb.commit()
    return

def downlinkRelay(data, id):
    mydb = connector()
    mycursor = mydb.cursor()
    dbnya = "device"
    val2 = hexD(data["data"][2:4])
    if data["data"][0:2] == "00": # Control Relay
        val1 = "relay_status"
    elif data["data"][0:2] == "01": # Change Brightness
        val1 = "brightness"
    elif data["data"][0:2] == "BB": #Control Bluetooth
        val1 = "blue_status"
    elif data["data"][0:2] == "AB": #Interval Send Data
        val1 = "interval_send"
        dbnya = "device_conf"
        val2 = hexD(data["data"][2:6])
    elif data["data"][0:2] == "AC": #set totalisator in kwh
        val1 = "energy"
        val2 = hexD(data["data"][2:14])/10000
    elif data["data"][0:2] == "AD": #calibration device time
        val1 = "device_time"
        val2 = deviceTime(data["data"][2:16])
    elif data["data"][0:2] == "AF": #Set lamp current treshold
        val1 = "lamp_treshold"
        val2 = hexD(data["data"][2:6])/100
    elif data["data"][0:2] == "BA": #Calibration Baterry Voltage
        val1 = "battery"
        val2 = hexD(data["data"][2:6])/100
        
    sql = "UPDATE " + dbnya + " SET " + val1 + " = %s WHERE id_device = %s"
    val = (val2, id)
    mycursor.execute(sql, val)
    mydb.commit()
    return

def downlinkCalibration(data, id):
    mydb = connector()
    mycursor = mydb.cursor()
    sql = "UPDATE device SET ampere = %s, voltage = %s, power = %s WHERE id_device = %s"
    val = (hexD(data["data"][2:6]), hexD(data["data"][6:10]), hexD(data["data"][10:16]), id)
    mycursor.execute(sql, val)
    mydb.commit()
    return