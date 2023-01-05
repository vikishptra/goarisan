package findgrupbyiduser

import "vikishptra/domain_goarisan/model/repository"

type Outport interface {
	repository.GetfindgrupbyidownerRepo
	repository.FindUserByIDRepo
}
