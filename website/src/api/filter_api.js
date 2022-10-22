import { url } from "./config"

export async function fetchTags({ service }) {

  var queryString = ""
  if (service) {
    queryString = service
  }


  console.log(
    "fetchTags",
    url(`bills/used-by-tags?${queryString}`)
  );

  return fetch(url(`bills/used-by-tags?${queryString}`))
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
}

export async function fetchServices({ tags = [] }) {
  var queryString = ""
  if (tags.length > 0) {
    queryString = `
    ${tags.map((item) => {
      return "tags=" + item
    }).join("&")
      }`
  }

  console.log(
    "fetchService",
    url(`bills/component-tags?${queryString}`)
  );

  return fetch(url(`bills/component-tags?${queryString}`))
    .then(async (res) => {
      const data = await res.json();
      return data;
    })
}
