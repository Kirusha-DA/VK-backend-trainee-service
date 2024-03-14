package actor

type Repository interface {
	CreateActor(actor *Actor)
	UpdateActor(actor *Actor)
	PartiallyUpdateActor(actor *Actor)
	DeleteActor(actorID string)
}
