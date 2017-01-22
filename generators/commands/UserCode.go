package commands
type UserCodeGet struct {
	node byte
	UserIdentifier byte
}

func NewUserCodeGet() UserCodeGet {
	return UserCodeGet{}
}

func (c *UserCodeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *UserCodeGet) Set(UserIdentifier byte,) error {
	c.UserIdentifier = UserIdentifier

	return nil
}

func (c *UserCodeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(UserCode),
		0x02,
		c.UserIdentifier,
		0x25,
	}
}
type UserCodeReport struct {
    *report
    UserIdentifier byte
    UserIDStatus byte
    USER_CODE byte
    data []byte
}

func NewUserCodeReport(data []byte) *UserCodeReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &UserCodeReport{
				UserIdentifier: data[0],
				UserIDStatus: data[1],
				USER_CODE: data[2],
        data: data,
    }
}

type UserCodeSet struct {
	node byte
	UserIdentifier byte
	UserIDStatus byte
	USER_CODE byte
}

func NewUserCodeSet() UserCodeSet {
	return UserCodeSet{}
}

func (c *UserCodeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *UserCodeSet) Set(UserIdentifier byte,UserIDStatus byte,USER_CODE byte,) error {
	c.UserIdentifier = UserIdentifier
	c.UserIDStatus = UserIDStatus
	c.USER_CODE = USER_CODE

	return nil
}

func (c *UserCodeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(UserCode),
		0x01,
		c.UserIdentifier,
		c.UserIDStatus,
		c.USER_CODE,
		0x25,
	}
}
type UsersNumberGet struct {
	node byte
}

func NewUsersNumberGet() UsersNumberGet {
	return UsersNumberGet{}
}

func (c *UsersNumberGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *UsersNumberGet) Set() error {

	return nil
}

func (c *UsersNumberGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(UserCode),
		0x04,
		0x25,
	}
}
type UsersNumberReport struct {
    *report
    SupportedUsers byte
    data []byte
}

func NewUsersNumberReport(data []byte) *UsersNumberReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &UsersNumberReport{
				SupportedUsers: data[0],
        data: data,
    }
}

