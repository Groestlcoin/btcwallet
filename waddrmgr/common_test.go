// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package waddrmgr

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcwallet/walletdb"
	_ "github.com/btcsuite/btcwallet/walletdb/bdb"
)

var (
	// seed is the master seed used throughout the tests.
	// BIP39 mnemonic:   "all all all all all all all all all all all all"
	// https://groestlcoin.org/bip39/
	seed = []byte{
		0xc7, 0x6c, 0x4a, 0xc4, 0xf4, 0xe4, 0xa0, 0x0d, 0x6b, 0x27, 0x4d, 0x5c,
		0x39, 0xc7, 0x00, 0xbb, 0x4a, 0x7d, 0xdc, 0x04, 0xfb, 0xc6, 0xf7, 0x8e,
		0x85, 0xca, 0x75, 0x00, 0x7b, 0x5b, 0x49, 0x5f, 0x74, 0xa9, 0x04, 0x3e,
		0xeb, 0x77, 0xbd, 0xd5, 0x3a, 0xa6, 0xfc, 0x3a, 0x0e, 0x31, 0x46, 0x22,
		0x70, 0x31, 0x6f, 0xa0, 0x4b, 0x8c, 0x19, 0x11, 0x4c, 0x87, 0x98, 0x70,
		0x6c, 0xd0, 0x2a, 0xc8,
	}

	originalSeed = []byte{
		0x2a, 0x64, 0xdf, 0x08, 0x5e, 0xef, 0xed, 0xd8, 0xbf,
		0xdb, 0xb3, 0x31, 0x76, 0xb5, 0xba, 0x2e, 0x62, 0xe8,
		0xbe, 0x8b, 0x56, 0xc8, 0x83, 0x77, 0x95, 0x59, 0x8b,
		0xb6, 0xc4, 0x40, 0xc0, 0x64,
	}

	pubPassphrase   = []byte("_DJr{fL4H0O}*-0\n:V1izc)(6BomK")
	privPassphrase  = []byte("81lUHXnOMZ@?XXd7O9xyDIWIbXX-lj")
	pubPassphrase2  = []byte("-0NV4P~VSJBWbunw}%<Z]fuGpbN[ZI")
	privPassphrase2 = []byte("~{<]08%6!-?2s<$(8$8:f(5[4/!/{Y")

	// fastScrypt are parameters used throughout the tests to speed up the
	// scrypt operations.
	fastScrypt = &FastScryptOptions

	// waddrmgrNamespaceKey is the namespace key for the waddrmgr package.
	waddrmgrNamespaceKey = []byte("waddrmgrNamespace")

	// expectedAddrs is the list of all expected addresses generated from the
	// seed.
	expectedAddrs = []expectedAddr{
		{
			address:     "Fj62rBJi8LvbmWu2jzkaUX1NFXLEqDLoZM",
			addressHash: hexToBytes("98af0aaca388a7e1024f505c033626d908e3b54a"),
			internal:    false,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("03b85cc59b67c35851eb5060cfc3a759a482254553c5857075c9e247d74d412c91"),
			privKey:     hexToBytes("3c3385ddc6fd95ba7282051aeb440bc75820b8c10db5c83c052d7586e3e98e84"),
			privKeyWIF:  "KyEjYKtiAqyERxq6f9SMQ29GinrThjVrEmfdUrKZz6ZPnPxr8Hor",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  0,
				Index:   0,
			},
		},
		{
			address:     "FYy3bTDYJiSaNhh4d2ptHGwAPNRc6heKy2",
			addressHash: hexToBytes("29ab7eb4d7a419548efab4f10013627e574a5d0a"),
			internal:    false,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("02cf5126ff54e38a80a919579d7091cafe24840eab1d30fe2b4d59bdd9d267cad8"),
			privKey:     hexToBytes("4562a53b7245ebf89494b25ca3da446be64ffe82ecca1e5cb33daea70e559e14"),
			privKeyWIF:  "KyYazbWftZUkCf2k9YFQr6UXtfjAw3vnZXyep6pz9PWATWm6wKaL",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  0,
				Index:   1,
			},
		},
		{
			address:     "FXHDsC5ZqWQHkDmShzgRVZ1MatpWhwxTAA",
			addressHash: hexToBytes("172b4e06e9b7881a48d2ee8062b495d0b2517fe8"),
			internal:    false,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("0331693756f749180aeed0a65a0fab0625a2250bd9abca502282a4cf0723152e67"),
			privKey:     hexToBytes("747c9cbc72fbd247f12d93cb44d2fef6d40a423037190997702cd8452dae53d8"),
			privKeyWIF:  "L189RB5TvaJX6p3mnjaoJ12R2GGzdxu1iDUvJPdT3d9Wh8c3g9q9",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  0,
				Index:   2,
			},
		},
		{
			address:     "FtM4zAn9aVYgHgxmamWBgWPyZsb6RhvkA9",
			addressHash: hexToBytes("fe40329c95c5598ac60752a5310b320cb52d18e6"),
			internal:    false,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("0286b2a6246bfed0f9a3a4e2ccb49b6989fe078177580b763bbe01e3d4fdfecacd"),
			privKey:     hexToBytes("15a862eaab1bc586448466a31974dade05f7a93bd126d75a5a841c2a86b70c4e"),
			privKeyWIF:  "KwwoyZrELnXJc1mvuviWCWc3xSZDBUfnpgwee81B6H83myTQ43y9",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  0,
				Index:   3,
			},
		},
		{
			address:     "FjE6TV6jcN12fbrzwn44GFDcmFQvU9driV",
			addressHash: hexToBytes("9a3561be88ec6a3d4d1ae93e282312f99f89d3fd"),
			internal:    false,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("03c20962ab16f4d97a4f6f8b83f73a05457794ced25debbf8299336e6ac48bf40d"),
			privKey:     hexToBytes("c5d40555caebd7534b848262595f710c6e4e7776fb8a978bc3674283f7e97fbe"),
			privKeyWIF:  "L3rGDCVjokG5caEwpxkQSUuDAQc7arsHCzSiFgzqMpJckyVVXAv9",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  0,
				Index:   4,
			},
		},
		{
			address:     "FmRaqvVBRrAp2Umfqx9V1ectZy8gw54QDN",
			addressHash: hexToBytes("b251e7b5a9f5fbd0770585e967ee6df28dc792c0"),
			internal:    true,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("036d7f78ae929a35acc4e955984f7ab22c2ff0ae9921068db01d162f1fc3f852cb"),
			privKey:     hexToBytes("adf68f0681dbc14bafdc31f04d9670657ce36f21026b86e995ca5d243fe5183a"),
			privKeyWIF:  "L33sZBsT58r8RiJAgqp6nJ7MBmuPPWHLb95SLkkhBjx3roAkSWY3",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  1,
				Index:   0,
			},
		},
		{
			address:     "Fmhtxeh7YdCBkyQF7AQG4QnY8y3rJg89di",
			addressHash: hexToBytes("b567aed7740bb1358557c08221be52e64cbe89fb"),
			internal:    true,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("0284b71f7ccb4b63e6fb370661ed665d5a57b435c510b5cf8c16daa5b0bfb788ed"),
			privKey:     hexToBytes("885330d93b2771df092466d3c9095c359bf570cfaba30477b22cf0c182a32321"),
			privKeyWIF:  "L1ni4eXUQ5SCTi8CLJViv9rfsXruNPE6wVdCHtG4Ng93m7VjH2Ze",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  1,
				Index:   1,
			},
		},
		{
			address:     "Fk1ujKDRdcNCrtBLu25od3ohffRUSkU1KQ",
			addressHash: hexToBytes("a2df8fc1515080de2d2608a29a58aeeadb48635a"),
			internal:    true,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("026f66ab5e316597c94741dec1455ec342a71cae4837f8cab47dedab21ff2a28bf"),
			privKey:     hexToBytes("529a8e9cccc205b7d14ea43c95dbe3d0783fcc6b3fcf570f61cc600bf46a24c8"),
			privKeyWIF:  "KyzHJ5RRWtfK94fj4eQRbCRd7yb9AYRc8YoSu9yLk98sUknofHxN",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  1,
				Index:   2,
			},
		},
		{
			address:     "FYHMC9gKwcfqEQgGkpyPssr8Y4RV5xiwd4",
			addressHash: hexToBytes("22298bf7482ba5cbc02e60465cac6244401a5d04"),
			internal:    true,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("03612b14000e8f927e1026cf55a04e937ce71f739ac431fdd8651f8e92ecc73986"),
			privKey:     hexToBytes("2bd81790c067dd55ff123787d10373bda6cbc5d213d6dae9d51c2d1ed0abf6e5"),
			privKeyWIF:  "KxgwN2RFvxXnEK63wy5tp3dpMw2z9tzca6PG5BGyXztkKXhY4q4k",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  1,
				Index:   3,
			},
		},
		{
			address:     "FmECuazFSW4av8u1Z4yM2PM18k7VC53MZ9",
			addressHash: hexToBytes("b02b0480e6a061adf5f7f12645e92633d1631fbf"),
			internal:    true,
			compressed:  true,
			imported:    false,
			pubKey:      hexToBytes("024667a988b010c2dfdba160aa36b94cf9d5948dec2d3be64fd7792ff6bd72f117"),
			privKey:     hexToBytes("a801a5a9810ff80d874db61b7b1c592af274416d40187cc828837bf9116ca67b"),
			privKeyWIF:  "L2rHyNLW2a74moTRZKdPb2h4HeVdHQAuYAMu8qC6K2tCBYXf5Hwq",
			derivationInfo: DerivationPath{
				Account: 0,
				Branch:  1,
				Index:   4,
			},
		},
	}

	// expectedExternalAddrs is the list of expected external addresses
	// generated from the seed
	expectedExternalAddrs = expectedAddrs[:5]

	// expectedInternalAddrs is the list of expected internal addresses
	// generated from the seed
	expectedInternalAddrs = expectedAddrs[5:]
)

// checkManagerError ensures the passed error is a ManagerError with an error
// code that matches the passed  error code.
func checkManagerError(t *testing.T, testName string, gotErr error,
	wantErrCode ErrorCode) bool {

	merr, ok := gotErr.(ManagerError)
	if !ok {
		t.Errorf("%s: unexpected error type - got %T, want %T",
			testName, gotErr, ManagerError{})
		return false
	}
	if merr.ErrorCode != wantErrCode {
		t.Errorf("%s: unexpected error code - got %s (%s), want %s",
			testName, merr.ErrorCode, merr.Description, wantErrCode)
		return false
	}

	return true
}

// hexToBytes is a wrapper around hex.DecodeString that panics if there is an
// error.  It MUST only be used with hard coded values in the tests.
func hexToBytes(origHex string) []byte {
	buf, err := hex.DecodeString(origHex)
	if err != nil {
		panic(err)
	}
	return buf
}

func emptyDB(t *testing.T) (tearDownFunc func(), db walletdb.DB) {
	dirName, err := ioutil.TempDir("", "mgrtest")
	if err != nil {
		t.Fatalf("Failed to create db temp dir: %v", err)
	}
	dbPath := filepath.Join(dirName, "mgrtest.db")
	db, err = walletdb.Create("bdb", dbPath, true)
	if err != nil {
		_ = os.RemoveAll(dirName)
		t.Fatalf("createDbNamespace: unexpected error: %v", err)
	}
	tearDownFunc = func() {
		db.Close()
		_ = os.RemoveAll(dirName)
	}
	return
}

// setupManager creates a new address manager and returns a teardown function
// that should be invoked to ensure it is closed and removed upon completion.
func setupManager(t *testing.T) (tearDownFunc func(), db walletdb.DB, mgr *Manager) {
	// Create a new manager in a temp directory.
	dirName, err := ioutil.TempDir("", "mgrtest")
	if err != nil {
		t.Fatalf("Failed to create db temp dir: %v", err)
	}
	dbPath := filepath.Join(dirName, "mgrtest.db")
	db, err = walletdb.Create("bdb", dbPath, true)
	if err != nil {
		_ = os.RemoveAll(dirName)
		t.Fatalf("createDbNamespace: unexpected error: %v", err)
	}
	err = walletdb.Update(db, func(tx walletdb.ReadWriteTx) error {
		ns, err := tx.CreateTopLevelBucket(waddrmgrNamespaceKey)
		if err != nil {
			return err
		}
		err = Create(
			ns, seed, pubPassphrase, privPassphrase,
			&chaincfg.MainNetParams, fastScrypt, time.Time{},
		)
		if err != nil {
			return err
		}
		mgr, err = Open(ns, pubPassphrase, &chaincfg.MainNetParams)
		return err
	})
	if err != nil {
		db.Close()
		_ = os.RemoveAll(dirName)
		t.Fatalf("Failed to create Manager: %v", err)
	}
	tearDownFunc = func() {
		mgr.Close()
		db.Close()
		_ = os.RemoveAll(dirName)
	}
	return tearDownFunc, db, mgr
}
