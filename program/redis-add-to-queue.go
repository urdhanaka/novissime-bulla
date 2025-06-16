func (s *Queue) AddToQueue(
    ctx context.Context,
    instanceRequest virtualization_model.CreateInstanceRequest,
) error {
    requestString, err := json.Marshal(instanceRequest)
    if err != nil {
        return err
    }

    _, err = s.redisClient.LPush(ctx, REDIS_MAIN_QUEUE, string(requestString)).Result()
    if err != nil {
        return err
    }

    return nil
}
