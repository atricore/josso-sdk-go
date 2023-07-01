package cli

// Test CreateCustomBrandingDefinitionDTO CRUD operations
func (s *AccTestSuite) TestServerInfo_get() {

	var t = s.T()

	s.client.Logger().Infof("Testing get server info")

	b, err := s.client.GetInfo()

	if err != nil {
		t.Error(err)
		return
	}

	s.client.Logger().Infof("Server info: %s/%s", *b.NodeId, *b.Version)

}
