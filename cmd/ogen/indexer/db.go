package indexer

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/golang-migrate/migrate/database"
	_ "github.com/mattn/go-sqlite3"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/database/sqlite3"

	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/olympus-protocol/ogen/pkg/logger"
	"github.com/olympus-protocol/ogen/pkg/primitives"
	"sync"
)

// State represents the last block saved
type State struct {
	Blocks        int
	LastBlockHash string
}

// Database represents an DB connection
type Database struct {
	log      logger.Logger
	db       *sql.DB
	canClose *sync.WaitGroup
	driver   string
}

func (d *Database) GetCurrentState() (State, error) {
	nextH, prevH, err := d.getNextHeight()
	if err != nil {
		return State{}, err
	}
	if nextH == 0 {
		return State{
			Blocks:        0,
			LastBlockHash: "",
		}, nil
	} else {
		return State{Blocks: nextH - 1, LastBlockHash: prevH}, nil
	}
}

func (d *Database) InsertBlock(block primitives.Block) error {

	nextHeight, prevHash, err := d.getNextHeight()
	if err != nil {
		d.log.Error(err)
		return err
	}

	if nextHeight > 0 && hex.EncodeToString(block.Header.PrevBlockHash[:]) != prevHash {
		d.log.Error(errorPrevBlockHash)
		return errorPrevBlockHash
	}

	// Insert into blocks table
	var queryVars []interface{}
	hash := block.Hash()
	queryVars = append(queryVars, hash.String(), hex.EncodeToString(block.Signature[:]), hex.EncodeToString(block.RandaoSignature[:]), nextHeight)
	err = d.insertRow("blocks", queryVars)
	if err != nil {
		d.log.Error(err)
	}

	// Block Headers
	queryVars = nil
	queryVars = append(queryVars, hash.String(), int(block.Header.Version), int(block.Header.Nonce),
		hex.EncodeToString(block.Header.TxMerkleRoot[:]), hex.EncodeToString(block.Header.TxMultiMerkleRoot[:]), hex.EncodeToString(block.Header.VoteMerkleRoot[:]),
		hex.EncodeToString(block.Header.DepositMerkleRoot[:]), hex.EncodeToString(block.Header.ExitMerkleRoot[:]), hex.EncodeToString(block.Header.VoteSlashingMerkleRoot[:]),
		hex.EncodeToString(block.Header.RANDAOSlashingMerkleRoot[:]), hex.EncodeToString(block.Header.ProposerSlashingMerkleRoot[:]),
		hex.EncodeToString(block.Header.GovernanceVotesMerkleRoot[:]), hex.EncodeToString(block.Header.PrevBlockHash[:]),
		int(block.Header.Timestamp), int(block.Header.Slot), hex.EncodeToString(block.Header.StateRoot[:]),
		hex.EncodeToString(block.Header.FeeAddress[:]))
	err = d.insertRow("block_headers", queryVars)
	if err != nil {
		d.log.Error(err)
	}

	// Votes
	for _, vote := range block.Votes {
		queryVars = nil
		queryVars = append(queryVars, hash.String(), hex.EncodeToString(vote.Sig[:]), hex.EncodeToString(vote.ParticipationBitfield), int(vote.Data.Slot), int(vote.Data.FromEpoch),
			hex.EncodeToString(vote.Data.FromHash[:]), int(vote.Data.ToEpoch), hex.EncodeToString(vote.Data.ToHash[:]), hex.EncodeToString(vote.Data.BeaconBlockHash[:]),
			int(vote.Data.Nonce), vote.Data.Hash().String())
		err = d.insertRow("votes", queryVars)
		if err != nil {
			d.log.Error(err)
			continue
		}
	}

	// Transactions Single
	for _, tx := range block.Txs {
		queryVars = nil
		queryVars = append(queryVars, hash.String(), 0, hex.EncodeToString(tx.To[:]), hex.EncodeToString(tx.FromPublicKey[:]),
			int(tx.Amount), int(tx.Nonce), int(tx.Fee), hex.EncodeToString(tx.Signature[:]))
		err = d.insertRow("tx_single", queryVars)
		if err != nil {
			d.log.Error(err)
			continue
		}
	}

	for _, deposit := range block.Deposits {
		queryVars = nil
		queryVars = append(queryVars, hash.String(), hex.EncodeToString(deposit.PublicKey[:]), hex.EncodeToString(deposit.Signature[:]),
			hex.EncodeToString(deposit.Data.PublicKey[:]), hex.EncodeToString(deposit.Data.ProofOfPossession[:]), hex.EncodeToString(deposit.Data.WithdrawalAddress[:]))
		err = d.insertRow("deposits", queryVars)
		if err != nil {
			d.log.Error(err)
			continue
		}
	}

	// Exits
	for _, exits := range block.Exits {
		queryVars = nil
		queryVars = append(queryVars, hash.String(), hex.EncodeToString(exits.ValidatorPubkey[:]), hex.EncodeToString(exits.WithdrawPubkey[:]),
			hex.EncodeToString(exits.Signature[:]))
		err = d.insertRow("exits", queryVars)
		if err != nil {
			d.log.Error(err)
			continue
		}
	}

	// Vote Slashings
	//for _, vs := range block.VoteSlashings {
	//	queryVars = nil
	//	queryVars = append(queryVars, hash, vote1Int, vote2Int)
	//	err = d.insertRow("vote_slashings", queryVars)
	//	if err != nil {
	//		d.log.Error(err)
	//		continue
	//	}
	//}
	//
	//// RANDAO Slashings
	//for _, rs := range block.RANDAOSlashings {
	//	queryVars = nil
	//	queryVars = append(queryVars, hash, hex.EncodeToString(rs.RandaoReveal[:]), int(rs.Slot), hex.EncodeToString(rs.ValidatorPubkey[:]))
	//	err = d.insertRow("RANDAO_slashings", queryVars)
	//	if err != nil {
	//		d.log.Error(err)
	//		continue
	//	}
	//}
	//
	//// Proposer Slashings
	//for _, ps := range block.ProposerSlashings {
	//	queryVars = nil
	//	queryVars = append(queryVars, hash, bh1Int, bh2Int, hex.EncodeToString(ps.Signature1[:]),
	//		hex.EncodeToString(ps.Signature2[:]))
	//	err = d.insertRow("proposer_slashings", queryVars)
	//	if err != nil {
	//		d.log.Error(err)
	//		continue
	//	}
	//}

	return nil
}

func (d *Database) getNextHeight() (int, string, error) {
	dw := goqu.Dialect(d.driver)
	ds := dw.From("blocks").Select("height")
	query, _, err := ds.ToSQL()
	if err != nil {
		return 0, "", err
	}
	var height string
	err = d.db.QueryRow(query).Scan(&height)
	if err != nil {
		return 0, "", nil
	}
	dw = goqu.Dialect(d.driver)
	ds = dw.From("blocks").Select("block_hash").Where(goqu.C("height").Eq(height))
	query, _, err = ds.ToSQL()
	if err != nil {
		return 0, "", err
	}
	var blockhash string
	err = d.db.QueryRow(query).Scan(&blockhash)
	if err != nil {
		return 0, "", err
	}
	heightNum, err := strconv.Atoi(height)
	if err != nil {
		return 0, "", err
	}
	return heightNum, blockhash, nil
}

func (d *Database) insertRow(tableName string, queryVars []interface{}) error {

	d.canClose.Add(1)
	defer d.canClose.Done()
	switch tableName {
	case "blocks":
		return d.insertBlockRow(queryVars)
	case "block_headers":
		return d.insertBlockHeadersRow(queryVars)
	case "votes":
		return d.insertVote(queryVars)
	case "tx_single":
		return d.insertTxSingle(queryVars)
	case "deposits":
		return d.insertDeposit(queryVars)
	case "exits":
		return d.insertExit(queryVars)

	}
	return nil
}

func (d *Database) insertBlockRow(queryVars []interface{}) error {
	dw := goqu.Dialect(d.driver)
	ds := dw.Insert("blocks").Rows(
		goqu.Record{
			"block_hash":             queryVars[0],
			"block_signature":        queryVars[1],
			"block_randao_signature": queryVars[2],
			"height":                 queryVars[3],
		},
	)
	query, _, err := ds.ToSQL()
	if err != nil {
		return err
	}
	_, err = d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) insertBlockHeadersRow(queryVars []interface{}) error {
	dw := goqu.Dialect(d.driver)
	ds := dw.Insert("block_headers").Rows(
		goqu.Record{
			"block_hash":                    queryVars[0],
			"version":                       queryVars[1],
			"nonce":                         queryVars[2],
			"tx_mekle_root":                 queryVars[3],
			"tx_multi_merkle_root":          queryVars[4],
			"vote_merkle_root":              queryVars[5],
			"deposit_merkle_root":           queryVars[6],
			"exit_merkle_root":              queryVars[7],
			"vote_slashing_merkle_root":     queryVars[8],
			"randao_slashing_merkle_root":   queryVars[9],
			"proposer_slashing_merkle_root": queryVars[10],
			"governance_votes_merkle_root":  queryVars[11],
			"previous_block_hash":           queryVars[12],
			"timestamp":                     queryVars[13],
			"slot":                          queryVars[14],
			"state_root":                    queryVars[15],
			"fee_address":                   queryVars[16],
		},
	)
	query, _, err := ds.ToSQL()
	if err != nil {
		return err
	}
	_, err = d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) insertVote(queryVars []interface{}) error {
	dw := goqu.Dialect(d.driver)
	ds := dw.Insert("votes").Rows(
		goqu.Record{
			"block_hash":             queryVars[0],
			"signature":              queryVars[1],
			"participation_bitfield": queryVars[2],
			"data_slot":              queryVars[3],
			"data_from_epoch":        queryVars[4],
			"data_from_hash":         queryVars[5],
			"data_to_epoch":          queryVars[6],
			"data_to_hash":           queryVars[7],
			"data_beacon_block_hash": queryVars[8],
			"data_nonce":             queryVars[9],
			"vote_hash":              queryVars[10],
		},
	)
	query, _, err := ds.ToSQL()
	if err != nil {
		return err
	}
	_, err = d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) insertTxSingle(queryVars []interface{}) error {
	dw := goqu.Dialect(d.driver)
	ds := dw.Insert("tx_single").Rows(
		goqu.Record{
			"block_hash":      queryVars[0],
			"tx_type":         queryVars[1],
			"to_addr":         queryVars[2],
			"from_public_key": queryVars[3],
			"amount":          queryVars[4],
			"nonce":           queryVars[5],
			"fee":             queryVars[6],
			"signature":       queryVars[7],
		},
	)
	query, _, err := ds.ToSQL()
	if err != nil {
		return err
	}
	_, err = d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) insertDeposit(queryVars []interface{}) error {
	dw := goqu.Dialect(d.driver)
	ds := dw.Insert("deposits").Rows(
		goqu.Record{
			"block_hash":               queryVars[0],
			"public_key":               queryVars[1],
			"signature":                queryVars[2],
			"data_public_key":          queryVars[3],
			"data_proof_of_possession": queryVars[4],
			"data_withdrawal_address":  queryVars[5],
		},
	)
	query, _, err := ds.ToSQL()
	if err != nil {
		return err
	}
	_, err = d.db.Exec(query)
	if err != nil {
		return err
	}

	return d.addValidator(queryVars[1])
}

func (d *Database) insertExit(queryVars []interface{}) error {
	dw := goqu.Dialect(d.driver)
	ds := dw.Insert("exits").Rows(
		goqu.Record{
			"block_hash":            queryVars[0],
			"validator_public_key":  queryVars[1],
			"withdrawal_public_key": queryVars[2],
			"signature":             queryVars[3],
		},
	)
	query, _, err := ds.ToSQL()
	if err != nil {
		return err
	}
	_, err = d.db.Exec(query)
	if err != nil {
		return err
	}
	return d.exitValidator(queryVars[1])
}

func (d *Database) addValidator(valPubKey interface{}) error {
	dw := goqu.Dialect(d.driver)
	ds := dw.Insert("validators").Rows(
		goqu.Record{
			"public_key": valPubKey,
			"exit":       false,
		},
	)
	query, _, err := ds.ToSQL()
	if err != nil {
		return err
	}
	_, err = d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) exitValidator(valPubKey interface{}) error {
	dw := goqu.Dialect(d.driver)
	ds := dw.Update("validators").Set(
		goqu.Record{
			"exit": true,
		}).Where(
		goqu.Ex{
			"public_key": valPubKey,
		},
	)
	query, _, err := ds.ToSQL()
	if err != nil {
		return err
	}
	_, err = d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// NewDB creates a db client
func NewDB(dbConnString string, log logger.Logger, wg *sync.WaitGroup, driver string) *Database {
	db, err := sql.Open(driver, dbConnString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = runMigrations(driver, db)
	if err != nil {
		log.Fatal(err)
	}

	dbclient := &Database{
		log:      log,
		db:       db,
		canClose: wg,
		driver:   driver,
	}

	dbclient.log.Info("Database connection established")

	return dbclient
}

func runMigrations(driver string, db *sql.DB) error {
	var driverWrapper database.Driver
	var err error
	var migrationsString string
	switch driver {
	case "mysql":
		migrationsString = "file://cmd/ogen/indexer/migrations/mysql"
		driverWrapper, err = mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			return err
		}
	case "sqlite3":
		migrationsString = "file://cmd/ogen/indexer/migrations/sqlite3"
		driverWrapper, err = sqlite3.WithInstance(db, &sqlite3.Config{})
		if err != nil {
			return err
		}
	default:
		return errors.New("driver not supported")
	}

	m, _ := migrate.NewWithDatabaseInstance(
		migrationsString,
		driver,
		driverWrapper,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
