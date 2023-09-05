package initapp

type DependencyService interface {
	Init() error
	Start() error
	Stop() error
}

type DependenciesInitializer struct {
	dependencyInitChain []DependencyService
}

func (i *DependenciesInitializer) InitDependencies() error {
	for _, dependencyInit := range i.dependencyInitChain {
		err := dependencyInit.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *DependenciesInitializer) StartDependencies() error {
	for _, dependencyInit := range i.dependencyInitChain {
		err := dependencyInit.Start()
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *DependenciesInitializer) StopDependencies() error {
	for _, dependencyInit := range i.dependencyInitChain {
		err := dependencyInit.Stop()
		if err != nil {
			return err
		}
	}

	return nil
}
