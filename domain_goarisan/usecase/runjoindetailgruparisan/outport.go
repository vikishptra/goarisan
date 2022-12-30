package runjoindetailgruparisan

import "vikishptra/domain_goarisan/model/repository"

type Outport interface {
	repository.SaveDetailGrupArisanRepo
	repository.FindUserByIDRepo
	repository.FindGrupArisanByIdRepo
	repository.FindGrupArisanByIdRepo
	repository.SaveUserRepo
	repository.SaveGrupArisanRepo
}
