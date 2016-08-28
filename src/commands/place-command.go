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
type PlaceCommand struct {
    robot *models.Robot
    
    tableTop *models.TableTop
    x int
    y int
    direction consts.DIRECTION
}

func NewPlaceCommand(robot *models.Robot, tableTop *models.TableTop, x int, y int, direction consts.DIRECTION) (*PlaceCommand) {
	return &PlaceCommand{ robot, tableTop, x, y, direction }
} 

func (this *PlaceCommand) Execute()  {
    var isValid = true;
    
    if this.tableTop == nil {
        isValid = false;
    }
    
    if this.x < 0 || this.y < 0 {
        isValid = false;
    }
    
    if this.x > this.tableTop.Width || this.y > this.tableTop.Length {
        isValid = false;
    }
    
    if isValid == true {
        this.robot.IsValid = isValid
        this.robot.TableTop = *this.tableTop
        this.robot.X = this.x
        this.robot.Y = this.y
        this.robot.Direction = this.direction
    }

}