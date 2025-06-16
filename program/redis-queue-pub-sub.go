func (s *Queue) Subscribe(ctx context.Context, channel string) *redis.PubSub {
    return s.redisClient.Subscribe(ctx, channel)
}

func (s *Queue) Publish(
    ctx context.Context,
    channel string,
    message any,
) error {
    _, err := s.redisClient.Publish(ctx, channel, message).Result()
    if err != nil {
        return err
    }

    return nil
}
