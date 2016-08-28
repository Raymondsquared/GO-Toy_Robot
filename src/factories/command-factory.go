package factories

import (
    "models"
    "consts"
    "commands"
)

func CommandFactory(robot *models.Robot, input models.Input) commands.Command{
    switch (input.Command) {
        case consts.PLACE_COMMAND:
            return commands.NewPlaceCommand(robot, input.TableTop, input.X, input.Y, input.Direction)
        case consts.MOVE_COMMAND:
            return commands.NewMoveCommand(robot)
        case consts.LEFT_COMMAND:
            return commands.NewTurnCommand(robot, consts.LEFT)
        case consts.RIGHT_COMMAND:
            return commands.NewTurnCommand(robot, consts.RIGHT)
        case consts.REPORT_COMMAND:
            return commands.NewReportCommand(robot)
        default:
            return nil
    }
}