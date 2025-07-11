func createBase(
    instanceName string,
    instanceConfig *virtualization_model.CreateInstanceRequest,
) (string, error) {
    ...
    seedFile := POOL_DIR + "/" + instanceName + ".iso"
    ...
    domConfig := &libvirtxml.Domain{
        ...
        Devices: &libvirtxml.DomainDeviceList{
            Emulator: "/usr/bin/qemu-system-x86_64",
            Disks: []libvirtxml.DomainDisk{
                ...
                {
                    ...
                    Source: &libvirtxml.DomainDiskSource{
                        File: &libvirtxml.DomainDiskSourceFile{
                            File: seedFile,
                        },
                    },
                    ...
                },
            },
            ...
        },
        ...
    }
    ...
}
