package currentTime

import (
	"context"
	"time"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) CurrentTime(ctx context.Context) (*string, error) {
	nowTime := time.Now()
	representation := nowTime.Format("2006-01-02 15:04:05")
	return &representation, nil
}
