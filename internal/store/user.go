package store

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt int64
}

func (s *Store) UserCount() (int64, error) {
	var count int64
	if err := s.DB.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
