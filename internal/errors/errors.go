package handlederror

var ErrTokenInvalid = UnauthorizedError("token is invalid")
var ErrInvalidCredentials = UnauthorizedError("credentials invalid").WithMessage("Wrong username or password")
var ErrEmptyRequestBody = ValidationError("empty request body").WithMessage("Please fill data")
var ErrInvalidRequestBody = ValidationError("invalid request body").WithMessage("Data is invalid")
