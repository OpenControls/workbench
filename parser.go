package main

import (
	"log"
	"os"
	"strings"
)
var COMPILE_ROOT = ""
func GenerateProject(project *RobotProject){
	COMPILE_ROOT = PWD()
	CreateProjectDirectory(project)
	CreateRepoRootFiles(project)
	CreateDotDirectories(project)
	CreateProjectRoot(project)
	ProcessSubsystems(project)

}

func EvaluatePackageAsFileDirectory(metadata *RobotProjectMetadata) string{

	raw := metadata.Package
	converted := strings.Replace(raw, ".", "/",-1)
	return converted
}


func PWD() string{
	path, err := os.Getwd()
	if err != nil {
		log.Println("Error identifying path in parser.go -> PWD()")
		log.Println(err)
		os.Exit(1001)
	}
	return path
}

func CreateDirectory(project *RobotProject,path string){
	dir := project.OutputDir+"/"+path
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil{
		log.Println("Error creating new Directory "+ dir)
		log.Println(err)
		os.Exit(1003)
	}
	log.Println("Successfully created "+ dir)
}
func CreateProjectDirectory(project *RobotProject){
	err := os.Chdir(PWD())
	if err != nil{
		log.Println("Error changing directories to "+PWD())
		log.Println(err)
		os.Exit(1002)
	}

	projectOutput := PWD() + "/build/src/main/java/"

	projectRoot := projectOutput + EvaluatePackageAsFileDirectory(project.Metadata)
	err = os.MkdirAll(projectRoot, os.ModePerm)
	if err != nil{
		log.Println("Error creating new Directory "+projectRoot)
		log.Println(err)
		os.Exit(1003)
	}
	log.Println("Successfully created "+projectRoot)

	project.OutputDir = projectRoot

	err = os.Chdir(projectRoot)

	if err != nil{
		log.Println("Could not change directories to "+projectRoot)
		os.Exit(1002)
	}
}


