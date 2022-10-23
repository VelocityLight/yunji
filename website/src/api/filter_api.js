import { url } from "./config"

export async function fetchServices() {
  console.log(
    "fetchServices",
    url(`bills/services`)
  );

  return fetch(url(`bills/services`))
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
}

export async function fetchTags() {
  console.log(
    "fetchTags",
    url(`bills/component-tags`)
  );

  return fetch(url(`bills/component-tags`))
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
}
