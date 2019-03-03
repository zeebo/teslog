package main

import (
	"fmt"
	"io"

	"github.com/zeebo/mon/monflux"
	"github.com/zeebo/teslog/teslib"
)

func monitor(w io.Writer, data *teslib.DataResponse) {
	if data != nil {
		monitorData(w, data)
	}
	monflux.Collector{Measurement: "mon"}.Write(w)
}

func monitorData(w io.Writer, data *teslib.DataResponse) {
	quote := func(x string) string { return fmt.Sprintf("%q", x) }
	output := func(m, f string, val interface{}) { fmt.Fprintf(w, "%s,id=%d %s=%v\n", m, data.Id, f, val) }

	output("status", "state", quote(data.State))

	{
		drive := data.DriveState
		output := func(f string, val interface{}) { output("drive_state", f, val) }

		output("latitude", drive.Latitude)
		output("longitude", drive.Longitude)
		output("power", drive.Power)
		if drive.ShiftState != nil {
			output("shift_state", quote(*drive.ShiftState))
		}
		if drive.Speed != nil {
			output("speed", *drive.Speed)
		}
	}

	{
		climate := data.ClimateState
		output := func(f string, val interface{}) { output("climate_state", f, val) }

		output("driver_temp_setting", climate.DriverTempSetting)
		output("inside_temp", climate.InsideTemp)
		output("is_auto_conditioning_on", climate.IsAutoConditioningOn)
		output("is_climate_on", climate.IsClimateOn)
		output("outside_temp", climate.OutsideTemp)
		output("passenger_temp_setting", climate.PassengerTempSetting)
	}

	{
		charge := data.ChargeState
		output := func(f string, val interface{}) { output("charge_state", f, val) }

		output("battery_level", charge.BatteryLevel)
		output("battery_range", charge.BatteryRange)
		output("charge_rate", charge.ChargeRate)
		output("charger_power", charge.ChargerPower)
		output("charger_voltage", charge.ChargerVoltage)
		output("est_battery_range", charge.EstBatteryRange)
		output("ideal_battery_range", charge.IdealBatteryRange)
		output("time_to_full_charge", charge.TimeToFullCharge)
		output("usable_battery_level", charge.UsableBatteryLevel)
		output("charging_state", quote(charge.ChargingState))
	}

	{
		vehicle := data.VehicleState
		output := func(f string, val interface{}) { output("vehicle_state", f, val) }

		output("car_version", quote(vehicle.CarVersion))
		output("homelink_nearby", vehicle.HomelinkNearby)
		output("is_user_present", vehicle.IsUserPresent)
		output("locked", vehicle.Locked)
		output("odometer", vehicle.Odometer)
	}
}
