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
/// REPORT will announce thWe X,Y and F of the robot. This can be in any form, but standard output is sufficient.
///</remarks>
type ReportCommand struct {
    robot *models.Robot
}

func NewReportCommand(robot *models.Robot) *ReportCommand {
	return &ReportCommand{ robot }
} 

func (this *ReportCommand) Execute()  {
    if this.robot.IsValid == true {
        
        var direction = ""
        
        switch this.robot.Direction {
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
        
    	fmt.Printf("%s,%s,%s\n", strconv.Itoa(this.robot.X), strconv.Itoa(this.robot.Y), direction)
    }
}