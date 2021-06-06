package service

import (
	"fmt"
)

func (s *service) Delete(fieldID string) error {
	err := s.repository.DeleteOne(fieldID)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
