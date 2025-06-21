func (s *NodeServer) CreateMaster(
    ctx context.Context,
    createMasterRequest *proto_model.CreateMasterRequest,
) (*proto_model.CreateMasterResponse, error) {
    provisionCtx, cancel := context.WithCancel(ctx)
    defer cancel()

    res := new(proto_model.CreateMasterResponse)

    instanceName := createMasterRequest.Requirements.NodeName

    virtSpecs := virtualization_model.CreateInstanceRequest{
        Name:     instanceName,
        IsMaster: true,
        Token:    createMasterRequest.ClusterToken,
        Cpu:      createMasterRequest.Requirements.Cpu,
        Memory:   createMasterRequest.Requirements.Memory,
        Storage:  createMasterRequest.Requirements.Storage,
    }

    err := s.queue.AddToQueue(provisionCtx, virtSpecs)
    if err != nil {
        return res, err
    }

    go sendLogs(instanceName, createMasterRequest.ClusterName)

    sub := s.queue.Subscribe(ctx, instanceName)
    defer sub.Close()

    msgCh := sub.Channel()
    select {
    case msg := <-msgCh:
        instanceRes := new(virtualization_model.VirtCreateInstanceResponse)

        _ = json.Unmarshal([]byte(msg.Payload), instanceRes)

        res.DashboardToken = instanceRes.DashboardToken
        res.MasterIpAddress = instanceRes.MasterIpAddress

    case <-time.After(PROVISIONING_TIMEOUT * time.Second):
        fmt.Println("timeout exceeded")
    }

    return res, nil
}
