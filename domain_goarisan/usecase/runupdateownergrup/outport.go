package runupdateownergrup

import "vikishptra/domain_goarisan/model/repository"

type Outport interface {
	repository.RunUpdateOwnerGrupRepo
	repository.SaveGrupArisanRepo
}
