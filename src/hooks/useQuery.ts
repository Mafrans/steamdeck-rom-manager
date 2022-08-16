import useSWR from "swr";
import { apiFetcher } from "../utils/apiFetcher";

export const useQuery = <T>(path: string) => {
  const { data, error } = useSWR<T>(path, apiFetcher);

  return {
    data,
    error,
    isLoading: !data && !error,
  };
};
