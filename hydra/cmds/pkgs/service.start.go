package pkgs

import "psbnb.com/greatsun/hydra/hydra/cmds/pkgs/service"

// Start Start
func (p *ServiceApp) Start(s service.Service) (err error) {
	return p.run()
}
