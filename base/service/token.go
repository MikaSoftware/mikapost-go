package service

import (
    "context"
    "time"
    "github.com/dgrijalva/jwt-go"
    "github.com/go-chi/jwtauth"

    "github.com/mikasoftware/mikapost-go/base/config"
)

// Our variables.
var tokenAuth *jwtauth.JWTAuth
var mySigningKey []byte

// Function will either lazily create and return the JWT token authority object,
// or return the JWT token authority if it was previously created.
func GetJWTTokenAuthority() *jwtauth.JWTAuth {
    // Get our applications secret signing key by lazily loading it if
    // we haven't already lazily loaded it yet.
    if mySigningKey == nil {
        secretString := config.GetSettingsVariableSigningSecretKey()
        mySigningKey = []byte(secretString)
    }

    if tokenAuth != nil {
        return tokenAuth
    }

    // Generate our token signing authority.
	tokenAuth = jwtauth.New("HS256", mySigningKey, nil)

    // Return our toke authority.
    return tokenAuth
}

// Function will generate a unique JWT token with the claims data of holding
// the `userID` paramter.
func GenerateJWTToken(userID uint64) string {
    // Get our JWT token authority.
    myTokenAuth := GetJWTTokenAuthority()

    // Create our claims data for the JWT token.
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    }

    // Generate our new JWT token.
    _, tokenString, _ := myTokenAuth.Encode(claims)

    // Return our new JWT token.
    return tokenString
}


// Function used to create a new context with the JWT authenticated `claims`
// already attached to the context. This function should be used in unit tests.
func NewContextWithJWTToken(token string) context.Context {
    // SPECIAL THANKS:
    // https://github.com/go-chi/jwtauth/blob/master/jwtauth.go

    // STEP 1: Get JWT Authority.
    ja := GetJWTTokenAuthority()

    // STEP 2: Generate our token
    token_obj, err := ja.Decode(token)

    // STEP 3: Create our context were the user is logged and attach it to
    //         our request.
    ctx := context.Background()
    ctx = jwtauth.NewContext(ctx, token_obj, err)
    return ctx
}


// Function used to extract the `user_id` from the JWT token that was passed in
// by the JWT middleware when a successful authentication happened.
func GetUserIDFromContext(ctx context.Context) uint64 {
    // Fetch the claims based on what the JWT token was encoded and
    // encrypted with. We will extrac the user ID value and look it up.
    _, claims, _ := jwtauth.FromContext(ctx)
	raw_user_id := claims["user_id"]
	if raw_user_id == nil {
		return 0
	}

    var user_id float64 = claims["user_id"].(float64) // Set to default format.
    var id uint64 = uint64(user_id)
    return id
}
