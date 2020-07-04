package nepodate

import (
	"context"
	"testing"
	time2 "time"
)

func setup()(srv Service, ctx context.Context)  {
	return NewService(),context.Background()

}
func TestDateService_Status(t *testing.T) {
	srv,ctx:=setup()

	s,err := srv.Status(ctx)
	if err != nil{
		t.Errorf("Error :%s",err)
	}

	//testing status
	if s !="ok"{
		t.Errorf("expected sevice to be ok")
	}
}

func TestDateService_Get(t *testing.T) {
	srv,ctx := setup()

	d,err := srv.Get(ctx)
	if err != nil{
		t.Errorf("Error :%s",err)
	}

	time:= time2.Now()
	today := time.Format("02/01/2006")

	if  today != d{
		t.Errorf("expected dates to be equal")
	}
}

func TestDateService_Validate(t *testing.T) {
	srv,ctx := setup()
	b,err := srv.Validate(ctx,"31/12/2020")
	if err != nil{
		t.Errorf("Error :%s",err)
	}
	//testing that the date is valid
	if !b {
		t.Errorf("date should be valid")
	}

	// testing an invalid date
	b,err = srv.Validate(ctx,"31/31/2020")
	if b{
		t.Errorf("date should be invalid")
	}

	//testing a USA date
	b,err = srv.Validate(ctx,"12/31/2020")
	if b{
		t.Errorf("date should be invalid")
	}
}