package api

import "fmt"

func (j *JSONInput) Strings() []string {
	s := []string{}

	s = append(s, j.Frontends.Strings()...)
	s = append(s, j.Backends.Strings()...)

	return s
}

func (f Frontend) Strings() []string {
	s := []string{
		fmt.Sprintf("frontend %s", f.Name),
		fmt.Sprintf("\tbind %s:%d", f.IP, f.Port),
	}

	s = append(s, f.BackendRules.Strings()...)

	return s
}

func (b Backend) Strings() []string {
	s := []string{
		fmt.Sprintf("backend %s", b.Name),
	}

	s = append(s, b.BalanceType.Strings()...)
	s = append(s, b.Mode.Strings()...)
	s = append(s, b.Options.Strings()...)
	s = append(s, b.Servers.Strings()...)

	return s
}

func (b RuleUseBackend) Strings() []string {
	return []string{fmt.Sprintf("\tuse_backend %s", b.Name)}
}

func (b BalanceType) Strings() []string {
	return []string{fmt.Sprintf("\tbalance %s", b)}
}

func (m Mode) Strings() []string {
	return []string{fmt.Sprintf("\tmode %s", m)}
}

func (o Option) Strings() []string {
	return []string{fmt.Sprintf("\toption %s", o)}
}

func (s Server) Strings() []string {
	return []string{fmt.Sprintf("\tserver %s %s:%d", s.Name, s.IP, s.Port)}
}
