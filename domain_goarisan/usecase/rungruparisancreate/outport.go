package rungruparisancreate

import "vikishptra/domain_goarisan/model/repository"

type Outport interface {
	repository.FindUserByIDRepo
	repository.SaveGrupArisanRepo
	repository.SaveUserRepo
	repository.SaveDetailGrupArisanRepo
}
