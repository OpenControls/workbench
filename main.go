package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/markbates/pkger"
	"log"
	"os"
)

func main() {
	var config RobotProject
	if _, err := toml.DecodeFile("robot.ocwb.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	json, err := json.Marshal(config)
	if(err != nil){
		fmt.Println(err)
		return
	}
	fmt.Println(string(json))
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
						Motors:        []*BaseMotorController{{
							Name:                "TestMotorController",
							MotorControllerType: MOTOR_FRC_TALONFX,
							CanID:               0,
							EnableFollow:        false,
							EnableInvert:        true,
							Invert:              INVERT_OPPOSE_MASTER,
							EnableNeutral:       true,
							Neutral:             NEUTRAL_BRAKE,
							NeutralEnum:         "",
							EnableFeedback:      false,
							FeedbackPort:        0,
							FeedbackPortEnum:    "",
							EnableSensorPhase:   false,
							SensorPhase:         0,
							SensorPhaseEnum:     "",
							Timeout:             0,
							PeakOutputForward:   0,
							PeakOutputReverse:   0,
							PID: &PIDController{
														KP: 0,
														KI: 0,
														KD: 0,
														KF: 0,
													},
												}},
						Sensors:       []*Sensor{
							&Sensor{
								Name:       "Test Sensor",
								SensorType: 0,
								Port:       0,
							},
						},
					},
					&Subsystem{
						Name:          "Lift",
						SubsystemType: 1,
						OnStateUpdateCode: "System.out.println(\"Hello\");\nSystem.out.println(\"1\");",
						Motors:        []*BaseMotorController{{
							Name:                "TestMotorController",
							MotorControllerType: MOTOR_FRC_TALONFX,
							CanID:               1,
							EnableFollow:        false,
							EnableInvert:        true,
							Invert:              INVERT_OPPOSE_MASTER,
							EnableNeutral:       true,
							Neutral:             NEUTRAL_BRAKE,
							NeutralEnum:         "",
							EnableFeedback:      false,
							FeedbackPort:        0,
							FeedbackPortEnum:    "",
							EnableSensorPhase:   false,
							SensorPhase:         0,
							SensorPhaseEnum:     "",
							Timeout:             0,
							PeakOutputForward:   0,
							PeakOutputReverse:   0,
							PID: &PIDController{
								KP: 0,
								KI: 0,
								KD: 0,
								KF: 0,
							},
						}},
						Sensors:       []*Sensor{
							&Sensor{
								Name:       "Test Sensor",
								SensorType: 0,
								Port:       0,
							},
						},
					},
				},
			},
			Sensors:       []*Sensor{
				&Sensor{
					Name:       "Test Global Sensor",
					SensorType: 1,
					Port:       1,
				},
			},
			GlobalManager: nil,
		},
	}

	f, err := os.Create("export.toml")
	if err != nil {
		// failed to create/open the file
		log.Fatal(err)
	}
	if err := toml.NewEncoder(f).Encode(project); err != nil {
		// failed to encode
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		// failed to close the file
		log.Fatal(err)

	}



	GenerateProject(&project)

	//TODO Add gradle files and vendordeps
}
