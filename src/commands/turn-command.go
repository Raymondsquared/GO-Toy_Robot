package commands

import (
    "consts"
    "models"
)

///<summary>
/// Concrete implementation of ICommand interface for Command Pattern
///</summary>
///<remarks>
/// Handling all turning commmands (left / right), send the command to the receiver to do some actions
/// LEFT and RIGHT will rotate the robot 90 degrees in the specified direction without changing the position of the robot.
///</remarks>
type TurnCommand struct {
    robot *models.Robot
    turn consts.TURN
}

func NewTurnCommand(robot *models.Robot, turn consts.TURN) *TurnCommand {
	return &TurnCommand{ robot, turn }
} 

func (this *TurnCommand) Execute()  {
    if this.robot.IsValid == true {
        switch this.robot.Direction {
            case consts.NORTH:
                if this.turn == consts.LEFT {
                    this.robot.Direction = consts.WEST
                }
                if this.turn == consts.RIGHT {
                    this.robot.Direction = consts.EAST
                }
            case consts.EAST:
                if this.turn == consts.LEFT {
                    this.robot.Direction = consts.NORTH
                }
                if this.turn == consts.RIGHT {
                    this.robot.Direction = consts.SOUTH
                }
            case consts.SOUTH:
                if this.turn == consts.LEFT {
                    this.robot.Direction = consts.EAST
                }
                if this.turn == consts.RIGHT {
                    this.robot.Direction = consts.WEST
                }
            case consts.WEST:
                if this.turn == consts.LEFT {
                    this.robot.Direction = consts.SOUTH
                }
                if this.turn == consts.RIGHT {
                    this.robot.Direction = consts.NORTH
                }
        }
    }
}