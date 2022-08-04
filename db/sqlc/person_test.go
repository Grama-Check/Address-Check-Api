package db

import (
	"context"
	"testing"

	"github.com/Grama-Check/Address-Check-Api/util"
	"github.com/stretchr/testify/require"
)

func TestGetPerson(t *testing.T) {

	person := createRandomPerson(t)

	person2, err := testQueries.GetPerson(context.Background(), person.Nic)

	require.NoError(t, err)
	require.NotEmpty(t, person2)

	require.Equal(t, person2.Address, person.Address)
	require.Equal(t, person2.Name, person.Name)
	require.Equal(t, person2.Nic, person.Nic)

}

func createRandomPerson(t *testing.T) Person {

	args := CreatePersonParams{
		Nic:     util.RandomID(),
		Address: util.RandomAddress(),
		Name:    util.RandomName(),
	}

	person, err := testQueries.CreatePerson(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, person)

	require.Equal(t, args.Address, person.Address)
	require.Equal(t, args.Name, person.Name)
	require.Equal(t, args.Nic, person.Nic)

	return person

}

func TestCreatePerson(t *testing.T) {
	createRandomPerson(t)
}
