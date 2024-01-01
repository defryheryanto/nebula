package handlederror

var ErrTokenInvalid = UnauthorizedError("token is invalid")
var ErrInvalidCredentials = UnauthorizedError("credentials invalid").WithMessage("Wrong username or password")
