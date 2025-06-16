func createBase(
    instanceName string,
    instanceConfig *virtualization_model.CreateInstanceRequest,
) (string, error) {
    instanceStorage := POOL_DIR + "/" + instanceName + ".qcow2"
    seedFile := POOL_DIR + "/" + instanceName + ".iso"
    logSocket := INSTANCE_LOGS_DIR + "/" + instanceName + ".sock"

    domConfig := &libvirtxml.Domain{
        Type: "kvm",
        Name: instanceName,
        ...
    }
    xmlConfig, err := domConfig.Marshal()
    if err != nil {
        return "", err
    }

    // hacky way to handle <channel>
    // for qemu-guest-agent channel
    res := strings.Replace(xmlConfig, "<channel>", `<channel type="unix">`, 1)

    return res, nil
}
