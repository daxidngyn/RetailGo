package transactions

import (
	"context"
	"fmt"

	"github.com/hktrib/RetailGo/internal/ent"
)

func StoreAndOwnerCreationTx(ctx context.Context, reqStore *ent.Store, reqUser *ent.User, dbClient *ent.Client) error {
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: starting a transaction: %w", err))
	}

	store, err := tx.Store.Create().SetStoreName(reqStore.StoreName).SetOwnerEmail(reqStore.OwnerEmail).
		SetStoreAddress(reqStore.StoreAddress).SetStoreType(reqStore.StoreType).Save(ctx)

	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to create store: %w", err))
	}
	_, err = tx.User.Create().
		SetEmail(reqUser.Email).
		SetIsOwner(reqUser.IsOwner).
		SetFirstName(reqUser.FirstName).
		SetLastName(reqUser.LastName).
		SetStoreID(store.ID).AddStoreIDs(store.ID).Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to create owner: %w", err))
	}

	_, err = tx.UserToStore.Update().
		AddPermissionLevel(0).Save(ctx)
	if err != nil {
		return rollback(tx, fmt.Errorf("tx_error: Unable to create owner: %w", err))
	}
	return tx.Commit()
}

func rollback(tx *ent.Tx, err error) error {

	// fmt.Println("Hit Rollback!!")
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
		fmt.Println(err)
	}
	return err
}
