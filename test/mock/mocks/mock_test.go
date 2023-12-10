package mocks

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_Eat(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPerson := NewMockPerson(ctrl)

	mockPerson.EXPECT().Eat("Apple").Times(1)
}

func Test_Sleep(t *testing.T) {

}
