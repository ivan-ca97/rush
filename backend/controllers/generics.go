package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	ce "github.com/ivan-ca97/rush/backend/custom_errors"
)

func GetIdFromParameters(request *http.Request) (uuid.UUID, error) {
	id, err := uuid.Parse(chi.URLParam(request, "id"))
	if err != nil {
		err = &ce.RequestError{Message: "Invalid ID"}
		return uuid.UUID{}, err
	}

	return id, nil
}

func DecodeRequestBody[BT any](request *http.Request) (*BT, error) {
	var body BT

	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		return nil, &ce.RequestError{Message: "Invalid request payload"}
	}

	return &body, nil
}

func Create[BT any, RT any](controller ControllersBasic, writer http.ResponseWriter, request *http.Request, service func(BT) (RT, error)) {
	body, err := DecodeRequestBody[BT](request)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}

	response, err := service(*body)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}

	err = controller.EncodeResponse(writer, response)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}
}

func GetAllPaginated[RT any](controller ControllersBasic, writer http.ResponseWriter, request *http.Request, service func(int, int) (RT, error)) {
	pageNumber, pageSize, err := controller.GetPaginationParameters(request)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}

	response, err := service(pageNumber, pageSize)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}

	err = controller.EncodeResponse(writer, response)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}
}

func GetByUuid[RT any](controller ControllersBasic, writer http.ResponseWriter, request *http.Request, service func(uuid.UUID) (RT, error)) {
	id, err := GetIdFromParameters(request)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}

	response, err := service(id)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}

	err = controller.EncodeResponse(writer, response)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}
}

func UpdateById[BT any, RT any](controller ControllersBasic, writer http.ResponseWriter, request *http.Request, service func(uuid.UUID, BT) (RT, error)) {

	id, err := GetIdFromParameters(request)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}

	body, err := DecodeRequestBody[BT](request)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}

	response, err := service(id, *body)

	err = controller.EncodeResponse(writer, response)
	if err != nil {
		controller.HandleError(writer, err)
		return
	}
}
