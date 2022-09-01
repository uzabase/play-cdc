package usecase

import "play-cdc/repository"

func InitRequests() {
	repository.ClearExecutedRequests()
}
