package postgresql

var returningIDSuffix = "RETURNING id"

var authTable = "public.auth"

// public.auth fields
var (
	authIdField    = "public.auth.id"
	authLoginField = "public.auth.login"
	authHashField  = "public.auth.password_hash"
)

var (
	authLoginSelectFields  = []string{authIdField, authLoginField, authHashField}
	authSignupInsertFields = []string{authLoginField, authHashField}
)

var userTable = "public.user"

// public.user fields
var (
	userIdField      = "public.user.id_user"
	userNameField    = "public.user.name"
	userSurnameField = "public.user.surname"
	userEmailField   = "public.user.email"
	userDOBField     = "public.user.dob"
)

var (
	newUserInsertFields  = []string{userIdField, userEmailField, userDOBField}
	userFullSelectFields = []string{userIdField, userNameField, userSurnameField, userEmailField, userDOBField}
)
