package main

import (
	"log"
	"os"
)
func CreateProjectRoot(project *RobotProject){
	CreateMainJava(project)
	CreateRobotJava(project)
}
func CreateMainJava(project *RobotProject){
	fileName := "main.java"
	inputFile := "/templates/project/main.tmpl.java"
	outputFile := project.OutputDir+"/Main.java"
	type LocalParse struct{
		Package string
	}
	parseSpec := LocalParse{
		Package: project.Metadata.Package,
	}

	f, err := os.Create(outputFile)
	if err != nil {
		log.Println("Error while creating "+fileName + " : ", err)
		os.Exit(1004)
	}
	t, e := CompileTemplate(inputFile)
	if e != nil {
		log.Println("Could not parse template "+inputFile)
		log.Println(err)
		os.Exit(2001)
	}
	err = t.Execute(f, parseSpec)

	if err != nil {
		log.Println("Could not save "+outputFile+" to filesystem: ", err)
		os.Exit(1005)
	}
	log.Println("Created "+fileName)
}

func CreateRobotJava(project *RobotProject){
	fileName := "main.java"
	inputFile := "/templates/project/robot.tmpl.java"
	outputFile := project.OutputDir+"/Robot.java"
	type LocalParse struct{
		Package string
	}
	parseSpec := LocalParse{
		Package: project.Metadata.Package,
	}

	f, err := os.Create(outputFile)
	if err != nil {
		log.Println("Error while creating "+fileName + " : ", err)
		os.Exit(1004)
	}
	t, e := CompileTemplate(inputFile)
	if e != nil {
		log.Println("Could not parse template "+inputFile)
		log.Println(err)
		os.Exit(2001)
	}
	err = t.Execute(f, parseSpec)

	if err != nil {
		log.Println("Could not save "+outputFile+" to filesystem: ", err)
		os.Exit(1005)
	}
	log.Println("Created "+fileName)
}


