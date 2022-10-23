import { url } from "./config"

export async function fetchRealtime({ }) {
  console.log(
    "fetchRealtime",
    url(`hacker/realtime`)
  );

  return fetch(url(`hacker/realtime`))
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
}
