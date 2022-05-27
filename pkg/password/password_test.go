package password

import (
	"fmt"
	"testing"

	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	//testHashPassword()
	password := utils.RandomString(10)

	hashPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotZero(t, hashPassword)
	require.NoError(t, CheckPassword(password, hashPassword))
	wrongPassword := utils.RandomString(10)
	require.Error(t, CheckPassword(wrongPassword, hashPassword))
}

func testHashPassword() {
	password := "123456"
	hashPassword, _ := HashPassword(password)
	fmt.Println(hashPassword)
}
