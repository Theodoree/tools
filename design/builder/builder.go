package builder

// builder
type User struct {
    NickName string
    Age      uint64
}

func createUser(b *builder) (u User) {
    u.Age = b.age
    u.NickName = b.nickName
    return
}

type builder struct {
    nickName string
    age      uint64
}

func Builder() *builder {
    return &builder{}
}

func (b *builder) SetNickName(name string) *builder {
    b.nickName = name
    return b
}

func (b *builder) SetAge(age uint64) *builder {
    b.age = age
    return b
}

func (b *builder) build() (User, error) {
    if len(b.nickName) == 0 {
        // do something
    }

    if b.age < 0 {
        // do something
    }

    return createUser(b), nil
}
