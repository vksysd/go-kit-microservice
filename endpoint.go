package nepodate

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetEndpoint endpoint.Endpoint
	StatusEndpoint endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

//MakeFetEndpoint returns the response from our service "Get"
func MakeGetEndpoint(srv Service) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// we use type assertion here to validate that the incoming request is of type "getRequest"
		_ = request.(getRequest) // we just need the request, we don't use any value from it
		d,err := srv.Get(ctx)
		if err != nil{
			return getResponse{
				Date: d,
				Err:  err.Error(),
			}, nil
		}
		return getResponse{
			Date: d,
			Err:  "",
		}, nil
	}
}

func MakeStatusEndpoint(srv Service) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_= request.(statusRequest)
		s, err := srv.Status(ctx)
		if err !=nil{
			return statusResponse{s}, err
		}
		return statusResponse{s}, nil
	}
}

func MakeValidateEndpoint(srv Service)endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(validateRequest)
		s,err := srv.Validate(ctx,req.Date)
		if err!= nil{
			return validateResponse{
				Valid: s,
				Err:   err.Error(),
			},err
		}
		return validateResponse{
			Valid: s,
			Err: "",
		},nil
	}
}

//Get endpoint mapping
func (e Endpoints) Get(ctx context.Context) (string, error)  {
	req:= getRequest{}
	resp,err := e.GetEndpoint(ctx,req)
	if err != nil{
		return "", err
	}
	getResp := resp.(getResponse)
	if getResp.Err != ""{
		return "", errors.New(getResp.Err)
	}
	return getResp.Date, nil
}

// Status endpoint mapping
func (e Endpoints) Status(ctx context.Context)(string,error)  {
	req := statusRequest{}
	resp,err := e.StatusEndpoint(ctx,req)
	if err != nil{
		return "", err
	}
	statusResp := resp.(statusResponse)
	return statusResp.Status, nil
}

// Validate endpoint mapping
func (e Endpoints)Validate(ctx context.Context, date string)(bool,error)  {
	req := validateRequest{Date: date}
	resp, err := e.ValidateEndpoint(ctx,req)
	if err != nil{
		return false, err
	}
	validateResp := resp.(validateResponse)
	if validateResp.Err != ""{
		return validateResp.Valid,errors.New(validateResp.Err)
	}
	return validateResp.Valid,nil
}
