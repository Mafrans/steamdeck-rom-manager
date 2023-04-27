package main

import (
	"bytes"
	"compress/zlib"
	_ "embed"
	"hash/crc32"
	"io"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

//go:embed games/games.compressed.buf
var compressedGameDB []byte
var gameDB *Games

func readGameDB() *Games {
	buffer := bytes.NewBuffer(compressedGameDB)
	reader, err := zlib.NewReader(buffer)

	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	dbProtoBuffer, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	games := &Games{}
	if err := proto.Unmarshal(dbProtoBuffer, games); err != nil {
		log.Fatal(err)
	}

	return games
}

// IdentifyGame matches a game file to the data by its CRC hash
func IdentifyGame(file string) GameMeta {
	if gameDB == nil {
		gameDB = readGameDB()
	}

	content, _ := os.ReadFile(file)
	crc := crc32.Checksum(content, crc32.IEEETable)

	match := GameMeta{}
	for _, game := range gameDB.Games {
		if game.CrcHash == crc {
			match.Game = game
		}
	}

	return match
}
