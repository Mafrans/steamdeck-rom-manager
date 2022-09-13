export type GameMeta = {
  game: {
    id: number;
    name: string;
    crcHash: number;
    console: string;
  };
  file: string;
  artwork: {
    cover: string;
  };
};
