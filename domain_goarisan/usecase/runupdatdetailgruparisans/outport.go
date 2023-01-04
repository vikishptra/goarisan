package runupdatdetailgruparisans

import "vikishptra/domain_goarisan/model/repository"

type Outport interface {
	repository.FindoneuserdetailgruparisansRepo
	repository.SaveDetailGrupArisanRepo
	repository.FindGrupArisanByIdRepo
	repository.FindUserByIDRepo
	repository.SaveUserRepo
}
