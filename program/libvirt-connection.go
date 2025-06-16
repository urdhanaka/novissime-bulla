func InitLibvirtConnection() *libvirt.Connect {
    c, err := libvirt.NewConnect("qemu:///system")
    if err != nil {
        slog.Error("error connecting to QEMU system",
            "error", err.Error(),
        )
        os.Exit(1)
    }

    return c
}
