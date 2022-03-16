package bridge

import (
	"errors"
	"strings"
)

// Bridge Design Pattern
// 将抽象和实现解耦，让它们可以独立变化

type db interface {
	Open()
	Close()
	Query()
}

type DbDriveEnum uint

var (
	driveMap = map[string]db{
		"mysql":  mysqlDrive{},
		"mongo":  mongoDrive{},
		"oracle": oracleDrive{},
	}
	WrongUrlError    = errors.New("url不能为空")
	UnSupportDbError = errors.New("不支持类型")
)

type service struct {
	db
}

func NewService(url string) (*service, error) {
	idx := strings.Index(url, ":")
	if len(url) == 0 || idx == -1 {
		return nil, WrongUrlError
	}

	drive, ok := driveMap[url[:idx]]
	if !ok {
		return nil, UnSupportDbError
	}
	return &service{
		db: drive,
	}, nil
}

// tip: mysql
type mysqlDrive struct{}

func (d mysqlDrive) Open() {
	// do something
	println("mysql")
}
func (d mysqlDrive) Close() {
	// do something
}
func (d mysqlDrive) Query() {
	// do something
}

// tip: mongo
type mongoDrive struct{}

func (d mongoDrive) Open() {
	// do something
	println("mongo")
}
func (d mongoDrive) Close() {
	// do something
}
func (d mongoDrive) Query() {
	// do something
}

// tip: oracle
type oracleDrive struct{}

func (d oracleDrive) Open() {
	// do something
	println("oracle")
}
func (d oracleDrive) Close() {
	// do something
}
func (d oracleDrive) Query() {
	// do something
}
