package token

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/emadghaffari/api-teacher/database/redis"
	model "github.com/emadghaffari/api-teacher/model/user"
	"github.com/emadghaffari/api-teacher/utils/hash"
	"github.com/spf13/viper"
)

var (
	// Conf variable instance of intef
	Conf intef = &wt{}
)

type intef interface {
	Generate(user model.User) (*TokenDetails, error)
	VerifyToken(string) (*AccessDetails, error)
}
type wt struct{}

func (j *wt) Generate(user model.User) (*TokenDetails, error) {

	td, err := j.genJWT()
	if err != nil {
		return nil, err
	}

	if err := j.genRefJWT(td); err != nil {
		return nil, err
	}

	if err := j.redis(user, td); err != nil {
		return nil, err
	}

	return td, nil
}

func (j *wt) VerifyToken(request string) (*AccessDetails, error) {
	token, err := jwt.Parse(request, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		AccessUUID, ok := claims["uuid"].(string)
		if !ok {
			return nil, fmt.Errorf("Error in claims uuid from client")
		}

		return &AccessDetails{
			AccessUUID: AccessUUID,
		}, nil
	}
	return nil, err
}

func (j *wt) genJWT() (*TokenDetails, error) {
	secret := viper.GetString("jwt.secret")

	// create new TokenDetails
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Duration(time.Minute * viper.GetDuration("jwt.expire"))).Unix()
	td.RtExpires = time.Now().Add(time.Duration(time.Minute * viper.GetDuration("jwt.RTexpire"))).Unix()
	td.AccessUUID = hash.Generate(30)
	td.RefreshUUID = hash.Generate(60)

	// New MapClaims for access token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["uuid"] = td.AccessUUID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	var err error
	td.AccessToken, err = at.SignedString([]byte(secret))
	if err != nil {
		log.WithFields(log.Fields{
			"uuid": td.AccessUUID,
		}).Warn(fmt.Sprintf("Error in Generate JWT: %s", err))
		return nil, err
	}
	return td, nil
}

func (j *wt) genRefJWT(td *TokenDetails) error {
	secret := viper.GetString("jwt.refSecret")

	// New MapClaims for refresh access token
	rtClaims := jwt.MapClaims{}
	rtClaims["uuid"] = td.RefreshUUID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	var err error
	td.RefreshToken, err = rt.SignedString([]byte(secret))
	if err != nil {
		log.WithFields(log.Fields{
			"ref_uuid": td.RefreshUUID,
		}).Warn(fmt.Sprintf("Error in Generate RefJWT: %s", err))
		return err
	}
	return nil
}

func (j *wt) redis(user model.User, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	// make map for store in redis
	us := make(map[string]string, 4)
	us["identitiy"] = user.Identitiy
	us["first_name"] = user.FirstName
	us["last_name"] = user.LastName
	us["role"] = user.Role.Name

	bt, err := json.Marshal(us)
	if err != nil {
		log.WithFields(log.Fields{
			"user": user,
		}).Warn(fmt.Sprintf("Error in Marshal User: %s", err))
		return err
	}
	if err := redis.DB.GetDB().Set(context.Background(), td.AccessUUID, string(bt), at.Sub(now)).Err(); err != nil {
		log.WithFields(log.Fields{
			"user_id":      user.ID,
			"access_token": td.AccessToken,
		}).Warn(fmt.Sprintf("Error in Store JWT in Redis: %s", err))
		return err
	}

	if err := redis.DB.GetDB().Set(context.Background(), td.RefreshUUID, string(bt), rt.Sub(now)).Err(); err != nil {
		log.WithFields(log.Fields{
			"user_id":      user.ID,
			"access_token": td.AccessToken,
			"ref_token":    td.RefreshToken,
		}).Warn(fmt.Sprintf("Error in Store RefJWT in Redis: %s", err))
		return err
	}
	return nil
}
