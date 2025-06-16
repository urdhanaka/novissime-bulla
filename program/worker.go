func (w *Worker) DoWork() {
    ctx := context.Background()

    for {
        instanceRequest, err := w.queue.PopQueue(ctx)
        if err != nil {
            slog.Error("error creating instance",
                "error", err,
                )
            continue
        }

        res, err := w.virtService.CreateInstance(ctx, instanceRequest)
        if err != nil {
            slog.Error("error creating instance",
                "error", err,
                )
            continue
        }

        jsonBytes, err := json.Marshal(res)
        if err != nil {
            slog.Error("error creating instance",
                "error", err,
                )
            continue
        }

        err = w.queue.Publish(ctx, instanceRequest.Name, string(jsonBytes))
        if err != nil {
            slog.Error("error creating instance",
                "error", err,
                )
            continue
        }
    }
}
