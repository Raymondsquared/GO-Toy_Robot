package commands

import (
    "consts"
    "models"
)

///<summary>
/// Concrete implementation of ICommand interface for Command Pattern
///</summary>
///<remarks>
/// Handling all move commmands, send the command to the receiver to do some actions
/// MOVE will move the toy robot one unit forward in the direction it is currently facing.
///</remarks>
type MoveCommand struct {
    robot *models.Robot
}

func NewMoveCommand(robot *models.Robot) *MoveCommand {
	return &MoveCommand{ robot }
} 

func (this *MoveCommand) Execute()  {
    if this.robot.IsValid == true {
        switch this.robot.Direction {
            case consts.NORTH:
                if this.robot.Y < this.robot.TableTop.Length {
                    this.robot.Y++
                }
            case consts.EAST:
                if this.robot.X < this.robot.TableTop.Width {
                    this.robot.X++
                }
            case consts.SOUTH:
                if this.robot.Y > 0 {
                    this.robot.Y--
                }
            case consts.WEST:
                if this.robot.X > 0 {
                    this.robot.X--
                }
        }
    }
}