func createNetwork() error {
    networkMut.Lock()
    defer networkMut.Unlock()

    filePath := BASE_POOL_DIR + "/" + "network-config"
    userDataContent := `network:
  version: 2
  ethernets:
    enp1s0:
      dhcp4: true`

    err := os.WriteFile(filePath, []byte(userDataContent), 0644)
    if err != nil {
        return err
    }

    return nil
}
