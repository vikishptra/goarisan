package runupdategruparisanbyidowner

import "vikishptra/domain_goarisan/model/repository"

type Outport interface {
	repository.FindOneGrupByOwnerRepo
	repository.SaveGrupArisanRepo
}
