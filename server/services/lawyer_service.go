package services

import (
	"../domain"
	"../models"
	"../utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetLawyer(email string, password string) (*models.Lawyers, *utils.ApplicationError) {
	return domain.GetLawyer(email, password)
}

func AddLawyer(lawyerInfo models.Lawyers) *utils.ApplicationError {
	return domain.AddLawyer(lawyerInfo)
}

func GetCase() ([]primitive.M, *utils.ApplicationError) {
	return domain.GetCase()
}

func GetEachCase(user models.Lawyers) ([]primitive.M, *utils.ApplicationError) {
	return domain.GetEachCase(user)
}

func TakeCase(caseID string, user models.Lawyers) *utils.ApplicationError {
	return domain.TakeCase(caseID, user)
}

func CreateMeeting(zoomInfo utils.ZoomMeeting, user models.Lawyers) *utils.ApplicationError {
	return domain.CreateMeeting(zoomInfo, user)
}

func LawyerSendTimeMeeting(meetingTimes utils.LawyerSendTime) *utils.ApplicationError {
	return domain.LawyerSendTimeMeeting(meetingTimes)
}
