func (n *NodeServer) CreateMaster(
	c context.Context,
	masterRequirement *proto_model.CreateMasterRequest,
) (*proto_model.CreateMasterResponse, error) {
	masterIpAddress, err := n.virtualizationService.CreateMaster(
		context.Background(),
		virtualization_model.NodeCreateRequest{
			IsMaster: true,
			Token:    masterRequirement.Token,
			Cpu:      masterRequirement.Requirements.GetCpu(),
			Memory:   masterRequirement.Requirements.GetMemory(),
			Storage:  masterRequirement.Requirements.GetStorage(),
		})
	if err != nil {
		return new(proto_model.CreateMasterResponse), err
	}

	return &proto_model.CreateMasterResponse{
		IpAddress: masterIpAddress,
	}, nil
}

func (n *NodeServer) CreateWorker(
	c context.Context,
	workerRequest *proto_model.CreateWorkerRequest,
) (*proto_model.CreateWorkerResponse, error) {
	err := n.virtualizationService.CreateWorker(context.Background(), virtualization_model.NodeCreateRequest{
		IsMaster:        false,
		MasterIpAddress: workerRequest.IpAddress,
		Token:           workerRequest.Token,
		Cpu:             workerRequest.Requirements.GetCpu(),
		Memory:          workerRequest.Requirements.GetMemory(),
		Storage:         workerRequest.Requirements.GetStorage(),
	})
	if err != nil {
		return new(proto_model.CreateWorkerResponse), err
	}

	return new(proto_model.CreateWorkerResponse), nil
}