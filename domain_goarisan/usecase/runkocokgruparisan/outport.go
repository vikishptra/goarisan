package runkocokgruparisan

import "vikishptra/domain_goarisan/model/repository"

type Outport interface {
	repository.FindUndianArisanUserRepo
	repository.SaveDetailGrupArisanRepo
	repository.FindOneGrupByOwnerRepo
}
