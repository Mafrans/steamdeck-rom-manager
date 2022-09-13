const baseURL = "http://localhost:3123";

export const apiFetcher = async (path: string, options: RequestInit) => {
  const url = new URL(path, baseURL);
  const res = await fetch(url.href, options);

  return await res.json();
};

export const getAPIUrl = (path: string) => new URL(path, baseURL).href;
