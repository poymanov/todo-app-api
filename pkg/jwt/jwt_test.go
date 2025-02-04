package jwt_test

import (
	"github.com/stretchr/testify/require"
	"poymanov/todo/pkg/jwt"
	"testing"
)

const secret = "test"
const expectedEmail = "test@test.ru"
const expectedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAdGVzdC5ydSJ9.fRLO5IdkhBlqOK9RreHzl-bikPgYeyBIqHkHOPxPtro"

func TestNewJWTSuccess(t *testing.T) {
	jwtLib := jwt.NewJWT(secret)

	require.NotNil(t, jwtLib)
	require.Equal(t, secret, jwtLib.Secret)
}

func TestCreateSuccess(t *testing.T) {
	jwtLib := jwt.NewJWT(secret)

	token, err := jwtLib.Create(jwt.JWTData{
		Email: expectedEmail,
	})

	require.NotNil(t, token)
	require.NoError(t, err)

	require.Equal(t, expectedToken, token)
}

func TestParseSuccess(t *testing.T) {
	jwtLib := jwt.NewJWT(secret)

	isSuccess, jwtData := jwtLib.Parse(expectedToken)

	require.True(t, isSuccess)
	require.NotNil(t, jwtData)
	require.Equal(t, expectedEmail, jwtData.Email)
}

func TestParseFailed(t *testing.T) {
	jwtLib := jwt.NewJWT("test2")

	isSuccess, jwtData := jwtLib.Parse(expectedToken)

	require.False(t, isSuccess)
	require.Nil(t, jwtData)
}
