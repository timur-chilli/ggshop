package customerSideService

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/timur-chilli/ggshop/customer_side/internal/models"
	"github.com/timur-chilli/ggshop/customer_side/internal/services/customer_side_service/mocks"
	"gotest.tools/v3/assert"
)

type CustomerSideServiceSuite struct {
	suite.Suite
	ctx            context.Context
	mockStorage *mocks.MockGGOrderStorage
	customerSideService *CustomerSideService 
}

func (s *CustomerSideServiceSuite) SetupTest() {
	s.mockStorage = mocks.NewMockGGOrderStorage(s.T())
	s.ctx = context.Background()
	s.customerSideService = NewCustomerSideService(s.ctx, s.mockStorage, 0, 100)
}


func (s *CustomerSideServiceSuite) TestInsertSuccess() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 1,
			CustomerName: "Timur",
			Email: "timur@timur.ru",
			Details: "Lorem ipsum",
		},
	}
	s.mockStorage.EXPECT().InsertGGOrderInfo(s.ctx, ggorders).Return(nil)
	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)
	assert.NilError(s.T(), err)
}


func (s *CustomerSideServiceSuite) TestInsertSuccess2() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 3,
			CustomerName: "Tim",
			Email: "google@gmail.ru",
			Details: "..............",
		},
		{
			ID: 2,
			CustomerName: "Pupupu",
			Email: "zuzuzu@zu.zu",
			Details: ",,,,,,,,,,,,,",
		},
	}
	s.mockStorage.EXPECT().InsertGGOrderInfo(s.ctx, ggorders).Return(nil)
	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)
	assert.NilError(s.T(), err)
}

func (s *CustomerSideServiceSuite) TestInsertStorageError() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 1,
			CustomerName: "Timur",
			Email: "timur@timur.ru",
			Details: "Lorem ipsum",
		},
	}
	wantErr := errors.New("error")
	s.mockStorage.EXPECT().InsertGGOrderInfo(s.ctx, ggorders).Return(wantErr)
	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)
	assert.ErrorIs(s.T(), err, wantErr)
}

func (s *CustomerSideServiceSuite) TestInsertEmptyNameError() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 23231,
			CustomerName: "",
			Email: "timur@timur.ru",
			Details: "Lorem ipsum",
		},
	}

	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)

	assert.Check(s.T(), err != nil)

}

func (s *CustomerSideServiceSuite) TestInsertCustomerNameTooLarge() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 1121211,
			CustomerName: strings.Repeat("x", 101),
			Email: "timur@timur.ru",
			Details: "Lorem ipsum",
		},
	}

	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)

	assert.Check(s.T(), err != nil)

}

func (s *CustomerSideServiceSuite) TestInsertDetailsTooLarge() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 99,
			CustomerName: "iamlargedetailscustomer",
			Email: "just@email.com",
			Details: strings.Repeat("x", 401),
		},
	}

	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)

	assert.Check(s.T(), err != nil)

}

func (s *CustomerSideServiceSuite) TestInsertDetailsEmpty() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 55,
			CustomerName: "iamemptydetailscustomer",
			Email: "just@email.com",
			Details: "",
		},
	}

	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)

	assert.Check(s.T(), err != nil)

}

func (s *CustomerSideServiceSuite) TestWrongEmail1() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 55,
			CustomerName: "iamemptydetailscustomer",
			Email: "@@@@@",
			Details: "b",
		},
	}

	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)

	assert.Check(s.T(), err != nil)

}

func (s *CustomerSideServiceSuite) TestEmptyEmail() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 56,
			CustomerName: "iamemptyemailscustomer",
			Email: "",
			Details: "a",
		},
	}

	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)

	assert.Check(s.T(), err != nil)

}


func (s *CustomerSideServiceSuite) TestTooBigEmail() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 57,
			CustomerName: "iamtoobigemailscustomer",
			Email: strings.Repeat("x", 254) + "@mail.ru",
			Details: "a",
		},
	}

	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)

	assert.Check(s.T(), err != nil)

}


func (s *CustomerSideServiceSuite) TestTooSmallEmail() {
	ggorders := []*models.GGOrderInfo{
		{
			ID: 58,
			CustomerName: "iamtoosmallemailscustomer",
			Email: "@m",
			Details: "a",
		},
	}

	err := s.customerSideService.InsertGGOrderInfo(s.ctx, ggorders)

	assert.Check(s.T(), err != nil)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(CustomerSideServiceSuite))
}