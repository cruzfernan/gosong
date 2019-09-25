package song

import (
	"fmt"
	"github.com/BitSongOfficial/go-bitsong/x/song/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - song genesis state
type GenesisState struct {
	Pool           Pool         `json:"pool"`
	Params         types.Params `json:"params"`
	SongTax        sdk.Dec      `json:"song_tax"`
	StartingSongID uint64       `json:"starting_id"`
	Songs          Songs        `json:"songs"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(pool Pool, startingSongID uint64, params types.Params) GenesisState {
	return GenesisState{
		Pool:           pool,
		StartingSongID: startingSongID,
		Params:         params,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Pool:           types.InitialPool(),
		StartingSongID: DefaultStartingSongID,
		Params:         types.DefaultParams(),
	}
}

// ValidateGenesis validates genesis state
func ValidateGenesis(data GenesisState) error {
	if data.SongTax.IsNegative() || data.SongTax.GT(sdk.OneDec()) {
		return fmt.Errorf("song parameter SongTax should non-negative and "+
			"less than one, is %s", data.SongTax.String())
	}

	return data.Pool.ValidateGenesis()
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	keeper.SetParams(ctx, data.Params)

	err := keeper.SetInitialSongID(ctx, data.StartingSongID)
	if err != nil {
		panic(err)
	}
	/*for _, song := range data.Songs {
		keeper.AddSong(ctx, *song)
	}*/

}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	startingSongID, _ := keeper.PeekCurrentSongID(ctx)
	return NewGenesisState(Pool{}, startingSongID, keeper.GetParams(ctx))
}
