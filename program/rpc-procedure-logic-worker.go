func (s *NodeServer) CreateWorker(
    ctx context.Context,
    createWorkerRequest *proto_model.CreateWorkerRequest,
) (*proto_model.CreateWorkerResponse, error) {
    provisionCtx, cancel := context.WithCancel(ctx)
    defer cancel()

    res := new(proto_model.CreateWorkerResponse)

    instanceName := createWorkerRequest.Requirements.NodeName

    virtSpecs := virtualization_model.CreateInstanceRequest{
        Name:            instanceName,
        IsMaster:        false,
        MasterIpAddress: createWorkerRequest.MasterIpAddress,
        Token:           createWorkerRequest.ClusterToken,
        Cpu:             createWorkerRequest.Requirements.Cpu,
        Memory:          createWorkerRequest.Requirements.Memory,
        Storage:         createWorkerRequest.Requirements.Storage,
    }

    err := s.queue.AddToQueue(provisionCtx, virtSpecs)
    if err != nil {
        return res, err
    }

    go sendLogs(instanceName, createWorkerRequest.ClusterName)

    return res, nil
}
