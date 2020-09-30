package main

import (
	"fmt"
	"github.com/markbates/pkger"
	"io"
	"log"
	"os"
	"strconv"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Root Files																										 //
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func CreateRepoRootFiles(project *RobotProject){
	CreateGitignore(project)
	CreateGradleFiles(project)

	fmt.Println("Created Project Root Files")
}
func CreateGradleFiles(project *RobotProject){
	CopyGradleW()
	CopyGradleWBat()
	CreateGradleBuild(project)

}

func CopyGradleW(){
	type ps = struct {
		project string
	}
	parseSpec := ps{
		"FRC",
	}

	f, err := os.Create(COMPILE_ROOT+"/build/gradlew")
	if err != nil {
		log.Println("Error while creating gradlew: ", err)
		os.Exit(1004)
	}
	t, e := CompileTemplate("/templates/gradlew.tmpl")
	if e != nil {
		log.Println("Could not parse template gradlew.tmpl")
		log.Println(err)
		os.Exit(2001)
	}
	err = t.Execute(f, parseSpec)

	if err != nil {
		log.Println("Could not save gradlew to filesystem: ", err)
		os.Exit(1005)
	}
	log.Println("Created gradlew")
}

func CopyGradleWBat(){
	sourceFileStat, err := os.Stat(COMPILE_ROOT+"/templates/gradlew.bat")
	if err != nil {
		log.Println("Could not open gradlew.bat")
		log.Println(err)
		os.Exit(2001)
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Println("gradlew.tmpl.bat is not a regular file")
		os.Exit(2001)
	}

	source, err := pkger.Open(COMPILE_ROOT+"/templates/gradlew.bat")
	if err != nil {
		log.Println("Could not open gradlew.tmpl.bat")
		log.Println(err)
		os.Exit(2001)
	}
	defer source.Close()

	destination, err := os.Create(COMPILE_ROOT+"/build/gradlew.bat")
	if err != nil {
		log.Println("Could not open the destination")
		log.Println(err)
		os.Exit(2001)
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	log.Println("Copied " + strconv.FormatInt(nBytes, 10) + " bytes")
	if err != nil{
		log.Println("Could not save gradlew.tmpl.bat to filesystem: ", err)
		os.Exit(1005)
	}
	log.Println("Created gradlew.tmpl.bat")
}

func CreateGradleBuild(project *RobotProject){
	type BuildGradleParse struct{
		GradleRioV string
		LaunchPKG string
	}

	parseSpec := BuildGradleParse{
		GradleRioV: project.Metadata.GradleRioV,
		LaunchPKG:  project.Metadata.Package+".Main",
	}

	f, err := os.Create(COMPILE_ROOT+"/build/build.gradle")
	if err != nil {
		log.Println("Error while creating build.gradle: ", err)
		os.Exit(1004)
	}
	t, e := CompileTemplate("/templates/build-gradle.tmpl")
	if e != nil {
		log.Println("Could not parse template build-gradle.tmpl")
		log.Println(err)
		os.Exit(2001)
	}
	err = t.Execute(f, parseSpec)

	if err != nil {
		log.Println("Could not save build.gradle to filesystem: ", err)
		os.Exit(1005)
	}
	log.Println("Created build.gradle")

}

func CreateGitignore(project *RobotProject){
	parseSpec := FileParse{
		Package: project.Metadata.Package,
		Body:    nil,
	}

	f, err := os.Create(COMPILE_ROOT+"/build/.gitignore")
	if err != nil {
		log.Println("Error while creating .gitignore: ", err)
		os.Exit(1004)
	}
	t, e := CompileTemplate("/templates/gitignore.tmpl")
	if e != nil {
		log.Println("Could not parse template gitignore.tmpl")
		log.Println(err)
		os.Exit(2001)
	}
	err = t.Execute(f, parseSpec)

	if err != nil {
		log.Println("Could not save .gitignore to filesystem: ", err)
		os.Exit(1005)
	}
	log.Println("Created .gitignore")
}
