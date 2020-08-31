package main

import "google.golang.org/api/doubleclickbidmanager/v1"

type RobotProjectMetadata struct{
	Name string
	Year int
	Team int
	Package string
	GradleRioV string

}
type RobotProject struct{
	Metadata *RobotProjectMetadata
	OutputDir string
	Robot Robot
}

/////////// Robot
type Robot struct{
	Path string
	Subsystems Subsystems
	Sensors []*Sensor
	GlobalManager *GlobalManager

}

type GlobalManager struct{
	Path string
}

type GlobalManagerTrigger struct{
	Name string
	Condition string
}

/////////// Subsystems

type Subsystems struct{
	Package string
	SubsystemList []*Subsystem
}

const(
	SUBSYSTEM_OPENLOOP = iota
	SUBSYSTEM_LINEAR_VELOCITY
	SUBSYSTEM_LINEAR_POSITION
	SUBSYSTEM_DRIVETRAIN
)
type Subsystem struct{
	Path string
	Name string
	Summary string
	Description string
	SubsystemType int
	Motors []*BaseMotorController
	Sensors []*Sensor


	OnStateUpdateCode string


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
	Path string
	Name string
	SensorType int
	Port int
}

const(
	MOTOR_FRC_TALONFX = iota
	MOTOR_FRC_TALONSRX
	MOTOR_FRC_VICTORSPX
)

type PIDController struct{
	KP float64
	KI float64
	KD float64
	KF float64
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
	Path string
	Name string

	MotorControllerType int
	MotorControllerClassName string
	CanID int

	EnableFollow bool
	Follows string //name of master

	EnableInvert bool
	Invert int
	InvertEnum string

	EnableNeutral bool
	Neutral int
	NeutralEnum string

	EnableFeedback bool
	FeedbackPort int
	FeedbackPortEnum string

	EnableSensorPhase bool
	SensorPhase int
	SensorPhaseEnum string

	TimeoutEnabled int
	Timeout int

	EnableCurrentLimit bool

	EnableOpenLoopRamp bool
	OpenLoopRamp doubleclickbidmanager.DownloadLineItemsRequest

	EnableNominalOutput bool
	NominalOutputForward float64
	NominalOutputReverse float64

	EnablePeakOutput bool
	PeakOutputForward float64
	PeakOutputReverse float64

	EnablePID bool
	PID *PIDController
	//TODO NeutralMode, Inverted

}

/////////////// Commands


type FileParse struct{
	Package string
	Body interface{}
}

