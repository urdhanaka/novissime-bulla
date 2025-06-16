func (c *LibvirtVirtualization) createInstance(
    ctx context.Context,
    virtRequest *virtualization_model.CreateInstanceRequest,
) (*virtualization_model.VirtCreateInstanceResponse, error) {
    ...

    slogFunction(virtRequest.Name, thisInstanceName, "waiting until the vm is ready..", nil)
    time.Sleep(CLOUD_INIT_TIMEOUT * time.Second)
    waitCloudInitCmd := "cloud-init status --wait"
    _, err = guestAgentExecStatus(dom, waitCloudInitCmd)
    if err != nil {
        slogFunction(virtRequest.Name, thisInstanceName, "could not spawn node", err)
        c.deleteInstance(thisInstanceName)

        return createRes, err
    }

    ...
}
