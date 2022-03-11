package hadiscovery

func (d AlarmControlPanel) RawID() string {
	return "alarm_control_panel"
}
func (d AlarmControlPanel) UniqueID() string {
	return d.uniqueID
}
func (d *AlarmControlPanel) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type AlarmControlPanel struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	Code                 string `json:"code"`
	CodeArmRequired      bool   `json:"code_arm_required"`
	CodeDisarmRequired   bool   `json:"code_disarm_required"`
	CodeTriggerRequired  bool   `json:"code_trigger_required"`
	CommandTemplate      string `json:"command_template"`
	CommandTopic         string `json:"command_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	EnabledByDefault       bool   `json:"enabled_by_default"`
	Encoding               string `json:"encoding"`
	EntityCategory         string `json:"entity_category"`
	Icon                   string `json:"icon"`
	JsonAttributesTemplate string `json:"json_attributes_template"`
	JsonAttributesTopic    string `json:"json_attributes_topic"`
	Name                   string `json:"name"`
	ObjectId               string `json:"object_id"`
	PayloadArmAway         string `json:"payload_arm_away"`
	PayloadArmCustomBypass string `json:"payload_arm_custom_bypass"`
	PayloadArmHome         string `json:"payload_arm_home"`
	PayloadArmNight        string `json:"payload_arm_night"`
	PayloadArmVacation     string `json:"payload_arm_vacation"`
	PayloadAvailable       string `json:"payload_available"`
	PayloadDisarm          string `json:"payload_disarm"`
	PayloadNotAvailable    string `json:"payload_not_available"`
	PayloadTrigger         string `json:"payload_trigger"`
	Qos                    int    `json:"qos"`
	Retain                 bool   `json:"retain"`
	StateTopic             string `json:"state_topic"`
	UniqueId               string `json:"unique_id"`
	ValueTemplate          string `json:"value_template"`
	uniqueID               string
	rawID                  string
}

func (d BinarySensor) RawID() string {
	return "binary_sensor"
}
func (d BinarySensor) UniqueID() string {
	return d.uniqueID
}
func (d *BinarySensor) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type BinarySensor struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	DeviceClass            string `json:"device_class"`
	EnabledByDefault       bool   `json:"enabled_by_default"`
	Encoding               string `json:"encoding"`
	EntityCategory         string `json:"entity_category"`
	ExpireAfter            int    `json:"expire_after"`
	ForceUpdate            bool   `json:"force_update"`
	Icon                   string `json:"icon"`
	JsonAttributesTemplate string `json:"json_attributes_template"`
	JsonAttributesTopic    string `json:"json_attributes_topic"`
	Name                   string `json:"name"`
	ObjectId               string `json:"object_id"`
	OffDelay               int    `json:"off_delay"`
	PayloadAvailable       string `json:"payload_available"`
	PayloadNotAvailable    string `json:"payload_not_available"`
	PayloadOff             string `json:"payload_off"`
	PayloadOn              string `json:"payload_on"`
	Qos                    int    `json:"qos"`
	StateTopic             string `json:"state_topic"`
	UniqueId               string `json:"unique_id"`
	ValueTemplate          string `json:"value_template"`
	uniqueID               string
	rawID                  string
}

func (d Camera) RawID() string {
	return "camera"
}
func (d Camera) UniqueID() string {
	return d.uniqueID
}
func (d *Camera) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Camera struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	EnabledByDefault       bool   `json:"enabled_by_default"`
	EntityCategory         string `json:"entity_category"`
	Icon                   string `json:"icon"`
	JsonAttributesTemplate string `json:"json_attributes_template"`
	JsonAttributesTopic    string `json:"json_attributes_topic"`
	Name                   string `json:"name"`
	ObjectId               string `json:"object_id"`
	Topic                  string `json:"topic"`
	UniqueId               string `json:"unique_id"`
	uniqueID               string
	rawID                  string
}

func (d Cover) RawID() string {
	return "cover"
}
func (d Cover) UniqueID() string {
	return d.uniqueID
}
func (d *Cover) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Cover struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	CommandTopic         string `json:"command_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	DeviceClass            string `json:"device_class"`
	EnabledByDefault       bool   `json:"enabled_by_default"`
	Encoding               string `json:"encoding"`
	EntityCategory         string `json:"entity_category"`
	Icon                   string `json:"icon"`
	JsonAttributesTemplate string `json:"json_attributes_template"`
	JsonAttributesTopic    string `json:"json_attributes_topic"`
	Name                   string `json:"name"`
	ObjectId               string `json:"object_id"`
	Optimistic             bool   `json:"optimistic"`
	PayloadAvailable       string `json:"payload_available"`
	PayloadClose           string `json:"payload_close"`
	PayloadNotAvailable    string `json:"payload_not_available"`
	PayloadOpen            string `json:"payload_open"`
	PayloadStop            string `json:"payload_stop"`
	PositionClosed         int    `json:"position_closed"`
	PositionOpen           int    `json:"position_open"`
	PositionTemplate       string `json:"position_template"`
	PositionTopic          string `json:"position_topic"`
	Qos                    int    `json:"qos"`
	Retain                 bool   `json:"retain"`
	SetPositionTemplate    string `json:"set_position_template"`
	SetPositionTopic       string `json:"set_position_topic"`
	StateClosed            string `json:"state_closed"`
	StateClosing           string `json:"state_closing"`
	StateOpen              string `json:"state_open"`
	StateOpening           string `json:"state_opening"`
	StateStopped           string `json:"state_stopped"`
	StateTopic             string `json:"state_topic"`
	TiltClosedValue        int    `json:"tilt_closed_value"`
	TiltCommandTemplate    string `json:"tilt_command_template"`
	TiltCommandTopic       string `json:"tilt_command_topic"`
	TiltMax                int    `json:"tilt_max"`
	TiltMin                int    `json:"tilt_min"`
	TiltOpenedValue        int    `json:"tilt_opened_value"`
	TiltOptimistic         bool   `json:"tilt_optimistic"`
	TiltStatusTemplate     string `json:"tilt_status_template"`
	TiltStatusTopic        string `json:"tilt_status_topic"`
	UniqueId               string `json:"unique_id"`
	ValueTemplate          string `json:"value_template"`
	uniqueID               string
	rawID                  string
}

func (d DeviceTracker) RawID() string {
	return "device_tracker"
}
func (d DeviceTracker) UniqueID() string {
	return d.uniqueID
}
func (d *DeviceTracker) PopulateDevice() {}

type DeviceTracker struct {
	Devices        []string `json:"devices"`
	PayloadHome    string   `json:"payload_home"`
	PayloadNotHome string   `json:"payload_not_home"`
	Qos            int      `json:"qos"`
	SourceType     string   `json:"source_type"`
	uniqueID       string
	rawID          string
}

func (d DeviceTrigger) RawID() string {
	return "device_trigger"
}
func (d DeviceTrigger) UniqueID() string {
	return d.uniqueID
}
func (d *DeviceTrigger) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type DeviceTrigger struct {
	AutomationType string `json:"automation_type"`
	Device         struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	Payload       string `json:"payload"`
	Qos           int    `json:"qos"`
	Subtype       string `json:"subtype"`
	Topic         string `json:"topic"`
	Type          string `json:"type"`
	ValueTemplate string `json:"value_template"`
	uniqueID      string
	rawID         string
}

func (d Fan) RawID() string {
	return "fan"
}
func (d Fan) UniqueID() string {
	return d.uniqueID
}
func (d *Fan) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Fan struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	CommandTemplate      string `json:"command_template"`
	CommandTopic         string `json:"command_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	EnabledByDefault           bool     `json:"enabled_by_default"`
	Encoding                   string   `json:"encoding"`
	EntityCategory             string   `json:"entity_category"`
	Icon                       string   `json:"icon"`
	JsonAttributesTemplate     string   `json:"json_attributes_template"`
	JsonAttributesTopic        string   `json:"json_attributes_topic"`
	Name                       string   `json:"name"`
	ObjectId                   string   `json:"object_id"`
	Optimistic                 bool     `json:"optimistic"`
	OscillationCommandTemplate string   `json:"oscillation_command_template"`
	OscillationCommandTopic    string   `json:"oscillation_command_topic"`
	OscillationStateTopic      string   `json:"oscillation_state_topic"`
	OscillationValueTemplate   string   `json:"oscillation_value_template"`
	PayloadAvailable           string   `json:"payload_available"`
	PayloadNotAvailable        string   `json:"payload_not_available"`
	PayloadOff                 string   `json:"payload_off"`
	PayloadOn                  string   `json:"payload_on"`
	PayloadOscillationOff      string   `json:"payload_oscillation_off"`
	PayloadOscillationOn       string   `json:"payload_oscillation_on"`
	PayloadResetPercentage     string   `json:"payload_reset_percentage"`
	PayloadResetPresetMode     string   `json:"payload_reset_preset_mode"`
	PercentageCommandTemplate  string   `json:"percentage_command_template"`
	PercentageCommandTopic     string   `json:"percentage_command_topic"`
	PercentageStateTopic       string   `json:"percentage_state_topic"`
	PercentageValueTemplate    string   `json:"percentage_value_template"`
	PresetModeCommandTemplate  string   `json:"preset_mode_command_template"`
	PresetModeCommandTopic     string   `json:"preset_mode_command_topic"`
	PresetModeStateTopic       string   `json:"preset_mode_state_topic"`
	PresetModeValueTemplate    string   `json:"preset_mode_value_template"`
	PresetModes                []string `json:"preset_modes"`
	Qos                        int      `json:"qos"`
	Retain                     bool     `json:"retain"`
	SpeedRangeMax              int      `json:"speed_range_max"`
	SpeedRangeMin              int      `json:"speed_range_min"`
	StateTopic                 string   `json:"state_topic"`
	StateValueTemplate         string   `json:"state_value_template"`
	UniqueId                   string   `json:"unique_id"`
	uniqueID                   string
	rawID                      string
}

func (d Humidifier) RawID() string {
	return "humidifier"
}
func (d Humidifier) UniqueID() string {
	return d.uniqueID
}
func (d *Humidifier) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Humidifier struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	CommandTemplate      string `json:"command_template"`
	CommandTopic         string `json:"command_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	DeviceClass                   string   `json:"device_class"`
	EnabledByDefault              bool     `json:"enabled_by_default"`
	Encoding                      string   `json:"encoding"`
	EntityCategory                string   `json:"entity_category"`
	Icon                          string   `json:"icon"`
	JsonAttributesTemplate        string   `json:"json_attributes_template"`
	JsonAttributesTopic           string   `json:"json_attributes_topic"`
	MaxHumidity                   int      `json:"max_humidity"`
	MinHumidity                   int      `json:"min_humidity"`
	ModeCommandTemplate           string   `json:"mode_command_template"`
	ModeCommandTopic              string   `json:"mode_command_topic"`
	ModeStateTemplate             string   `json:"mode_state_template"`
	ModeStateTopic                string   `json:"mode_state_topic"`
	Modes                         []string `json:"modes"`
	Name                          string   `json:"name"`
	ObjectId                      string   `json:"object_id"`
	Optimistic                    bool     `json:"optimistic"`
	PayloadAvailable              string   `json:"payload_available"`
	PayloadNotAvailable           string   `json:"payload_not_available"`
	PayloadOff                    string   `json:"payload_off"`
	PayloadOn                     string   `json:"payload_on"`
	PayloadResetHumidity          string   `json:"payload_reset_humidity"`
	PayloadResetMode              string   `json:"payload_reset_mode"`
	Qos                           int      `json:"qos"`
	Retain                        bool     `json:"retain"`
	StateTopic                    string   `json:"state_topic"`
	StateValueTemplate            string   `json:"state_value_template"`
	TargetHumidityCommandTemplate string   `json:"target_humidity_command_template"`
	TargetHumidityCommandTopic    string   `json:"target_humidity_command_topic"`
	TargetHumidityStateTemplate   string   `json:"target_humidity_state_template"`
	TargetHumidityStateTopic      string   `json:"target_humidity_state_topic"`
	UniqueId                      string   `json:"unique_id"`
	uniqueID                      string
	rawID                         string
}

func (d Climate) RawID() string {
	return "climate"
}
func (d Climate) UniqueID() string {
	return d.uniqueID
}
func (d *Climate) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Climate struct {
	ActionTemplate             string `json:"action_template"`
	ActionTopic                string `json:"action_topic"`
	AuxCommandTopic            string `json:"aux_command_topic"`
	AuxStateTemplate           string `json:"aux_state_template"`
	AuxStateTopic              string `json:"aux_state_topic"`
	AvailabilityMode           string `json:"availability_mode"`
	AvailabilityTemplate       string `json:"availability_template"`
	AvailabilityTopic          string `json:"availability_topic"`
	CurrentTemperatureTemplate string `json:"current_temperature_template"`
	CurrentTemperatureTopic    string `json:"current_temperature_topic"`
	Device                     struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	EnabledByDefault               bool     `json:"enabled_by_default"`
	Encoding                       string   `json:"encoding"`
	EntityCategory                 string   `json:"entity_category"`
	FanModeCommandTemplate         string   `json:"fan_mode_command_template"`
	FanModeCommandTopic            string   `json:"fan_mode_command_topic"`
	FanModeStateTemplate           string   `json:"fan_mode_state_template"`
	FanModeStateTopic              string   `json:"fan_mode_state_topic"`
	FanModes                       []string `json:"fan_modes"`
	Icon                           string   `json:"icon"`
	Initial                        int      `json:"initial"`
	JsonAttributesTemplate         string   `json:"json_attributes_template"`
	JsonAttributesTopic            string   `json:"json_attributes_topic"`
	MaxTemp                        float64  `json:"max_temp"`
	MinTemp                        float64  `json:"min_temp"`
	ModeCommandTemplate            string   `json:"mode_command_template"`
	ModeCommandTopic               string   `json:"mode_command_topic"`
	ModeStateTemplate              string   `json:"mode_state_template"`
	ModeStateTopic                 string   `json:"mode_state_topic"`
	Modes                          []string `json:"modes"`
	Name                           string   `json:"name"`
	ObjectId                       string   `json:"object_id"`
	PayloadAvailable               string   `json:"payload_available"`
	PayloadNotAvailable            string   `json:"payload_not_available"`
	PayloadOff                     string   `json:"payload_off"`
	PayloadOn                      string   `json:"payload_on"`
	PowerCommandTopic              string   `json:"power_command_topic"`
	Precision                      float64  `json:"precision"`
	PresetModeCommandTemplate      string   `json:"preset_mode_command_template"`
	PresetModeCommandTopic         string   `json:"preset_mode_command_topic"`
	PresetModeStateTopic           string   `json:"preset_mode_state_topic"`
	PresetModeValueTemplate        string   `json:"preset_mode_value_template"`
	PresetModes                    []string `json:"preset_modes"`
	Qos                            int      `json:"qos"`
	Retain                         bool     `json:"retain"`
	SwingModeCommandTemplate       string   `json:"swing_mode_command_template"`
	SwingModeCommandTopic          string   `json:"swing_mode_command_topic"`
	SwingModeStateTemplate         string   `json:"swing_mode_state_template"`
	SwingModeStateTopic            string   `json:"swing_mode_state_topic"`
	SwingModes                     []string `json:"swing_modes"`
	TempStep                       float64  `json:"temp_step"`
	TemperatureCommandTemplate     string   `json:"temperature_command_template"`
	TemperatureCommandTopic        string   `json:"temperature_command_topic"`
	TemperatureHighCommandTemplate string   `json:"temperature_high_command_template"`
	TemperatureHighCommandTopic    string   `json:"temperature_high_command_topic"`
	TemperatureHighStateTemplate   string   `json:"temperature_high_state_template"`
	TemperatureHighStateTopic      string   `json:"temperature_high_state_topic"`
	TemperatureLowCommandTemplate  string   `json:"temperature_low_command_template"`
	TemperatureLowCommandTopic     string   `json:"temperature_low_command_topic"`
	TemperatureLowStateTemplate    string   `json:"temperature_low_state_template"`
	TemperatureLowStateTopic       string   `json:"temperature_low_state_topic"`
	TemperatureStateTemplate       string   `json:"temperature_state_template"`
	TemperatureStateTopic          string   `json:"temperature_state_topic"`
	TemperatureUnit                string   `json:"temperature_unit"`
	UniqueId                       string   `json:"unique_id"`
	ValueTemplate                  string   `json:"value_template"`
	uniqueID                       string
	rawID                          string
}

func (d Light) RawID() string {
	return "light"
}
func (d Light) UniqueID() string {
	return d.uniqueID
}
func (d *Light) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Light struct {
	AvailabilityMode          string `json:"availability_mode"`
	AvailabilityTemplate      string `json:"availability_template"`
	AvailabilityTopic         string `json:"availability_topic"`
	BrightnessCommandTemplate string `json:"brightness_command_template"`
	BrightnessCommandTopic    string `json:"brightness_command_topic"`
	BrightnessScale           int    `json:"brightness_scale"`
	BrightnessStateTopic      string `json:"brightness_state_topic"`
	BrightnessValueTemplate   string `json:"brightness_value_template"`
	ColorModeStateTopic       string `json:"color_mode_state_topic"`
	ColorModeValueTemplate    string `json:"color_mode_value_template"`
	ColorTempCommandTemplate  string `json:"color_temp_command_template"`
	ColorTempCommandTopic     string `json:"color_temp_command_topic"`
	ColorTempStateTopic       string `json:"color_temp_state_topic"`
	ColorTempValueTemplate    string `json:"color_temp_value_template"`
	CommandTopic              string `json:"command_topic"`
	Device                    struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	EffectCommandTemplate  string   `json:"effect_command_template"`
	EffectCommandTopic     string   `json:"effect_command_topic"`
	EffectList             []string `json:"effect_list"`
	EffectStateTopic       string   `json:"effect_state_topic"`
	EffectValueTemplate    string   `json:"effect_value_template"`
	EnabledByDefault       bool     `json:"enabled_by_default"`
	Encoding               string   `json:"encoding"`
	EntityCategory         string   `json:"entity_category"`
	HsValueTemplate        string   `json:"hs_value_template"`
	Icon                   string   `json:"icon"`
	JsonAttributesTemplate string   `json:"json_attributes_template"`
	JsonAttributesTopic    string   `json:"json_attributes_topic"`
	MaxMireds              int      `json:"max_mireds"`
	MinMireds              int      `json:"min_mireds"`
	Name                   string   `json:"name"`
	ObjectId               string   `json:"object_id"`
	OnCommandType          string   `json:"on_command_type"`
	Optimistic             bool     `json:"optimistic"`
	PayloadAvailable       string   `json:"payload_available"`
	PayloadNotAvailable    string   `json:"payload_not_available"`
	PayloadOff             string   `json:"payload_off"`
	PayloadOn              string   `json:"payload_on"`
	Qos                    int      `json:"qos"`
	Retain                 bool     `json:"retain"`
	RgbCommandTemplate     string   `json:"rgb_command_template"`
	RgbCommandTopic        string   `json:"rgb_command_topic"`
	RgbStateTopic          string   `json:"rgb_state_topic"`
	RgbValueTemplate       string   `json:"rgb_value_template"`
	Schema                 string   `json:"schema"`
	StateTopic             string   `json:"state_topic"`
	StateValueTemplate     string   `json:"state_value_template"`
	UniqueId               string   `json:"unique_id"`
	WhiteCommandTopic      string   `json:"white_command_topic"`
	WhiteScale             int      `json:"white_scale"`
	XyCommandTopic         string   `json:"xy_command_topic"`
	XyStateTopic           string   `json:"xy_state_topic"`
	XyValueTemplate        string   `json:"xy_value_template"`
	uniqueID               string
	rawID                  string
}

func (d Lock) RawID() string {
	return "lock"
}
func (d Lock) UniqueID() string {
	return d.uniqueID
}
func (d *Lock) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Lock struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	CommandTopic         string `json:"command_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	EnabledByDefault       bool   `json:"enabled_by_default"`
	Encoding               string `json:"encoding"`
	EntityCategory         string `json:"entity_category"`
	Icon                   string `json:"icon"`
	JsonAttributesTemplate string `json:"json_attributes_template"`
	JsonAttributesTopic    string `json:"json_attributes_topic"`
	Name                   string `json:"name"`
	ObjectId               string `json:"object_id"`
	Optimistic             bool   `json:"optimistic"`
	PayloadAvailable       string `json:"payload_available"`
	PayloadLock            string `json:"payload_lock"`
	PayloadNotAvailable    string `json:"payload_not_available"`
	PayloadOpen            string `json:"payload_open"`
	PayloadUnlock          string `json:"payload_unlock"`
	Qos                    int    `json:"qos"`
	Retain                 bool   `json:"retain"`
	StateLocked            string `json:"state_locked"`
	StateTopic             string `json:"state_topic"`
	StateUnlocked          string `json:"state_unlocked"`
	UniqueId               string `json:"unique_id"`
	ValueTemplate          string `json:"value_template"`
	uniqueID               string
	rawID                  string
}

func (d Number) RawID() string {
	return "number"
}
func (d Number) UniqueID() string {
	return d.uniqueID
}
func (d *Number) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Number struct {
	AvailabilityMode  string `json:"availability_mode"`
	AvailabilityTopic string `json:"availability_topic"`
	CommandTemplate   string `json:"command_template"`
	CommandTopic      string `json:"command_topic"`
	Device            struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	EnabledByDefault       bool    `json:"enabled_by_default"`
	Encoding               string  `json:"encoding"`
	EntityCategory         string  `json:"entity_category"`
	Icon                   string  `json:"icon"`
	JsonAttributesTemplate string  `json:"json_attributes_template"`
	JsonAttributesTopic    string  `json:"json_attributes_topic"`
	Max                    float64 `json:"max"`
	Min                    float64 `json:"min"`
	Name                   string  `json:"name"`
	ObjectId               string  `json:"object_id"`
	Optimistic             bool    `json:"optimistic"`
	PayloadReset           string  `json:"payload_reset"`
	Qos                    int     `json:"qos"`
	Retain                 bool    `json:"retain"`
	StateTopic             string  `json:"state_topic"`
	Step                   float64 `json:"step"`
	UniqueId               string  `json:"unique_id"`
	UnitOfMeasurement      string  `json:"unit_of_measurement"`
	ValueTemplate          string  `json:"value_template"`
	uniqueID               string
	rawID                  string
}

func (d Scene) RawID() string {
	return "scene"
}
func (d Scene) UniqueID() string {
	return d.uniqueID
}
func (d *Scene) PopulateDevice() {}

type Scene struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	CommandTopic         string `json:"command_topic"`
	EnabledByDefault     bool   `json:"enabled_by_default"`
	EntityCategory       string `json:"entity_category"`
	Icon                 string `json:"icon"`
	Name                 string `json:"name"`
	ObjectId             string `json:"object_id"`
	PayloadAvailable     string `json:"payload_available"`
	PayloadNotAvailable  string `json:"payload_not_available"`
	PayloadOn            string `json:"payload_on"`
	Qos                  int    `json:"qos"`
	Retain               bool   `json:"retain"`
	UniqueId             string `json:"unique_id"`
	uniqueID             string
	rawID                string
}

func (d Select) RawID() string {
	return "select"
}
func (d Select) UniqueID() string {
	return d.uniqueID
}
func (d *Select) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Select struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	CommandTemplate      string `json:"command_template"`
	CommandTopic         string `json:"command_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	EnabledByDefault       bool     `json:"enabled_by_default"`
	Encoding               string   `json:"encoding"`
	EntityCategory         string   `json:"entity_category"`
	Icon                   string   `json:"icon"`
	JsonAttributesTemplate string   `json:"json_attributes_template"`
	JsonAttributesTopic    string   `json:"json_attributes_topic"`
	Name                   string   `json:"name"`
	ObjectId               string   `json:"object_id"`
	Optimistic             bool     `json:"optimistic"`
	Options                []string `json:"options"`
	Qos                    int      `json:"qos"`
	Retain                 bool     `json:"retain"`
	StateTopic             string   `json:"state_topic"`
	UniqueId               string   `json:"unique_id"`
	ValueTemplate          string   `json:"value_template"`
	uniqueID               string
	rawID                  string
}

func (d Sensor) RawID() string {
	return "sensor"
}
func (d Sensor) UniqueID() string {
	return d.uniqueID
}
func (d *Sensor) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Sensor struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	DeviceClass            string `json:"device_class"`
	EnabledByDefault       bool   `json:"enabled_by_default"`
	Encoding               string `json:"encoding"`
	EntityCategory         string `json:"entity_category"`
	ExpireAfter            int    `json:"expire_after"`
	ForceUpdate            bool   `json:"force_update"`
	Icon                   string `json:"icon"`
	JsonAttributesTemplate string `json:"json_attributes_template"`
	JsonAttributesTopic    string `json:"json_attributes_topic"`
	LastResetValueTemplate string `json:"last_reset_value_template"`
	Name                   string `json:"name"`
	ObjectId               string `json:"object_id"`
	PayloadAvailable       string `json:"payload_available"`
	PayloadNotAvailable    string `json:"payload_not_available"`
	Qos                    int    `json:"qos"`
	StateClass             string `json:"state_class"`
	StateTopic             string `json:"state_topic"`
	UniqueId               string `json:"unique_id"`
	UnitOfMeasurement      string `json:"unit_of_measurement"`
	ValueTemplate          string `json:"value_template"`
	uniqueID               string
	rawID                  string
}

func (d Switch) RawID() string {
	return "switch"
}
func (d Switch) UniqueID() string {
	return d.uniqueID
}
func (d *Switch) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Switch struct {
	AvailabilityMode     string `json:"availability_mode"`
	AvailabilityTemplate string `json:"availability_template"`
	AvailabilityTopic    string `json:"availability_topic"`
	CommandTopic         string `json:"command_topic"`
	Device               struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	DeviceClass            string `json:"device_class"`
	EnabledByDefault       bool   `json:"enabled_by_default"`
	Encoding               string `json:"encoding"`
	EntityCategory         string `json:"entity_category"`
	Icon                   string `json:"icon"`
	JsonAttributesTemplate string `json:"json_attributes_template"`
	JsonAttributesTopic    string `json:"json_attributes_topic"`
	Name                   string `json:"name"`
	ObjectId               string `json:"object_id"`
	Optimistic             bool   `json:"optimistic"`
	PayloadAvailable       string `json:"payload_available"`
	PayloadNotAvailable    string `json:"payload_not_available"`
	PayloadOff             string `json:"payload_off"`
	PayloadOn              string `json:"payload_on"`
	Qos                    int    `json:"qos"`
	Retain                 bool   `json:"retain"`
	StateOff               string `json:"state_off"`
	StateOn                string `json:"state_on"`
	StateTopic             string `json:"state_topic"`
	UniqueId               string `json:"unique_id"`
	ValueTemplate          string `json:"value_template"`
	uniqueID               string
	rawID                  string
}

func (d Tag) RawID() string {
	return "tag"
}
func (d Tag) UniqueID() string {
	return d.uniqueID
}
func (d *Tag) PopulateDevice() {
	d.Device.Manufacturer = Manufacturer
	d.Device.Model = SoftwareName
	d.Device.Name = InstanceName
	d.Device.SwVersion = SWVersion
}

type Tag struct {
	Device struct {
		ConfigurationUrl string   `json:"configuration_url"`
		Connections      []string `json:"connections"`
		Identifiers      []string `json:"identifiers"`
		Manufacturer     string   `json:"manufacturer"`
		Model            string   `json:"model"`
		Name             string   `json:"name"`
		SuggestedArea    string   `json:"suggested_area"`
		SwVersion        string   `json:"sw_version"`
		ViaDevice        string   `json:"via_device"`
	} `json:"device"`
	Topic         string `json:"topic"`
	ValueTemplate string `json:"value_template"`
	uniqueID      string
	rawID         string
}

func (d Vacuum) RawID() string {
	return "vacuum"
}
func (d Vacuum) UniqueID() string {
	return d.uniqueID
}
func (d *Vacuum) PopulateDevice() {}

type Vacuum struct {
	AvailabilityMode       string   `json:"availability_mode"`
	AvailabilityTemplate   string   `json:"availability_template"`
	AvailabilityTopic      string   `json:"availability_topic"`
	BatteryLevelTemplate   string   `json:"battery_level_template"`
	BatteryLevelTopic      string   `json:"battery_level_topic"`
	ChargingTemplate       string   `json:"charging_template"`
	ChargingTopic          string   `json:"charging_topic"`
	CleaningTemplate       string   `json:"cleaning_template"`
	CleaningTopic          string   `json:"cleaning_topic"`
	CommandTopic           string   `json:"command_topic"`
	DockedTemplate         string   `json:"docked_template"`
	DockedTopic            string   `json:"docked_topic"`
	EnabledByDefault       bool     `json:"enabled_by_default"`
	Encoding               string   `json:"encoding"`
	EntityCategory         string   `json:"entity_category"`
	ErrorTemplate          string   `json:"error_template"`
	ErrorTopic             string   `json:"error_topic"`
	FanSpeedList           []string `json:"fan_speed_list"`
	FanSpeedTemplate       string   `json:"fan_speed_template"`
	FanSpeedTopic          string   `json:"fan_speed_topic"`
	Icon                   string   `json:"icon"`
	JsonAttributesTemplate string   `json:"json_attributes_template"`
	JsonAttributesTopic    string   `json:"json_attributes_topic"`
	Name                   string   `json:"name"`
	ObjectId               string   `json:"object_id"`
	PayloadAvailable       string   `json:"payload_available"`
	PayloadCleanSpot       string   `json:"payload_clean_spot"`
	PayloadLocate          string   `json:"payload_locate"`
	PayloadNotAvailable    string   `json:"payload_not_available"`
	PayloadReturnToBase    string   `json:"payload_return_to_base"`
	PayloadStartPause      string   `json:"payload_start_pause"`
	PayloadStop            string   `json:"payload_stop"`
	PayloadTurnOff         string   `json:"payload_turn_off"`
	PayloadTurnOn          string   `json:"payload_turn_on"`
	Qos                    int      `json:"qos"`
	Retain                 bool     `json:"retain"`
	Schema                 string   `json:"schema"`
	SendCommandTopic       string   `json:"send_command_topic"`
	SetFanSpeedTopic       string   `json:"set_fan_speed_topic"`
	SupportedFeatures      []string `json:"supported_features"`
	UniqueId               string   `json:"unique_id"`
	uniqueID               string
	rawID                  string
}
type Device interface {
	AlarmControlPanel | BinarySensor | Camera | Cover | DeviceTracker | DeviceTrigger | Fan | Humidifier | Climate | Light | Lock | Number | Scene | Select | Sensor | Switch | Tag | Vacuum
	RawID() string
	UniqueID() string
	PopulateDevice()
}
