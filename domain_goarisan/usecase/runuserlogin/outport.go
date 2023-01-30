package runuserlogin

import "vikishptra/domain_goarisan/model/repository"

type Outport interface {
	repository.RunLoginRepo
	repository.SaveUserRepo
}
