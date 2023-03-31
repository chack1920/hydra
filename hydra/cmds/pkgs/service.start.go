package pkgs

import "github.com/chack1920/hydra/hydra/cmds/pkgs/service"

// Start Start
func (p *ServiceApp) Start(s service.Service) (err error) {
	return p.run()
}
