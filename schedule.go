package main

import (
	"fmt"
	"os"
	"time"

	"github.com/zeebo/teslog/teslib"
)

type state string

const (
	stateUnknown      state = "Unknown"
	stateOffline      state = "Offline"
	stateDriving      state = "Driving"
	stateCharging     state = "Charging"
	stateConditioning state = "Conditioning"
	stateOnline       state = "Online"
	stateQuiescent    state = "Quiescent"
)

var lastState state = stateUnknown

func getState(data *teslib.DataResponse) state {
	if data == nil {
		return stateOffline
	} else if shift := data.DriveState.ShiftState; shift != nil &&
		(*shift == "R" || *shift == "D" || *shift == "N") {
		return stateDriving
	} else if speed := data.DriveState.Speed; speed != nil && *speed > 0 {
		return stateDriving
	} else if charge := data.ChargeState.ChargingState; charge == "Charging" || charge == "Starting" {
		return stateCharging
	} else if data.ClimateState.IsClimateOn {
		return stateConditioning
	} else if data.VehicleState.CenterDisplayState != 0 {
		return stateOnline
	}
	return stateQuiescent
}

func scheduleNext(data *teslib.DataResponse) time.Duration {
	state := getState(data)
	if state != lastState {
		fmt.Fprintf(os.Stderr, "%s: state change from %s to %s.\n", time.Now(), lastState, state)
	}
	lastState = state

	switch state {
	case stateUnknown, stateQuiescent:
		return 11 * time.Minute
	default:
		return time.Minute
	}
}
