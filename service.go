package nepodate

import (
	"context"
	"time"
)

// Service provides some "date capabilities" to your app
type Service interface {
	Status(ctx context.Context) (string,error)
	Get (ctx context.Context) (string,error)
	Validate(ctx context.Context,date string) (bool,error)
}

type dateService struct {}

func (d dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

func (d dateService) Get(ctx context.Context) (string, error) {
	now:= time.Now()
	//layouts must use the reference time Jan 2 15:04:05 MST 2006 to show the pattern with which to format a given time/string
	return now.Format("02/01/2006"), nil
}

func (d dateService) Validate(ctx context.Context, date string) (bool, error) {
	_,err := time.Parse("02/01/2006",date)
	if err != nil{
		return false, err
	}
	return true, nil
}

func NewService() Service{
	return dateService{}
}


