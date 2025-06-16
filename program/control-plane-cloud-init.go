func createCloudInitMaster(instanceName string, clusterToken string) error {
    cloudInitMut.Lock()
    defer cloudInitMut.Unlock()

    filePath := BASE_POOL_DIR + "/" + "user-data"
    networkPath := BASE_POOL_DIR + "/" + "network-config"
    userDataContent := fmt.Sprintf(`#cloud-config
        ...
        `, instanceName, clusterToken)

    err := os.WriteFile(filePath, []byte(userDataContent), 0644)
    if err != nil {
        return err
    }

    // create the iso
    cmd := exec.Command("cloud-localds", "-N", networkPath, POOL_DIR+"/"+instanceName+".iso", filePath)
    err = cmd.Run()
    if err != nil {
        return err
    }

    return nil
}
