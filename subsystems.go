package main

import (
	"log"
	"os"
)

func ProcessSubsystems(project *RobotProject) {
	CreateSubsystemDir(project)
	for i:=0; i < len(project.Robot.Subsystems.SubsystemList); i++{
		subsys := project.Robot.Subsystems.SubsystemList[i]
		CreateSubsystem(project, subsys)
	}
}

func CreateSubsystemDir(project *RobotProject) {
	CreateDirectory(project, "subsystems")
}

func CreateSubsystem(project *RobotProject, subsystem *Subsystem) {
	fileName := "main.java"
	inputFile := "/templates/project/subsystem.tmpl.java"
	outputFile := project.OutputDir + "/subsystems/"+subsystem.Name+".java"



	for i := 0; i < len(subsystem.Motors); i++ {
		motor := subsystem.Motors[i]

		if motor.MotorControllerType == MOTOR_FRC_TALONFX{
			motor.MotorControllerClassName = "FRCTalonFX"
		}else if motor.MotorControllerType == MOTOR_FRC_TALONSRX{
			motor.MotorControllerClassName = "FRCTalonSRX"
		}else if motor.MotorControllerType == MOTOR_FRC_VICTORSPX{
			motor.MotorControllerClassName = "FRCVictorSPX"
		}

		if motor.Neutral == NEUTRAL_COAST {
			motor.NeutralEnum = "NeutralMode.Coast"
		} else if motor.Neutral == NEUTRAL_BRAKE {
			motor.NeutralEnum = "NeutralMode.Brake"
		} else {
			motor.Neutral = NEUTRAL_BRAKE
			motor.NeutralEnum = "NeutralMode.Brake"

		}

		if motor.EnableInvert {
			if motor.Invert == INVERT_NONE {
				motor.InvertEnum = "InvertType.None"
			} else if motor.Invert == INVERT_INVERT_MOTOR_OUTPUT {
				motor.InvertEnum = "InvertType.InvertMotorOutput"
			} else if motor.Invert == INVERT_FOLLOW_MASTER {
				motor.InvertEnum = "InvertType.FollowMaster"
			} else if motor.Invert == INVERT_OPPOSE_MASTER {
				motor.InvertEnum = "InvertType.OpposeMaster"
			}
		}

	}
	type LocalParse struct {
		Package string
		Class   string
		Motors  []*BaseMotorController
		Subsystem *Subsystem
	}
	parseSpec := LocalParse{
		Package: project.Metadata.Package + ".subsystems",
		Class:   subsystem.Name,
		Motors:  subsystem.Motors,
		Subsystem: subsystem,
	}

	f, err := os.Create(outputFile)
	if err != nil {
		log.Println("Error while creating "+fileName+" : ", err)
		os.Exit(1004)
	}
	t, e := CompileTemplate(inputFile)
	if e != nil {
		log.Println("Could not parse template " + inputFile)
		log.Println(err)
		os.Exit(2001)
	}
	err = t.Execute(f, parseSpec)

	if err != nil {
		log.Println("Could not save "+outputFile+" to filesystem: ", err)
		os.Exit(1005)
	}
	log.Println("Created " + fileName)
}
