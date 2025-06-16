func (s *Queue) PopQueue(
    ctx context.Context,
) (*virtualization_model.CreateInstanceRequest, error) {
    instanceRequest := new(virtualization_model.CreateInstanceRequest)

    job, err := s.redisClient.BRPop(ctx, 0, REDIS_MAIN_QUEUE).Result()
    if err != nil {
        return instanceRequest, err
    }

    // job value is the second element at the array
    err = json.Unmarshal([]byte(job[1]), instanceRequest)
    if err != nil {
        return instanceRequest, err
    }

    return instanceRequest, nil
}
