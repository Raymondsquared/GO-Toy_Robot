package commands

import (
    "consts"
    "models"
)

///<summary>
/// Concrete implementation of ICommand interface for Command Pattern
///</summary>
///<remarks>
/// Handling all place commmands, send the command to the receiver to do some actions
/// PLACE will put the toy robot on the table in position X,Y and facing NORTH, SOUTH, EAST or WEST.
/// The origin (0,0) can be considered to be the SOUTH WEST most corner.
///</remarks>
func PlaceCommand(robot *models.Robot, tableTop *models.TableTop, x int, y int, direction consts.DIRECTION) { 
    
    var isValid = true;
    
    if tableTop == nil {
        isValid = false;
    }
    
    if x < 0 || y < 0 {
        isValid = false;
    }
    
    if x > tableTop.Width || y > tableTop.Length {
        isValid = false;
    }
    
    if isValid == true {
        robot.IsValid = isValid
        robot.TableTop = *tableTop
        robot.X = x
        robot.Y = y
        robot.Direction = direction
    }

}