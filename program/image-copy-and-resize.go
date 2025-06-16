func copyImage(
    instanceName string,
    virtRequest *virtualization_model.CreateInstanceRequest,
) error {
    imageMut.Lock()
    defer imageMut.Unlock()

    baseImage := BASE_POOL_DIR + "/" + BASE_IMAGE_NAME
    destinationPath := POOL_DIR + "/" + instanceName + ".qcow2"

    data, err := os.ReadFile(baseImage)
    if err != nil {
        return err
    }

    err = os.WriteFile(destinationPath, data, 0644)
    if err != nil {
        return err
    }

    // resize the qcow2
    // resizeCmd := exec.Command("qemu-img", "resize", destinationPath, "+10G")
    resizeCmd := exec.Command("qemu-img", "resize", destinationPath, fmt.Sprintf("+%dG", virtRequest.Storage))
    err = resizeCmd.Run()
    if err != nil {
        return err
    }

    return nil
}
