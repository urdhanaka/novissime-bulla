func StartGrpcServer(connection *InitStruct) {
    lis, err := net.Listen("tcp", GRPC_PORT)
    if err != nil {
        slog.Error("could not start grpc server",
            "error", err.Error(),
        )
        os.Exit(1)
    }

    s := grpc.NewServer()
    proto_model.RegisterNodeServiceServer(s, &NodeServer{
        queue: connection.QueueService,
    })

    // background job
    go startWorker(connection)

    // connect to main server
    slog.Info("node is ready, connecting to main server")
    res, err := connectToServer()
    if err != nil {
        slog.Error("StartGrpcServer(): failed to connect to main server",
            "error", err.Error(),
        )
        os.Exit(1)
    }
    slog.Info("main server is responding",
        "response", res)

    slog.Info(fmt.Sprintf("starting grpc server at %s", GRPC_PORT))

    if err = s.Serve(lis); err != nil {
        slog.Error("StartGrpcServer(): failed to serve",
            "error", err.Error(),
        )
        os.Exit(1)
    }
}
