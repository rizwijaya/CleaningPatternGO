package device

import "time"

type Device struct {
	No_device         int
	Id_device         string
	Energy            string
	Power             string
	Voltage           string
	Ampere            string
	PowerF            string
	LightSensor       string
	Battery           string
	Relay_status      int
	Brightness        string
	Device_time       time.Time
	Alarm_type        int
	Device_registered time.Time
	Restart_status    int
	Blue_status       int
	Lamp_treshold     string
}
