import { useSWRConfig } from "swr";
import { apiFetcher } from "../utils/apiFetcher";

export const useMutate = <T>(path: string, options?: RequestInit) => {
  const { mutate } = useSWRConfig();
  return (data: T) => {
    apiFetcher(path, {
      ...options,
      method: options?.method ?? "POST",
      body: JSON.stringify(data),
    });
    mutate(path);
  };
};
