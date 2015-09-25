package api

func (f Frontends) Strings() []string {
	s := []string{}

	for _, r := range f {
		s = append(s, r.Strings()...)
	}

	return s
}

func (b Backends) Strings() []string {
	s := []string{}

	for _, r := range b {
		s = append(s, r.Strings()...)
	}

	return s
}

func (o Options) Strings() []string {
	s := []string{}

	for _, r := range o {
		s = append(s, r.Strings()...)
	}

	return s
}

func (b BackendRules) Strings() []string {
	s := []string{}

	for _, r := range b {
		s = append(s, r.Strings()...)
	}

	return s
}

func (b Servers) Strings() []string {
	s := []string{}

	for _, r := range b {
		s = append(s, r.Strings()...)
	}

	return s
}
