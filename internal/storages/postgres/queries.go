package postgresql

var returningIDSuffix = "RETURNING id"

var authTable = "public.auth"

// public.auth fields
var (
	publicAuthIdField       = "public.auth.id"
	publicAuthUsernameField = "public.auth.username"
	publicAuthHashField     = "public.auth.password_hash"
)

// internal public.auth fields
var (
	authIdField       = "id"
	authUsernameField = "username"
	authHashField     = "password_hash"
)

var (
	authLoginSelectFields  = []string{authIdField, authUsernameField, authHashField}
	authSignupInsertFields = []string{authUsernameField, authHashField}
)

var userTable = "public.user"

// public.user fields
var (
	publicUserIdField      = "public.user.id_user"
	publicUserNameField    = "public.user.name"
	publicUserSurnameField = "public.user.surname"
	publicUserEmailField   = "public.user.email"
	publicUserDOBField     = "public.user.dob"
)

// internal public.user fields
var (
	userIdField      = "id_user"
	userNameField    = "name"
	userSurnameField = "surname"
	userEmailField   = "email"
	userDOBField     = "dob"
)

var (
	newUserInsertFields  = []string{userIdField, userEmailField, userDOBField}
	userFullSelectFields = []string{userIdField, userNameField, userSurnameField, userEmailField, userDOBField}
)
