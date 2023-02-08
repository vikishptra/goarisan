package runchangepasswordwithgmail

import "vikishptra/domain_goarisan/model/repository"

type Outport interface {
	repository.FindUserByIDRepo
	repository.SaveUserRepo
}
