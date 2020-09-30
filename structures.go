package main

type RobotProjectMetadata struct{
	Name string `toml:"name"`
	Year int `toml:"year"`
	Team int `toml:"team"`
	Package string `toml:"package"`
	GradleRioV string `toml:"gradle_rio_version"`

}
type RobotProject struct{
	Metadata *RobotProjectMetadata `toml:"metadata"`
	OutputDir string `toml:"output_directory"`
	Robot Robot `toml:"robot"`
}

/////////// Robot
type Robot struct{
	Path string
	Subsystems Subsystems `toml:"subsystem"`
	Sensors []*Sensor `toml:"sensors"`
	GlobalManager *GlobalManager `toml:"gloal_manager"`

}

type GlobalManager struct{
	Path string `toml:"path"`
}

type GlobalManagerTrigger struct{
	Name string `toml:"name"`
	Condition string `toml:"condition"`
}

/////////// Subsystems

type Subsystems struct{
	Package string `toml:"package"`
	SubsystemList []*Subsystem `toml:"subsystem_list"`
}

const(
	SUBSYSTEM_OPENLOOP = iota
	SUBSYSTEM_LINEAR_VELOCITY
	SUBSYSTEM_LINEAR_POSITION
	SUBSYSTEM_DRIVETRAIN
)
type Subsystem struct{
	Path string `toml:""package`
	Name string `toml:"name"`
	Summary string `toml:"summary"`
	Description string `toml:"description"`
	SubsystemType int `toml:"subsystem_type"`
	Motors []*BaseMotorController `toml:"motor"`
	Sensors []*Sensor `toml:"sensors"`


	OnStateUpdateCode string `toml:"on_state_update_code"`


}

const(
	SENSOR_TYPE_ANALOG = iota
	SENSOR_TYPE_DIGITAL
)

const(
	SENSOR_PORT_TYPE_DIRECT_FEEDBACK = iota
	SENSOR_PORT_TYPE_ROBORIO_DIO
	SENSOR_PORT_TYPE_ROBORIO_ANALOG
)
type Sensor struct{
	Path string `toml:"path"`
	Name string `toml:"name"`
	SensorType int `toml:"sensor_type"`
	Port int `toml:"port"`
}

const(
	MOTOR_FRC_TALONFX = iota
	MOTOR_FRC_TALONSRX
	MOTOR_FRC_VICTORSPX
)

type PIDController struct{
	KP float64 `toml:"kP"`
	KI float64 `toml:"kI"`
	KD float64 `toml:"kD"`
	KF float64 `toml:"kF"`
}

const(
	INVERT_NONE = iota
	INVERT_INVERT_MOTOR_OUTPUT
	INVERT_FOLLOW_MASTER
	INVERT_OPPOSE_MASTER
)

const(
	NEUTRAL_EEPROMSETTING = iota
	NEUTRAL_COAST
	NEUTRAL_BRAKE
)
type BaseMotorController struct{
	Path string `toml:"path"`
	Name string `toml:"name"`

	MotorControllerType int `toml:"motor_controller_type"`
	MotorControllerClassName string `toml:"motor_controller_class_name"`
	CanID int `toml:"can_id"`

	EnableFollow bool `toml:"enable_follow"`
	Follows string `toml:"follows"`//name of master

	EnableInvert bool `toml:"enable_invert"`
	Invert int `toml:"invert"`
	InvertEnum string `toml:"-"`

	EnableNeutral bool `toml:"enable_neutral"`
	Neutral int `toml:"neutral"`
	NeutralEnum string `toml:"-"`

	EnableFeedback bool `toml:"enable_feedback"`
	FeedbackPort int `toml:"feedback_port"`
	FeedbackPortEnum string `toml:"-"`

	EnableSensorPhase bool `toml:"enable_sensor_phase"`
	SensorPhase int `toml:"sensor_phase"`
	SensorPhaseEnum string `toml:"-"`

	TimeoutEnabled int `toml:"timeout_enabled"`
	Timeout int `toml:"timeout"`

	EnableCurrentLimit bool `toml:"enable_current_limit"`

	EnableOpenLoopRamp bool `toml:"enable_open_loop_ramp"`
	OpenLoopRamp float64 `toml:"open_loop_ramp"`

	EnableNominalOutput bool `toml:"enable_nominal_output"`
	NominalOutputForward float64 `toml:"nominal_output_forward"`
	NominalOutputReverse float64 `toml:"nominal_output_reverse"`

	EnablePeakOutput bool `toml:"enable_peak_output"`
	PeakOutputForward float64 `toml:"peak_output_forward"`
	PeakOutputReverse float64 `toml:"peak_output_reverse"`

	EnablePID bool `toml:"enable_pid"`
	PID *PIDController `toml:"pid"`
	//TODO NeutralMode, Inverted

}

/////////////// Commands


type FileParse struct{
	Package string
	Body interface{}
}

