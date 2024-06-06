package usecase

import (
	"play-cdc/domain"
	"play-cdc/repository"
)

func GenerateSpec() {
	for _, env := range repository.GetProviderEnvs() {
		spec, contracts := load(env)

		errors := save(env, spec, contracts)
		if errors != nil {
			panic(errors)
		}
	}
}

func load(env domain.ProviderEnv) (*domain.Spec, domain.Contracts) {
	requests := repository.LoadRecordedRequests(env.ProviderEndpoint)

	contracts := domain.ToContracts(requests)

	spec := contracts.ToSpec(repository.GetConsumerName(), env.ProviderName)
	spec.SortScenarios()

	return spec, contracts
}

func save(env domain.ProviderEnv, spec *domain.Spec, contracts domain.Contracts) []error {
	err := repository.SaveSpec(spec, domain.SpecFilePath(outputPath(env), repository.GetConsumerName()))
	if err != nil {
		return []error{err}
	}

	errors := repository.SaveRequestBodies(contracts, domain.RequestBodiesPath(outputPath(env), repository.GetConsumerName()))
	if len(errors) > 0 {
		return errors
	}

	return nil
}

func outputPath(env domain.ProviderEnv) string {
	outputBasePath := repository.GetOutputBasePath()
	return domain.OutputPath(outputBasePath, env.ProviderName)
}
