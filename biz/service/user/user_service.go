package service

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"

	"qnc/biz/dal/db"
	"qnc/biz/model/common"
	user "qnc/biz/model/user"
	"qnc/pkg/errno"
	"qnc/pkg/utils"
)

type UserService struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewUserService create user service
func NewUserService(ctx context.Context, c *app.RequestContext) *UserService {
	return &UserService{ctx: ctx, c: c}
}

// UserRegister register user return user id.
func (s *UserService) UserRegister(req *user.RegisterRequest) (user_id int64, err error) {
	user, err := db.QueryUser(req.Username)
	if err != nil {
		return 0, err
	}
	if *user != (db.User{}) {
		return 0, errno.UserAlreadyExistErr
	}

	passWord, err := utils.Crypt(req.Password)
	if err != nil {
		return 0, err
	}

	ts := time.Now().Unix()
	user_id, err = db.CreateUser(&db.User{
		UserName:   req.Username,
		Password:   passWord,
		Email:      req.Email,
		AvatarUrl:  "",
		Coin:       0.0,
		CreateTime: ts,
		UpdateTime: ts,
	})
	if err != nil {
		return 0, err
	}
	return user_id, nil
}

// UserInfo the function of user api
func (s *UserService) UserInfo(req *user.Request) (*common.User, error) {
	queryUserId := req.UserId
	currentUserId, exists := s.c.Get("current_user_id")
	if !exists {
		currentUserId = 0
	}
	return s.GetUserInfo(queryUserId, currentUserId.(int64))
}

// GetUserInfo
//
//	@Description: Query the information of query_user_id according to the current user user_id
//	@receiver *UserService
//	@param query_user_id int64
//	@param user_id int64  "Currently logged-in user id, may be 0"
//	@return *user.User
//	@return error
func (s *UserService) GetUserInfo(queryUserId, userId int64) (*common.User, error) {
	u := &common.User{}
	errChan := make(chan error, 7)
	defer close(errChan)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		dbUser, err := db.QueryUserById(queryUserId)
		if err != nil {
			errChan <- err
		} else {
			u.Name = dbUser.UserName
			u.Coin = dbUser.Coin
			// u.Avatar = utils.URLconvert(s.ctx, s.c, dbUser.Avatar)
			// u.BackgroundImage = utils.URLconvert(s.ctx, s.c, dbUser.BackgroundImage)
			// u.Signature = dbUser.Signature
		}
		wg.Done()
	}()

	// go func() {
	// 	WorkCount, err := db.GetWorkCount(query_user_id)
	// 	if err != nil {
	// 		errChan <- err
	// 	} else {
	// 		u.WorkCount = WorkCount
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	FollowCount, err := db.GetFollowCount(query_user_id)
	// 	if err != nil {
	// 		errChan <- err
	// 		return
	// 	} else {
	// 		u.FollowCount = FollowCount
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	FollowerCount, err := db.GetFollowerCount(query_user_id)
	// 	if err != nil {
	// 		errChan <- err
	// 	} else {
	// 		u.FollowerCount = FollowerCount
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	if user_id != 0 {
	// 		IsFollow, err := db.QueryFollowExist(user_id, query_user_id)
	// 		if err != nil {
	// 			errChan <- err
	// 		} else {
	// 			u.IsFollow = IsFollow
	// 		}
	// 	} else {
	// 		u.IsFollow = false
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	FavoriteCount, err := db.GetFavoriteCountByUserID(query_user_id)
	// 	if err != nil {
	// 		errChan <- err
	// 	} else {
	// 		u.FavoriteCount = FavoriteCount
	// 	}
	// 	wg.Done()
	// }()

	// go func() {
	// 	TotalFavorited, err := db.QueryTotalFavoritedByAuthorID(query_user_id)
	// 	if err != nil {
	// 		errChan <- err
	// 	} else {
	// 		u.TotalFavorited = TotalFavorited
	// 	}
	// 	wg.Done()
	// }()

	wg.Wait()
	select {
	case result := <-errChan:
		return &common.User{}, result
	default:
	}
	u.Id = queryUserId
	return u, nil
}
