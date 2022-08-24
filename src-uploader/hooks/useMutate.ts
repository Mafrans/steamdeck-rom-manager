import { useSWRConfig } from "swr";
import { apiFetcher } from "../utils/apiFetcher";

export const useMutate = <T = any, K = unknown>(
  path: string,
  options?: RequestInit
): ((data: T) => Promise<K>) => {
  const { mutate } = useSWRConfig();
  return async (data: T) => {
    const response = await apiFetcher(path, {
      ...options,
      method: options?.method ?? "POST",
      body: JSON.stringify(data),
    });

    mutate(path);
    return response as K;
  };
};
