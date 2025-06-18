func (c *LibvirtVirtualization) deleteInstance(
    domainName string,
) {
    domain, err := c.libvirtConnection.LookupDomainByName(domainName)
    if err != nil {
        slog.Error("error getting the domain",
            "error", err,
        )
    }

    // every domain is set to delete when shutdown
    err = domain.Shutdown()
    if err != nil {
        slog.Error("could not shutdown the domain, retrying after this",
            "error", err,
        )

        for i := 1; i <= SHUTDOWN_RETRIES; i++ {
            err = domain.Shutdown()
            if err != nil {
                slog.Error("could not shutdown the domain, retrying after this",
                    "error", err,
                )
            }
        }
    }

    // domain files cleanup
    deleteFilesCommand := fmt.Sprintf("rm %s/%s.*", POOL_DIR, domainName)
    cmd := exec.Command("/bin/bash", "-c", deleteFilesCommand)
    err = cmd.Run()
    if err != nil {
        slog.Error("could not clean domain files",
            "error", err,
        )
    }
    deleteFilesCommand = fmt.Sprintf("rm %s/%s.*", NVRAM_DIR, domainName)
    cmd = exec.Command("/bin/bash", "-c", deleteFilesCommand)
    err = cmd.Run()
    if err != nil {
        slog.Error("could not clean domain files",
            "error", err,
        )
    }
}
