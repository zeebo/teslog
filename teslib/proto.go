package teslib

const AuthURL = "/oauth/token"

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

const RefreshURL = "/oauth/token"

type RefreshRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

const VehiclesURL = "/api/1/vehicles"

type VehiclesResponse struct {
	Count    int64             `json:"count"`
	Response []VehicleResponse `json:"response"`
}

type VehicleResponse struct {
	Id                     int64       `json:"id"`
	VehicleId              int64       `json:"vehicle_id"`
	VIN                    string      `json:"vin"`
	DisplayName            string      `json:"display_name"`
	OptionCodes            string      `json:"option_codes"`
	Color                  interface{} `json:"color"`
	Tokens                 []string    `json:"tokens"`
	State                  string      `json:"state"`
	InService              bool        `json:"in_service"`
	IdString               string      `json:"id_s"`
	CalendarEnabled        bool        `json:"calendar_enabled"`
	APIVersion             int64       `json:"api_version"`
	BackseatToken          interface{} `json:"backseat_token"`
	BackseatTokenUpdatedAt interface{} `json:"backseat_token_updated_at"`
}

const DataURL = "/api/1/vehicles/%v/vehicle_data"

type DataResponse struct {
	Id                     int64                 `json:"id"`
	UserId                 int64                 `json:"user_id"`
	VehicleId              int64                 `json:"vehicle_id"`
	VIN                    string                `json:"vin"`
	DisplayName            string                `json:"display_name"`
	OptionCodes            string                `json:"option_codes"`
	Color                  interface{}           `json:"color"`
	Tokens                 []string              `json:"tokens"`
	State                  string                `json:"state"`
	InService              bool                  `json:"in_service"`
	IdString               string                `json:"id_s"`
	CalendarEnabled        bool                  `json:"calendar_enabled"`
	APIVersion             int64                 `json:"api_version"`
	BackseatToken          interface{}           `json:"backseat_token"`
	BackseatTokenUpdatedAt interface{}           `json:"backseat_token_updated_at"`
	DriveState             DriveStateResponse    `json:"drive_state"`
	ClimateState           ClimateStateResponse  `json:"climate_state"`
	ChargeState            ChargeStateResponse   `json:"charge_state"`
	GUISettings            GUISettingsResponse   `json:"gui_settings"`
	VehicleState           VehicleStateResponse  `json:"vehicle_state"`
	VehicleConfig          VehicleConfigResponse `json:"vehicle_config"`
}

const DriveStateURL = "/api/1/vehicles/%v/data_request/drive_state"

type DriveStateResponse struct {
	GPSAsOf                 int64    `json:"gps_as_of"`
	Heading                 int64    `json:"heading"`
	Latitude                float64  `json:"latitude"`
	Longitude               float64  `json:"longitude"`
	NativeLatitude          float64  `json:"native_latitude"`
	NativeLocationSupported int64    `json:"native_location_supported"`
	NativeLongitude         float64  `json:"native_longitude"`
	NativeType              string   `json:"native_type"`
	Power                   int64    `json:"power"`
	ShiftState              *string  `json:"shift_state"`
	Speed                   *float64 `json:"speed"`
	Timestamp               int64    `json:"timestamp"`
}

const ClimateStateURL = "/api/1/vehicles/%v/data_request/climate_state"

type ClimateStateResponse struct {
	BatteryHeater           bool        `json:"battery_heater"`
	BatteryHeaterNoPower    bool        `json:"battery_heater_no_power"`
	DriverTempSetting       float64     `json:"driver_temp_setting"`
	FanStatus               int64       `json:"fan_status"`
	InsideTemp              interface{} `json:"inside_temp"`
	IsAutoConditioningOn    interface{} `json:"is_auto_conditioning_on"`
	IsClimateOn             bool        `json:"is_climate_on"`
	IsFrontDefrosterOn      bool        `json:"is_front_defroster_on"`
	IsPreconditioning       bool        `json:"is_preconditioning"`
	IsRearDefrosterOn       bool        `json:"is_rear_defroster_on"`
	LeftTempDirection       interface{} `json:"left_temp_direction"`
	MaxAvailTemp            float64     `json:"max_avail_temp"`
	MinAvailTemp            float64     `json:"min_avail_temp"`
	OutsideTemp             interface{} `json:"outside_temp"`
	PassengerTempSetting    float64     `json:"passenger_temp_setting"`
	RightTempDirection      interface{} `json:"right_temp_direction"`
	SeatHeaterLeft          interface{} `json:"seat_heater_left"`
	SeatHeaterRearCenter    interface{} `json:"seat_heater_rear_center"`
	SeatHeaterRearLeft      interface{} `json:"seat_heater_rear_left"`
	SeatHeaterRearLeftBack  interface{} `json:"seat_heater_rear_left_back"`
	SeatHeaterRearRight     interface{} `json:"seat_heater_rear_right"`
	SeatHeaterRearRightBack interface{} `json:"seat_heater_rear_right_back"`
	SeatHeaterRight         interface{} `json:"seat_heater_right"`
	SideMirrorHeaters       interface{} `json:"side_mirror_heaters"`
	SmartPreconditioning    bool        `json:"smart_preconditioning"`
	SteeringWheelHeater     interface{} `json:"steering_wheel_heater"`
	Timestamp               int64       `json:"timestamp"`
	WiperBladeHeater        interface{} `json:"wiper_blade_heater"`
}

const ChargeStateURL = "/api/1/vehicles/%v/data_request/charge_state"

type ChargeStateResponse struct {
	BatteryHeaterOn             bool        `json:"battery_heater_on"`
	BatteryLevel                int64       `json:"battery_level"`
	BatteryRange                float64     `json:"battery_range"`
	ChargeCurrentRequest        int64       `json:"charge_current_request"`
	ChargeCurrentRequestMax     int64       `json:"charge_current_request_max"`
	ChargeEnableRequest         bool        `json:"charge_enable_request"`
	ChargeEnergyAdded           float64     `json:"charge_energy_added"`
	ChargeLimitSoc              int64       `json:"charge_limit_soc"`
	ChargeLimitSocMax           int64       `json:"charge_limit_soc_max"`
	ChargeLimitSocMin           int64       `json:"charge_limit_soc_min"`
	ChargeLimitSocStd           int64       `json:"charge_limit_soc_std"`
	ChargeMilesAddedIdeal       float64     `json:"charge_miles_added_ideal"`
	ChargeMilesAddedRated       float64     `json:"charge_miles_added_rated"`
	ChargePortDoorOpen          bool        `json:"charge_port_door_open"`
	ChargePortLatch             string      `json:"charge_port_latch"`
	ChargeRate                  float64     `json:"charge_rate"`
	ChargeToMaxRange            bool        `json:"charge_to_max_range"`
	ChargerActualCurrent        int64       `json:"charger_actual_current"`
	ChargerPhases               interface{} `json:"charger_phases"`
	ChargerPilotCurrent         int64       `json:"charger_pilot_current"`
	ChargerPower                int64       `json:"charger_power"`
	ChargerVoltage              int64       `json:"charger_voltage"`
	ChargingState               string      `json:"charging_state"`
	ConnChargeCable             string      `json:"conn_charge_cable"`
	EstBatteryRange             float64     `json:"est_battery_range"`
	FastChargerBrand            string      `json:"fast_charger_brand"`
	FastChargerPresent          bool        `json:"fast_charger_present"`
	FastChargerType             string      `json:"fast_charger_type"`
	IdealBatteryRange           float64     `json:"ideal_battery_range"`
	ManagedChargingActive       bool        `json:"managed_charging_active"`
	ManagedChargingStartTime    *int64      `json:"managed_charging_start_time"`
	ManagedChargingUserCanceled bool        `json:"managed_charging_user_canceled"`
	MaxRangeChargeCounter       int64       `json:"max_range_charge_counter"`
	NotEnoughPowerToHeat        bool        `json:"not_enough_power_to_heat"`
	ScheduledChargingPending    bool        `json:"scheduled_charging_pending"`
	ScheduledChargingStartTime  *int64      `json:"scheduled_charging_start_time"`
	TimeToFullCharge            float64     `json:"time_to_full_charge"`
	Timestamp                   int64       `json:"timestamp"`
	TripCharging                bool        `json:"trip_charging"`
	UsableBatteryLevel          int64       `json:"usable_battery_level"`
	UserChargeEnableRequest     interface{} `json:"user_charge_enable_request"`
}

const GUISettingsURL = "/api/1/vehicles/%v/data_request/gui_settings"

type GUISettingsResponse struct {
	GUI24HourTime       bool   `json:"gui_24_hour_time"`
	GUIChargeRateUnits  string `json:"gui_charge_rate_units"`
	GUIdistanceUnits    string `json:"gui_distance_units"`
	GuiRangeDisplay     string `json:"gui_range_display"`
	GUITemperatureUnits string `json:"gui_temperature_units"`
	Timestamp           int64  `json:"timestamp"`
}

const VehicleStateURL = "/api/1/vehicles/%v/data_request/vehicle_state"

type VehicleStateResponse struct {
	APIVersion         int64  `json:"api_version"`
	AutoparkStateV2    string `json:"autopark_state_v2"`
	AutoparkStyle      string `json:"autopark_style"`
	CalendarSupported  bool   `json:"calendar_supported"`
	CarVersion         string `json:"car_version"`
	CenterDisplayState int64  `json:"center_display_state"`
	DF                 int64  `json:"df"`
	DR                 int64  `json:"dr"`
	FT                 int64  `json:"ft"`
	HomelinkNearby     bool   `json:"homelink_nearby"`
	IsUserPresent      bool   `json:"is_user_present"`
	LastAutoparkError  string `json:"last_autopark_error"`
	Locked             bool   `json:"locked"`
	MediaState         struct {
		RemoteControlEnabled bool `json:"remote_control_enabled"`
	} `json:"media_state"`
	NotificationsSupported  bool    `json:"notifications_supported"`
	Odometer                float64 `json:"odometer"`
	ParsedCalendarSupported bool    `json:"parsed_calendar_supported"`
	PF                      int64   `json:"pf"`
	PR                      int64   `json:"pr"`
	RemoteStart             bool    `json:"remote_start"`
	RemoteStartSupported    bool    `json:"remote_start_supported"`
	RT                      int64   `json:"rt"`
	SoftwareUpdate          struct {
		ExpectedDurationSec int64  `json:"expected_duration_sec"`
		Status              string `json:"status"`
	} `json:"software_update"`
	SpeedLimitMode struct {
		Active          bool    `json:"active"`
		CurrentLimitMPH float64 `json:"current_limit_mph"`
		MaxLimitMPH     int64   `json:"max_limit_mph"`
		MinLimitMPH     int64   `json:"min_limit_mph"`
		PINCodeSet      bool    `json:"pin_code_set"`
	} `json:"speed_limit_mode"`
	SunRoofPercentOpen int64  `json:"sun_roof_percent_open"`
	SunRoofState       string `json:"sun_roof_state"`
	Timestamp          int64  `json:"timestamp"`
	ValetMode          bool   `json:"valet_mode"`
	ValetPinNeeded     bool   `json:"valet_pin_needed"`
	VehicleName        string `json:"vehicle_name"`
}

const VehicleConfigURL = "/api/1/vehicles/%v/data_request/vehicle_config"

type VehicleConfigResponse struct {
	CanAcceptNavigationRequests bool   `json:"can_accept_navigation_requests"`
	CanActuateTrunks            bool   `json:"can_actuate_trunks"`
	CarSpecialType              string `json:"car_special_type"`
	CarType                     string `json:"car_type"`
	ChargePortType              string `json:"charge_port_type"`
	EuVehicle                   bool   `json:"eu_vehicle"`
	ExteriorColor               string `json:"exterior_color"`
	HasAirSuspension            bool   `json:"has_air_suspension"`
	HasLudicrousMode            bool   `json:"has_ludicrous_mode"`
	MotorizedChargePort         bool   `json:"motorized_charge_port"`
	PerfConfig                  string `json:"perf_config"`
	PLG                         bool   `json:"plg"`
	RearSeatHeaters             int64  `json:"rear_seat_heaters"`
	RearSeatType                int64  `json:"rear_seat_type"`
	RHD                         bool   `json:"rhd"`
	RoofColor                   string `json:"roof_color"`
	SeatType                    int64  `json:"seat_type"`
	SpoilerType                 string `json:"spoiler_type"`
	SunRoofInstalled            int64  `json:"sun_roof_installed"`
	ThirdRowSeats               string `json:"third_row_seats"`
	Timestamp                   int64  `json:"timestamp"`
	TrimBadging                 string `json:"trim_badging"`
	WheelType                   string `json:"wheel_type"`
}
