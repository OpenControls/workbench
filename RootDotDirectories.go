package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"text/template"
)

func CreateDotDirectories(project *RobotProject){
	InitializeDotDirs()
	CopyLaunch()
	CopySettings()
	CreateWPILibPrefs(project)
}

func InitializeDotDirs(){
	vscodeDir := COMPILE_ROOT+"/build/.vscode"
	err := os.MkdirAll(vscodeDir, os.ModePerm)
	if err != nil{
		log.Println("Error creating new Directory "+vscodeDir)
		log.Println(err)
		os.Exit(1003)
	}
	log.Println("Successfully created "+vscodeDir)

	wpilibDir := COMPILE_ROOT+"/build/.wpilib"
	err = os.MkdirAll(wpilibDir, os.ModePerm)
	if err != nil{
		log.Println("Error creating new Directory "+wpilibDir)
		log.Println(err)
		os.Exit(1003)
	}
	log.Println("Successfully created "+wpilibDir)


}

func CopyLaunch(){
	fileName := "launch.json"

	fromFile := COMPILE_ROOT + "/templates/dotfiles/launch.json"
	toFile := COMPILE_ROOT + "/build/.vscode/"+fileName

	sourceFileStat, err := os.Stat(fromFile)
	if err != nil {
		log.Println("Could not open "+fileName)
		log.Println(err)
		os.Exit(2001)
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Println(fileName+" is not a regular file")
		os.Exit(2001)
	}

	source, err := os.Open(fromFile)
	if err != nil {
		log.Println("Could not open "+fileName)
		log.Println(err)
		os.Exit(2001)
	}
	defer source.Close()

	destination, err := os.Create(toFile)
	if err != nil {
		log.Println("Could not open the destination")
		log.Println(err)
		os.Exit(2001)
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	log.Println("Copied " + strconv.FormatInt(nBytes, 10) + " bytes")
	if err != nil{
		log.Println("Could not save "+fileName+ " to filesystem: ", err)
		os.Exit(1005)
	}
	log.Println("Created "+fileName)
}

func CopySettings(){
	fileName := "settings.json"

	fromFile := COMPILE_ROOT + "/templates/dotfiles/settings.json"
	toFile := COMPILE_ROOT + "/build/.vscode/"+fileName


	sourceFileStat, err := os.Stat(fromFile)
	if err != nil {
		log.Println("Could not open "+fileName)
		log.Println(err)
		os.Exit(2001)
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Println(fileName+" is not a regular file")
		os.Exit(2001)
	}

	source, err := os.Open(fromFile)
	if err != nil {
		log.Println("Could not open "+fileName)
		log.Println(err)
		os.Exit(2001)
	}
	defer source.Close()

	destination, err := os.Create(toFile)
	if err != nil {
		log.Println("Could not open the destination")
		log.Println(err)
		os.Exit(2001)
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	log.Println("Copied " + strconv.FormatInt(nBytes, 10) + " bytes")
	if err != nil{
		log.Println("Could not save "+fileName+ " to filesystem: ", err)
		os.Exit(1005)
	}
	log.Println("Created "+fileName)
}

func CreateWPILibPrefs(project *RobotProject){
	fileName := "wpilib_preferences.json"
	inputFile := COMPILE_ROOT + "/templates/dotfiles/wpilib_preferences.tmpl"
	outputFile := COMPILE_ROOT+"/build/.wpilib/"+fileName
	type LocalParse struct{
		Year int
		Team int
	}
	parseSpec := LocalParse{
		Year: project.Metadata.Year,
		Team: project.Metadata.Team,
	}

	f, err := os.Create(outputFile)
	if err != nil {
		log.Println("Error while creating "+fileName + " : ", err)
		os.Exit(1004)
	}
	t, e := template.ParseFiles(inputFile)
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

