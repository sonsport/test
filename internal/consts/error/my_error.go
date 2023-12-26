package constsError

import (
	"errors"
)

var InsufficientBalanceError = errors.New("insufficient balance")
var InsufficientPermissionsError = errors.New("InsufficientPermissionsError")
var AbnormalStateError = errors.New("AbnormalStateError")
var RoomLiveError = errors.New("RoomLiveError")
var RoomIsFullError = errors.New("RoomIsFullErr")

var RepeatBillingError = errors.New("Repeat Billing")

var UserBankNotSetError = errors.New("UserBankNotSetError")

var WithdrawalCheckError = errors.New("Withdrawal check exception")

var BankNotSupportError = errors.New("Bank not support")
