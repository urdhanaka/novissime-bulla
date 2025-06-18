func (c *LibvirtVirtualization) createMaster(
	ctx context.Context,
	virtRequest *virtualization_model.CreateInstanceRequest,
) (*virtualization_model.VirtCreateInstanceResponse, error) {
    ...

	// spawning
	slogFunction(virtRequest.Name, thisInstanceName, "spawning node", nil)
	dom, err := c.libvirtConnection.DomainCreateXML(domainXmlConfig, libvirt.DomainCreateFlags(0))
	if err != nil {
		slogFunction(virtRequest.Name, thisInstanceName, "could not spawn node", err)
		c.deleteInstance(thisInstanceName)

		return createRes, err
	}

    ...
}
