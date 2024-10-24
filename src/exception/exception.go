package exception

import (
	"go-pix-api/src/models"
)

func InvalidLoginError() *models.AppError {
	return &models.AppError{
		Code:    403,
		Message: "MSG_INVALID_LOGIN",
	}
}

func UserExistError() *models.AppError {
	return &models.AppError{
		Code:    409,
		Message: "MSG_USER_EXISTS",
	}
}

func PasswordNotSecureError() *models.AppError {
	return &models.AppError{
		Code:    409,
		Message: "MSG_PASSWORD_NOT_SECURE",
	}
}

func MissingFieldError() *models.AppError {
	return &models.AppError{
		Code:    400,
		Message: "MSG_MISSING_FIELD",
	}
}

func WrongDataTypeError() *models.AppError {
	return &models.AppError{
		Code:    400,
		Message: "MSG_WROND_DATA_TYPE",
	}
}

func ImageCanNotProcessError() *models.AppError {
	return &models.AppError{
		Code:    400,
		Message: "MSG_IMAGE_CAN_NOT_PROCESS",
	}
}

func InternalServerError() *models.AppError {
	return &models.AppError{
		Code:    400,
		Message: "MSG_INTERNAL_SERVER_ERROR",
	}
}

func InvalidAccessTokenError() *models.AppError {
	return &models.AppError{
		Code:    401,
		Message: "MSG_INVALID_ACCESS_TOKEN",
	}
}

func PostNotExistsError() *models.AppError {
	return &models.AppError{
		Code:    404,
		Message: "MSG_POST_NOT_EXISTS",
	}
}

func CommentNotExistsError() *models.AppError {
	return &models.AppError{
		Code:    404,
		Message: "MSG_COMMENT_NOT_EXISTS",
	}
}

func UserNotExistsError() *models.AppError {
	return &models.AppError{
		Code:    404,
		Message: "MSG_USER_NOT_EXISTS",
	}
}

func PermissionDenyError() *models.AppError {
	return &models.AppError{
		Code:    403,
		Message: "MSG_PERMISSION_DENY",
	}
}
