package models

import (
  "time"
  "errors"
  "strings"
  "html"

  "github.com/badoux/checkmail"
  "golang.org/x/crypto/bcrypt"
  "github.com/jinzhu/gorm"
)



type User struct {
  ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
  Name      string `gorm:"size:255;not null" json:"name"`
  Username  string `gorm:"size:255;not null;unique" json:"username"`
  Email     string `gorm:"size:255;not null;unique" json:"email"`
  Phone     string `gorm:"size:255" json:"phone"`
  ClientID  uint64 `gorm:"not null" json:"client_id"`
  EmailVerifiedAt time.Time `json:"email_verified_at"`
  Password  string `gorm:"size:255;not null;" json:"password"`
  Timezone  string `gorm:"size:255" json:"timezone"`
  RememberToken string `gorm:"size:100" json:"remember_token"`
  CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
  UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) GetUser(db *gorm.DB, usr uint64) (*User, error)  {
		var err error
		err = db.Debug().Model(&User{}).Where("id = ?", usr).Take(&u).Error
		if err != nil {
			return &User{}, err
		}
		return u, nil
}

func (u *User) GetUserName(db *gorm.DB, usr string) (*User, error)  {
		var err error
		err = db.Debug().Model(&User{}).Where("username = ?", usr).Take(&u).Error
		if err != nil {
			return &User{}, err
		}
		return u, nil
}

func Hash(password string) ([]byte, error) {
  return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error  {
  return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func (u *User) BeforeSave() error  {
  hashedPassword, err := Hash(u.Password)
  if err != nil {
    return err
  }
  u.Password = string(hashedPassword)
  return nil
}

func (u *User) Prepare()  {
  u.ID = 0
  u.Name = html.EscapeString(strings.TrimSpace(u.Name))
  u.Email = html.EscapeString(strings.TrimSpace(u.Email))
  u.CreatedAt = time.Now()
  u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error  {
  switch strings.ToLower(action) {
  case "update":
    if u.Name == "" {
      return errors.New("Required Nickname")
    }
    if u.Password == "" {
      return errors.New("Required Password")
    }
    if u.Email == "" {
      return errors.New("Required Email")
    }
    if err:= checkmail.ValidateFormat(u.Email); err != nil {
      return errors.New("Invalid Email")
    }

    return nil

  case "login":
    if u.Password == "" {
      return errors.New("Required Password")
    }
    if u.Email == "" {
      return errors.New("Invalid Email")
    }
    return nil

  default:
    if u.Name == "" {
      return errors.New("Required Nickname")
    }
    if u.Password == "" {
      return errors.New("Required Password")
    }
    if u.Email == "" {
      return errors.New("Required Email")
    }
    if err := checkmail.ValidateFormat(u.Email); err != nil {
      return errors.New("Invalid Email")
    }
    return nil

  }
}

