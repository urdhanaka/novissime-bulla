func guestAgentExecStatus(
    dom *libvirt.Domain,
    execString string,
) (*ExecStatusQemuGuestAgent, error) {
    ...

    for {
        cmd := fmt.Sprintf(`{"execute":"guest-exec-status","arguments":{"pid":%d}}`, pidStruct.Return.PID)
        status, err := dom.QemuAgentCommand(cmd, libvirt.DOMAIN_QEMU_AGENT_COMMAND_BLOCK, 0)
        if err != nil {
            return res, err
        }

        ...

        time.Sleep(10 * time.Second)
    }

    return res, nil
}
