import { useQuery } from "../../../src-uploader/hooks/useQuery";
import { GameTile } from "../../components/GameTile";
import { GameMeta } from "../../types/GameMeta";

export function HomeView({}): JSX.Element {
  const { data: games } = useQuery<GameMeta[]>("games");

  return (
    <div class="inline-grid gap-4 p-4 sm:!grid-cols-3 md:!grid-cols-4 xs:grid-cols-2">
      {games?.map((game) => (
        <GameTile gamemeta={game} />
      ))}
    </div>
  );
}
