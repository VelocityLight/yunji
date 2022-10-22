import { url } from "./config"

export async function fetchTeamCosts({ }) {
  console.log(
    "fetchTeamCost",
    url(`costs/`)
  );

  return fetch(url(`costs/`))
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
  // .catch((e) => {
  //   return e
  // });
}
