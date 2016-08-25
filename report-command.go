package main

import (
    "fmt"
)

func ReportCommand(robot *Robot) { 
    if robot.IsValid == true {
    	fmt.Printf("Test.")
    }
}