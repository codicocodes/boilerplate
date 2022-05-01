package userservice

type BadUserInput struct {
    msg string
}

func (e BadUserInput) Error() string {
    return e.msg
}

func newBadUserInputError(msg string) BadUserInput  {
    return BadUserInput{msg: msg}
}

type FailedUserCreation struct {
    msg string
}

func (e FailedUserCreation) Error() string {
    return e.msg
}

func newFailedUserCreationError(msg string) FailedUserCreation  {
    return FailedUserCreation{msg: msg}
}
