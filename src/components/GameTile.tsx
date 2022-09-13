import { GameMeta } from "../types/GameMeta";
import { getAPIUrl } from "../utils/apiFetcher";

type GameTileProps = {
  gamemeta: GameMeta;
};

export function GameTile({ gamemeta }: GameTileProps): JSX.Element {
  return (
    <div class="">
      <img
        class="aspect-[1/1] w-full object-cover"
        src={getAPIUrl(`/games/${gamemeta.game.id}/cover`)}
      />
      <p class="text-center">{gamemeta.game.name}</p>
    </div>
  );
}
