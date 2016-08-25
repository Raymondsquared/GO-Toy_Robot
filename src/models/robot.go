package models

import (
    "consts"
)

///<summary>
/// Concrete implementation of Receiver for Command Pattern
///</summary>
type Robot struct {
  TableTop
  X, Y int
  Direction consts.DIRECTION
  IsValid bool
}