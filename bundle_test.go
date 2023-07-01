package cli

// Test CreateCustomBrandingDefinitionDTO CRUD operations
func (s *AccTestSuite) TestAccCliBundles_get() {

	var t = s.T()

	s.client.Logger().Infof("Testing get bundles")

	b, err := s.client.GetOSGiBundles()

	if err != nil {
		t.Error(err)
		return
	}

	// print bundle names:
	for _, bundle := range b {
		s.client.Logger().Infof("Bundle %s", *bundle.Name)
	}

}
