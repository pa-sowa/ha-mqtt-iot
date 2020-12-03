package iotconfig

import (
	"log"
	"os/exec"

	"../hadiscovery"
)

func (sw SensorHA) constructStateFunc() (f func() string) {
	var err error
	return func() string {
		var out []byte
		if len(sw.CommandState) > 1 {
			out, err = exec.Command(sw.CommandState[0], sw.CommandState[1:]...).Output()
		} else {

			out, err = exec.Command(sw.CommandState[0]).Output()
		}
		if err != nil {
			log.Printf("%s", err)
		}
		return string(out)
	}
}

func (se SensorHA) Translate() hadiscovery.Sensor {
	nse := hadiscovery.Sensor{}
	nse.UpdateInterval = se.UpdateInterval
	nse.ExpireAfter = int(nse.UpdateInterval + 1)
	nse.ForceUpdateMQTT = se.ForceUpdateMQTT
	nse.UnitOfMeasurement = se.UnitOfMeasurement
	if !se.ForceUpdateMQTT {
		nse.ExpireAfter = 0
	}
	nse.Name = se.Info.Name
	nse.UniqueID = se.Info.ID + "_" + hadiscovery.NodeID
	if se.Info.Icon != "" {
		nse.Icon = se.Info.Icon
	}

	if len(se.CommandState) > 0 {
		nse.StateFunc = se.constructStateFunc()
	} else {
		log.Fatalln("Missing state func, needed for sensors!")
	}

	nse.Initialize()

	return nse
}
