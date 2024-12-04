package domain

type UserInfo struct {
	UserId       string
	Avatar       string
	Birthday     string
	Email        string
	Name         string
	AccessToken  string
	RefreshToken string
}

func (u *UserInfo) SetUserId(userId string) *UserInfo {
	u.UserId = userId
	return u
}

func (u *UserInfo) SetAvatar(avatar string) *UserInfo {
	u.Avatar = avatar
	return u
}

func (u *UserInfo) SetBirthday(birthday string) *UserInfo {
	u.Birthday = birthday
	return u
}

func (u *UserInfo) SetEmail(email string) *UserInfo {
	u.Email = email
	return u
}

func (u *UserInfo) SetName(name string) *UserInfo {
	u.Name = name
	return u
}

func (u *UserInfo) SetAccessToken(accessToken string) *UserInfo {
	u.AccessToken = accessToken
	return u
}

func (u *UserInfo) SetRefreshToken(refreshToken string) *UserInfo {
	u.RefreshToken = refreshToken
	return u
}
