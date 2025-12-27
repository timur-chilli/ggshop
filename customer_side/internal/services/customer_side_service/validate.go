package customerSideService

import (
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"github.com/timur-chilli/ggshop/customer_side/internal/models"
)

func (s *CustomerSideService) validateInfo(GGOrdersInfos []*models.GGOrderInfo) error {
	for _, info := range GGOrdersInfos {
		if len(info.CustomerName) <= s.minNameLen || len(info.CustomerName) >= s.maxNameLen {
			return errors.New("имя не должно быть пустым и не должно превышать 100 символов")
		}
		if len(info.Details) <= 0 || len(info.Details) > 400 {
			return fmt.Errorf("неправильно указано описание заказа %v", info.Details)
		}
		if !s.isValidEmail(info.Email) {
			return fmt.Errorf("некорректный email у покупателя %v", info.Email)
		}
	}
	return nil
}

func (s *CustomerSideService) isValidEmail(email string) bool {
	if len(email) < 3 || len(email) > 254 {
		return false
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	if len(parts[1]) == 0 || len(parts[1]) > 253 {
		return false
	}

	return true
}
