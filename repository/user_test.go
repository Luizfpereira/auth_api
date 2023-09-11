package repository

import (
	"auth_api/entity"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRepoSuite struct {
	suite.Suite
	repo *UserRepositoryPSQL
}

func TestUserRepoSuite(t *testing.T) {
	suite.Run(t, new(UserRepoSuite))
}

func (s *UserRepoSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		s.Suite.T().Fatal(err)
	}
	db.AutoMigrate(entity.User{})

	repo := NewUserRepositoryPSQL(db)
	s.repo = repo

	u1, err := entity.NewUser("Luiz", "luiz.test@gmail.com", "*Mudar123")
	s.Assert().Nil(err)
	s.Assert().Nil(u1.Validate())

	user1, err := repo.CreateUser(u1)
	s.Assert().Nil(err)
	s.Assert().Equal(1, user1.ID)

	u2, err := entity.NewUser("Thais", "thais.montovani@gmail.com", "!Change456")
	s.Assert().Nil(err)
	s.Assert().Nil(u2.Validate())

	user2, err := repo.CreateUser(u2)
	s.Assert().Nil(err)
	s.Assert().Equal(2, user2.ID)
}

func (s *UserRepoSuite) TestCreateUser() {
	u3, err := entity.NewUser("Shirley", "shirley.test@gmail.com", "*Mudar123456")
	s.Assert().Nil(err)
	s.Assert().Nil(u3.Validate())
	user3, err := s.repo.CreateUser(u3)
	s.Assert().Nil(err)
	s.Assert().Equal(3, user3.ID)

	u4, err := entity.NewUser("Shirley", "shirley.test@gmail.com", "*Mudar123456")
	s.Assert().Nil(err)
	s.Assert().Nil(u4.Validate())
	user4, err := s.repo.CreateUser(u4)
	s.Assert().NotNil(err)
	s.Assert().Nil(user4)
}
