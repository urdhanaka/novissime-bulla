func (c *LibvirtVirtualization) createInstance(
    ctx context.Context,
    virtRequest *virtualization_model.CreateInstanceRequest,
) (*virtualization_model.VirtCreateInstanceResponse, error) {
    ...

    // base xml for the vm
    slogFunction(virtRequest.Name, thisInstanceName, "creating instance base xml", nil)
    domainXmlConfig, err := createBase(thisInstanceName, virtRequest)
    if err != nil {
        slogFunction(virtRequest.Name, thisInstanceName, "could not create node base xml", err)
        c.deleteInstance(thisInstanceName)

        return createRes, err
    }

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
