package models

import (
    "consts"
)

///<summary>
/// Concrete implementation of Receiver for Command Pattern
///</summary>
type Input struct {
  Command consts.COMMAND
  TableTop *TableTop
  X, Y int
  Direction consts.DIRECTION
}