package main

import "github.com/markbates/pkger"

func main() {
	pkger.Include("/templates")
	metadata := RobotProjectMetadata{
		Name:    "Test Transcribed Project",
		Year:    2020,
		Team: 	 100,
		Package: "frc.robot",
		GradleRioV: "2020.3.2",
	}
	project := RobotProject{
		Metadata: &metadata,
		OutputDir: "",
		Robot: Robot{
			Subsystems:    Subsystems{
				Package:       metadata.Package+".subsystems",
				SubsystemList: []*Subsystem{
					&Subsystem{
						Name:          "Drivetrain",
						SubsystemType: 0,
						OnStateUpdateCode: "System.out.println(\"Hello\");\nSystem.out.println(\"1\");",
						Motors:        []*BaseMotorController{&BaseMotorController{
							Name:                     "TestMotorController",
							MotorControllerType:      MOTOR_FRC_TALONFX,
							CanID:                    0,
							EnableFollow:             false,
							EnableInvert:             true,
							Invert:                   INVERT_OPPOSE_MASTER,
							EnableNeutral:            true,
							Neutral:                  NEUTRAL_BRAKE,
							NeutralEnum:              "",
							EnableFeedback:           false,
							FeedbackPort:             0,
							FeedbackPortEnum:         "",
							EnableSensorPhase:        false,
							SensorPhase:              0,
							SensorPhaseEnum:          "",
							Timeout:                  0,
							PeakOutputForward:        0,
							PeakOutputReverse:        0,
							PID:                      &PIDController{
								KP: 0,
								KI: 0,
								KD: 0,
								KF: 0,
							},
						}},
						Sensors:       nil,
					},
				},
			},
			Sensors:       nil,
			GlobalManager: nil,
		},
	}

	GenerateProject(&project)

	//TODO Add gradle files and vendordeps
}
