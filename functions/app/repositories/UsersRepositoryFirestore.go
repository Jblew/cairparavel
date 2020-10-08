import 	(
	"cloud.google.com/go/firestore"
	"context"
)

type UsersRepositoryFirestore struct {
	Firestore *firestore.Client
	Context *context.Context
}

var usersCols = []string{
	"envs/prod/users",
	"envs/test/users",
}

func (UsersRepositoryFirestore *repo) StoreUser(user User) error {
	for _, usersCol := range usersCols {
		docRef := repo.Firestore.Collection(usersCol).Doc(userDoc.UID)
		_, err := docRef.Create(repo.Context, userDoc)
		if err != nil {
			return err
		}
	}
	return nil
}
