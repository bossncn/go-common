package error_code

const UnknownError string = "ERR00000"
const InternalServerError string = "ERR00001"
const ValidationError string = "ERR00002"
const Unauthorized string = "ERR00003"
const Forbidden string = "ERR00004"
const NotFound string = "ERR00005"

func GetHTTPStatusCode(code string) int {
	switch code {
	case ValidationError:
		return 400
	case Unauthorized:

		return 401
	case Forbidden:

		return 403
	case NotFound:
		return 404
	case InternalServerError:
		return 500
	default:
		return 500
	}
}

func GetErrorCode(code string) string {
	switch code {
	case ValidationError:
		return ValidationError
	case Unauthorized:
		return Unauthorized
	case Forbidden:
		return Forbidden
	case InternalServerError:
		return InternalServerError
	default:
		return UnknownError
	}
}

func GetErrorMessage(code string) string {
	switch code {
	case ValidationError:
		return "ValidationError"
	case Unauthorized:

		return "Unauthorized"
	case Forbidden:

		return "Forbidden"
	case InternalServerError:
		return "InternalServerError"
	default:
		return code
	}
}
