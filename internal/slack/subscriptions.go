package slack

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Subscription struct {
	MessageID string
	Project string
	Tag string
}

func NewSubscription(project string, tag string, messageID string) *Subscription {
	return &Subscription{
		MessageID: messageID,
		Project:   project,
		Tag:       tag,
	}
}

func NewSubscriptionFromKV(key *SubscriptionKey, value string) *Subscription {
	return &Subscription{
		MessageID: value,
		Project: key.Project,
		Tag: key.Tag,
	}
}

func (s *Subscription) GetMessageID () string {
	return s.MessageID
}

func (s *Subscription) GetSubscriptionKey() *SubscriptionKey {
	return NewSubscriptionKey(s.Project, s.Tag)
}

// SubscriptionKey represents key parts
type SubscriptionKey struct {
	Project string
	Tag string
}

func NewSubscriptionKey(project string, tag string) *SubscriptionKey {
	return &SubscriptionKey{
		Project: project,
		Tag:     tag,
	}
}

// Serializes key to owner/repo:tag
// E.g srgkas/most-useful-slackbot:v0.0.1
func (sk *SubscriptionKey) String() string  {
	return sk.Project + ":" + sk.Tag
}

type SubscriptionRepo interface {
	Get(key *SubscriptionKey) (*Subscription, error)
	Store(s *Subscription) error
	Delete(key *SubscriptionKey) error
}

type RedisSubscriptionRepo struct {
	client *redis.Client
}

func (r *RedisSubscriptionRepo) Get(key *SubscriptionKey) (*Subscription, error) {
	ctx := context.Background()
	cmd := r.client.Get(ctx, key.String())
	result, err := cmd.Result()
	if err != nil {
		return nil, err
	}

	return NewSubscriptionFromKV(key, result), nil
}

func (r *RedisSubscriptionRepo) Store(s *Subscription) error {
	ctx := context.Background()
	return r.client.Set(
		ctx,
		s.GetSubscriptionKey().String(),
		s.GetMessageID(),
		0,
	).Err()
}

func (r *RedisSubscriptionRepo) Delete(key *SubscriptionKey) error {
	ctx := context.Background()
	return r.client.Del(ctx, key.String()).Err()
}

func NewSubscriptionRepo(client *redis.Client) SubscriptionRepo {
	return &RedisSubscriptionRepo{client: client}
}