package commands

import (
    "consts"
    "models"
    "fmt"
    "strconv"
)

///<summary>
/// Concrete implementation of ICommand interface for Command Pattern
///</summary>
///<remarks>
/// Handling all reporting commmands, send the command to the receiver to do some actions
/// REPORT will announce the X,Y and F of the robot. This can be in any form, but standard output is sufficient.
///</remarks>
func ReportCommand(robot *models.Robot) { 
    if robot.IsValid == true {
        
        var direction = "";
        
        switch robot.Direction {
            case consts.NORTH:
                direction = "NORTH"
            case consts.EAST:
                direction = "EAST"
            case consts.SOUTH:
                direction = "SOUTH"
            case consts.WEST:
                direction = "WEST"
            default:
                direction = "UNKNOWN_DIRECTION"
        }
        
    	fmt.Printf("%s,%s,%s\n", strconv.Itoa(robot.X), strconv.Itoa(robot.Y), direction)
    }
}